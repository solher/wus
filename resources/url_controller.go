// @SubApi Url resource [/urls]
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

// @Title RedirectUrl
// @Description Redirect a short Url to the original one
// @Accept  json
// @Param   shortHandle path string true "Shortened Url"
// @Success 301 {object} error "Request was successful"
// @Router /{shortHandle} [get]
// @Resource urls
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

// @Title Create
// @Description Create one or multiple Url instances
// @Accept  json
// @Param   Url body domain.Url true "Url instance(s) data"
// @Success 201 {object} domain.Url "Request was successful"
// @Router /urls [post]
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

// @Title Find
// @Description Find all Url instances matched by filter
// @Accept  json
// @Param   filter query string false "JSON filter defining fields and includes"
// @Success 200 {object} domain.Url "Request was successful"
// @Router /urls [get]
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
// @Description Find a Url instance
// @Accept  json
// @Param   id path int true "Url id"
// @Param   filter query string false "JSON filter defining fields and includes"
// @Success 200 {object} domain.Url "Request was successful"
// @Router /urls/{id} [get]
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

// @Title Upsert
// @Description Upsert one or multiple Url instances
// @Accept  json
// @Param   Url body domain.Url true "Url instance(s) data"
// @Success 201 {object} domain.Url "Request was successful"
// @Router /urls [put]
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

// @Title UpdateByID
// @Description Update attributes of a Url instance
// @Accept  json
// @Param   id path int true "Url id"
// @Param   Url body domain.Url true "Url instance data"
// @Success 201 {object} domain.Url
// @Router /urls/{id} [put]
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

// @Title DeleteAll
// @Description Delete all Url instances matched by filter
// @Accept  json
// @Param   filter query string false "JSON filter defining fields and includes"
// @Success 204 {object} error "Request was successful"
// @Router /urls [delete]
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

// @Title DeleteByID
// @Description Delete a Url instance
// @Accept  json
// @Param   id path int true "Url id"
// @Success 204 {object} error "Request was successful"
// @Router /urls/{id} [delete]
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

// @Title CreateRelated
// @Description Create one or multiple Url instances of a related resource
// @Accept  json
// @Param   pk path int true "Url id"
// @Param   relatedResource path string true "Related resource name"
// @Param   Url body domain.Url true "Url instance(s) data"
// @Success 201 {object} domain.Url "Request was successful"
// @Router /urls/{pk}/{relatedResource} [post]
func (c *UrlCtrl) CreateRelated(w http.ResponseWriter, r *http.Request, params map[string]string) {
	c.related(w, r, params)
}

// @Title FindRelated
// @Description Find all Url instances  of a related resource matched by filter
// @Accept  json
// @Param   pk path int true "Url id"
// @Param   relatedResource path string true "Related resource name"
// @Param   filter query string false "JSON filter defining fields and includes"
// @Success 200 {object} domain.Url "Request was successful"
// @Router /urls/{pk}/{relatedResource} [get]
func (c *UrlCtrl) FindRelated(w http.ResponseWriter, r *http.Request, params map[string]string) {
	c.related(w, r, params)
}

// @Title UpsertRelated
// @Description Upsert one or multiple Url instances of a related resource
// @Accept  json
// @Param   pk path int true "Url id"
// @Param   relatedResource path string true "Related resource name"
// @Param   Url body domain.Url true "Url instance(s) data"
// @Success 201 {object} domain.Url "Request was successful"
// @Router /urls/{pk}/{relatedResource} [put]
func (c *UrlCtrl) UpsertRelated(w http.ResponseWriter, r *http.Request, params map[string]string) {
	c.related(w, r, params)
}

// @Title DeleteAllRelated
// @Description Delete all Url instances of a related resource matched by filter
// @Accept  json
// @Param   pk path int true "Url id"
// @Param   relatedResource path string true "Related resource name"
// @Param   filter query string false "JSON filter defining fields and includes"
// @Success 204 {object} error "Request was successful"
// @Router /urls/{pk}/{relatedResource} [delete]
func (c *UrlCtrl) DeleteAllRelated(w http.ResponseWriter, r *http.Request, params map[string]string) {
	c.related(w, r, params)
}

func (c *UrlCtrl) related(w http.ResponseWriter, r *http.Request, params map[string]string) {
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

// @Title FindByIDRelated
// @Description Find a Url instance of a related resource
// @Accept  json
// @Param   pk path int true "Url id"
// @Param   relatedResource path string true "Related resource name"
// @Param   fk path int true "Related resource id"
// @Param   filter query string false "JSON filter defining fields and includes"
// @Success 200 {object} domain.Url "Request was successful"
// @Router /urls/{pk}/{relatedResource}/{fk} [get]
func (c *UrlCtrl) FindByIDRelated(w http.ResponseWriter, r *http.Request, params map[string]string) {
	c.relatedOne(w, r, params)
}

// @Title UpdateByIDRelated
// @Description Update attributes of a Url instance of a related resource
// @Accept  json
// @Param   pk path int true "Url id"
// @Param   relatedResource path string true "Related resource name"
// @Param   fk path int true "Related resource id"
// @Param   Url body domain.Url true "Url instance data"
// @Success 201 {object} domain.Url
// @Router /urls/{pk}/{relatedResource}/{fk} [put]
func (c *UrlCtrl) UpdateByIDRelated(w http.ResponseWriter, r *http.Request, params map[string]string) {
	c.relatedOne(w, r, params)
}

// @Title DeleteByIDRelated
// @Description Delete a Url instance of a related resource
// @Accept  json
// @Param   pk path int true "Url id"
// @Param   relatedResource path string true "Related resource name"
// @Param   fk path int true "Related resource id"
// @Success 204 {object} error "Request was successful"
// @Router /urls/{pk}/{relatedResource}/{fk} [delete]
func (c *UrlCtrl) DeleteByIDRelated(w http.ResponseWriter, r *http.Request, params map[string]string) {
	c.relatedOne(w, r, params)
}

func (c *UrlCtrl) relatedOne(w http.ResponseWriter, r *http.Request, params map[string]string) {
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
