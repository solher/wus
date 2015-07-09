// @SubApi URL management API [/urls]
package resources

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/solher/wus/domain"
	"github.com/solher/zest/apierrors"
	"github.com/solher/zest/interfaces"
	"github.com/solher/zest/internalerrors"
	"github.com/solher/zest/usecases"
)

func init() {
	usecases.DependencyDirectory.Register(NewUrlCtrl)
}

type AbstractUrlInter interface {
	Create(urls []domain.Url) ([]domain.Url, error)
	CreateOne(url *domain.Url) (*domain.Url, error)
	Find(context usecases.QueryContext) ([]domain.Url, error)
	FindByID(id int, context usecases.QueryContext) (*domain.Url, error)
	Upsert(urls []domain.Url, context usecases.QueryContext) ([]domain.Url, error)
	UpsertOne(url *domain.Url, context usecases.QueryContext) (*domain.Url, error)
	UpdateByID(id int, url *domain.Url, context usecases.QueryContext) (*domain.Url, error)
	DeleteAll(context usecases.QueryContext) error
	DeleteByID(id int, context usecases.QueryContext) error
	GetLongUrl(shortHandle string) (string, error)
}

type UrlCtrl struct {
	interactor AbstractUrlInter
	render     interfaces.AbstractRender
	routeDir   *usecases.RouteDirectory
}

func NewUrlCtrl(interactor AbstractUrlInter, render interfaces.AbstractRender, routeDir *usecases.RouteDirectory) *UrlCtrl {
	controller := &UrlCtrl{interactor: interactor, render: render, routeDir: routeDir}

	if routeDir != nil {
		setUrlRoutes(routeDir, controller)
	}

	return controller
}

func (c *UrlCtrl) RedirectUrl(w http.ResponseWriter, r *http.Request, params map[string]string) {
	shortHandle := params["shortHandle"]

	longurl, err := c.interactor.GetLongUrl(shortHandle)
	if err != nil {
		c.render.JSONError(w, http.StatusUnauthorized, apierrors.Unauthorized, err)
		return
	}

	w.Header().Set("Location", longurl)
	c.render.JSON(w, http.StatusMovedPermanently, nil)
}

func (c *UrlCtrl) Create(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	url := &domain.Url{}
	var urls []domain.Url

	buffer, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(buffer, url)
	if err != nil {
		err := json.Unmarshal(buffer, &urls)
		if err != nil {
			c.render.JSONError(w, http.StatusBadRequest, apierrors.BodyDecodingError, err)
			return
		}
	}

	lastResource := interfaces.GetLastResource(r)

	if urls == nil {
		url.SetRelatedID(lastResource.IDKey, lastResource.ID)
		url, err = c.interactor.CreateOne(url)
	} else {
		for i := range urls {
			(&urls[i]).SetRelatedID(lastResource.IDKey, lastResource.ID)
		}
		urls, err = c.interactor.Create(urls)
	}

	if err != nil {
		switch err.(type) {
		case *internalerrors.ViolatedConstraint:
			c.render.JSONError(w, 422, apierrors.ViolatedConstraint, err)
		default:
			c.render.JSONError(w, http.StatusInternalServerError, apierrors.InternalServerError, err)
		}
		return
	}

	if urls == nil {
		url.BeforeRender()
		c.render.JSON(w, http.StatusCreated, url)
	} else {
		for i := range urls {
			(&urls[i]).BeforeRender()
		}
		c.render.JSON(w, http.StatusCreated, urls)
	}
}

func (c *UrlCtrl) Find(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	filter, err := interfaces.GetQueryFilter(r)
	if err != nil {
		c.render.JSONError(w, http.StatusBadRequest, apierrors.FilterDecodingError, err)
		return
	}

	filter = interfaces.FilterIfLastResource(r, filter)
	filter = interfaces.FilterIfOwnerRelations(r, filter)
	relations := interfaces.GetOwnerRelations(r)

	urls, err := c.interactor.Find(usecases.QueryContext{Filter: filter, OwnerRelations: relations})
	if err != nil {
		c.render.JSONError(w, http.StatusInternalServerError, apierrors.InternalServerError, err)
		return
	}

	for i := range urls {
		(&urls[i]).BeforeRender()
	}
	c.render.JSON(w, http.StatusOK, urls)
}

// @Title FindByID
// @Description Get URL by ID
// @Accept  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {object} string
// @Failure 400 {object} apierrors.APIError
// @Router /urls/{some_id} [get]
func (c *UrlCtrl) FindByID(w http.ResponseWriter, r *http.Request, params map[string]string) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		c.render.JSONError(w, http.StatusBadRequest, apierrors.InvalidPathParams, err)
		return
	}

	filter, err := interfaces.GetQueryFilter(r)
	if err != nil {
		c.render.JSONError(w, http.StatusBadRequest, apierrors.FilterDecodingError, err)
		return
	}

	filter = interfaces.FilterIfOwnerRelations(r, filter)
	relations := interfaces.GetOwnerRelations(r)

	url, err := c.interactor.FindByID(id, usecases.QueryContext{Filter: filter, OwnerRelations: relations})
	if err != nil {
		switch err {
		case internalerrors.NotFound:
			c.render.JSONError(w, http.StatusUnauthorized, apierrors.Unauthorized, err)
		default:
			c.render.JSONError(w, http.StatusInternalServerError, apierrors.InternalServerError, err)
		}
		return
	}

	url.BeforeRender()
	c.render.JSON(w, http.StatusOK, url)
}

func (c *UrlCtrl) Upsert(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	url := &domain.Url{}
	var urls []domain.Url

	buffer, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(buffer, url)
	if err != nil {
		err := json.Unmarshal(buffer, &urls)
		if err != nil {
			c.render.JSONError(w, http.StatusBadRequest, apierrors.BodyDecodingError, err)
			return
		}
	}

	lastResource := interfaces.GetLastResource(r)
	filter := interfaces.FilterIfOwnerRelations(r, nil)
	relations := interfaces.GetOwnerRelations(r)

	if urls == nil {
		url.SetRelatedID(lastResource.IDKey, lastResource.ID)
		url, err = c.interactor.UpsertOne(url, usecases.QueryContext{Filter: filter, OwnerRelations: relations})
	} else {
		for i := range urls {
			(&urls[i]).SetRelatedID(lastResource.IDKey, lastResource.ID)
		}
		urls, err = c.interactor.Upsert(urls, usecases.QueryContext{Filter: filter, OwnerRelations: relations})
	}

	if err != nil {
		switch err.(type) {
		case *internalerrors.ViolatedConstraint:
			c.render.JSONError(w, 422, apierrors.ViolatedConstraint, err)
		}

		switch err {
		case internalerrors.NotFound:
			c.render.JSONError(w, http.StatusUnauthorized, apierrors.Unauthorized, err)
		default:
			c.render.JSONError(w, http.StatusInternalServerError, apierrors.InternalServerError, err)
		}

		return
	}

	if urls == nil {
		url.BeforeRender()
		c.render.JSON(w, http.StatusCreated, url)
	} else {
		for i := range urls {
			(&urls[i]).BeforeRender()
		}
		c.render.JSON(w, http.StatusCreated, urls)
	}
}

func (c *UrlCtrl) UpdateByID(w http.ResponseWriter, r *http.Request, params map[string]string) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		c.render.JSONError(w, http.StatusBadRequest, apierrors.InvalidPathParams, err)
		return
	}

	url := &domain.Url{}

	err = json.NewDecoder(r.Body).Decode(url)
	if err != nil {
		c.render.JSONError(w, http.StatusBadRequest, apierrors.BodyDecodingError, err)
		return
	}

	lastResource := interfaces.GetLastResource(r)
	filter := interfaces.FilterIfOwnerRelations(r, nil)
	relations := interfaces.GetOwnerRelations(r)

	url.SetRelatedID(lastResource.IDKey, lastResource.ID)
	url, err = c.interactor.UpdateByID(id, url, usecases.QueryContext{Filter: filter, OwnerRelations: relations})

	if err != nil {
		switch err {
		case internalerrors.NotFound:
			c.render.JSONError(w, http.StatusUnauthorized, apierrors.Unauthorized, err)
		default:
			c.render.JSONError(w, http.StatusInternalServerError, apierrors.InternalServerError, err)
		}
		return
	}

	url.BeforeRender()
	c.render.JSON(w, http.StatusCreated, url)
}

func (c *UrlCtrl) DeleteAll(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	filter, err := interfaces.GetQueryFilter(r)
	if err != nil {
		c.render.JSONError(w, http.StatusBadRequest, apierrors.FilterDecodingError, err)
		return
	}

	filter = interfaces.FilterIfLastResource(r, filter)
	filter = interfaces.FilterIfOwnerRelations(r, filter)
	relations := interfaces.GetOwnerRelations(r)

	err = c.interactor.DeleteAll(usecases.QueryContext{Filter: filter, OwnerRelations: relations})
	if err != nil {
		c.render.JSONError(w, http.StatusInternalServerError, apierrors.InternalServerError, err)
		return
	}

	c.render.JSON(w, http.StatusNoContent, nil)
}

func (c *UrlCtrl) DeleteByID(w http.ResponseWriter, r *http.Request, params map[string]string) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		c.render.JSONError(w, http.StatusBadRequest, apierrors.InvalidPathParams, err)
		return
	}

	filter := interfaces.FilterIfOwnerRelations(r, nil)
	relations := interfaces.GetOwnerRelations(r)

	err = c.interactor.DeleteByID(id, usecases.QueryContext{Filter: filter, OwnerRelations: relations})
	if err != nil {
		switch err {
		case internalerrors.NotFound:
			c.render.JSONError(w, http.StatusUnauthorized, apierrors.Unauthorized, err)
		default:
			c.render.JSONError(w, http.StatusInternalServerError, apierrors.InternalServerError, err)
		}
		return
	}

	c.render.JSON(w, http.StatusNoContent, nil)
}

func (c *UrlCtrl) Related(w http.ResponseWriter, r *http.Request, params map[string]string) {
	pk, err := strconv.Atoi(params["pk"])
	if err != nil {
		c.render.JSONError(w, http.StatusBadRequest, apierrors.InvalidPathParams, err)
		return
	}

	related := params["related"]
	key := usecases.NewDirectoryKey(related)

	var handler usecases.HandlerFunc
	switch r.Method {
	case "POST":
		handler = c.routeDir.Get(key.For("Create")).EffectiveHandler
	case "GET":
		handler = c.routeDir.Get(key.For("Find")).EffectiveHandler
	case "PUT":
		handler = c.routeDir.Get(key.For("Upsert")).EffectiveHandler
	case "DELETE":
		handler = c.routeDir.Get(key.For("DeleteAll")).EffectiveHandler
	}

	if handler == nil {
		c.render.JSON(w, http.StatusNotFound, nil)
		return
	}

	context.Set(r, "lastResource", &interfaces.Resource{Name: related, IDKey: "urlID", ID: pk})

	handler(w, r, params)
}

func (c *UrlCtrl) RelatedOne(w http.ResponseWriter, r *http.Request, params map[string]string) {
	pk, err := strconv.Atoi(params["pk"])
	if err != nil {
		c.render.JSONError(w, http.StatusBadRequest, apierrors.InvalidPathParams, err)
		return
	}

	params["id"] = params["fk"]

	related := params["related"]
	key := usecases.NewDirectoryKey(related)

	var handler usecases.HandlerFunc

	switch r.Method {
	case "GET":
		handler = c.routeDir.Get(key.For("FindByID")).EffectiveHandler
	case "PUT":
		handler = c.routeDir.Get(key.For("UpdateByID")).EffectiveHandler
	case "DELETE":
		handler = c.routeDir.Get(key.For("DeleteByID")).EffectiveHandler
	}

	if handler == nil {
		c.render.JSON(w, http.StatusNotFound, nil)
		return
	}

	context.Set(r, "lastResource", &interfaces.Resource{Name: related, IDKey: "urlID", ID: pk})

	handler(w, r, params)
}
