package resources

import (
	"database/sql"

	"github.com/solher/wus/domain"
	"github.com/solher/zest/internalerrors"
	"github.com/solher/zest/usecases"
)

func init() {
	usecases.DependencyDirectory.Register(NewUrlInter)
	usecases.DependencyDirectory.Register(PopulateUrlInter)
}

type AbstractUrlRepo interface {
	Create(urls []domain.Url) ([]domain.Url, error)
	CreateOne(url *domain.Url) (*domain.Url, error)
	Find(context usecases.QueryContext) ([]domain.Url, error)
	FindByID(id int, context usecases.QueryContext) (*domain.Url, error)
	Update(urls []domain.Url, context usecases.QueryContext) ([]domain.Url, error)
	UpdateByID(id int, url *domain.Url, context usecases.QueryContext) (*domain.Url, error)
	DeleteAll(context usecases.QueryContext) error
	DeleteByID(id int, context usecases.QueryContext) error
	Raw(query string, values ...interface{}) (*sql.Rows, error)
}

type UrlInter struct {
	repo AbstractUrlRepo
}

func NewUrlInter(repo AbstractUrlRepo) *UrlInter {
	return &UrlInter{repo: repo}
}

func PopulateUrlInter(urlInter *UrlInter, repo AbstractUrlRepo) {
	if urlInter.repo == nil {
		urlInter.repo = repo
	}
}

func (i *UrlInter) GetLongUrl(shortHandle string) (string, error) {
	filter := &usecases.Filter{
		Limit: 1,
		Where: map[string]interface{}{"shortHandle": shortHandle},
	}

	urls, err := i.repo.Find(usecases.QueryContext{Filter: filter})
	if err != nil {
		return "", err
	}

	if len(urls) == 0 {
		return "", internalerrors.NotFound
	}

	return urls[0].LongUrl, nil
}

func (i *UrlInter) Create(urls []domain.Url) ([]domain.Url, error) {
	urls, err := i.BeforeCreate(urls)
	if err != nil {
		return nil, err
	}

	urls, err = i.repo.Create(urls)
	if err != nil {
		return nil, err
	}

	urls, err = i.AfterCreate(urls)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func (i *UrlInter) CreateOne(url *domain.Url) (*domain.Url, error) {
	urls, err := i.Create([]domain.Url{*url})
	if err != nil {
		return nil, err
	}

	return &urls[0], nil
}

func (i *UrlInter) Find(context usecases.QueryContext) ([]domain.Url, error) {
	urls, err := i.repo.Find(context)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func (i *UrlInter) FindByID(id int, context usecases.QueryContext) (*domain.Url, error) {
	url, err := i.repo.FindByID(id, context)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (i *UrlInter) Upsert(urls []domain.Url, context usecases.QueryContext) ([]domain.Url, error) {
	urlsToUpdate := []domain.Url{}
	urlsToCreate := []domain.Url{}

	for k := range urls {
		if urls[k].ID != 0 {
			urlsToUpdate = append(urlsToUpdate, urls[k])
		} else {
			urlsToCreate = append(urlsToCreate, urls[k])
		}
	}

	urlsToUpdate, err := i.BeforeUpdate(urlsToUpdate)
	if err != nil {
		return nil, err
	}

	urlsToUpdate, err = i.repo.Update(urlsToUpdate, context)
	if err != nil {
		return nil, err
	}

	urlsToUpdate, err = i.AfterUpdate(urlsToUpdate)
	if err != nil {
		return nil, err
	}

	urlsToCreate, err = i.BeforeCreate(urlsToCreate)
	if err != nil {
		return nil, err
	}

	urlsToCreate, err = i.repo.Create(urlsToCreate)
	if err != nil {
		return nil, err
	}

	urlsToCreate, err = i.AfterCreate(urlsToCreate)
	if err != nil {
		return nil, err
	}

	return append(urlsToUpdate, urlsToCreate...), nil
}

func (i *UrlInter) UpsertOne(url *domain.Url, context usecases.QueryContext) (*domain.Url, error) {
	urls, err := i.Upsert([]domain.Url{*url}, context)
	if err != nil {
		return nil, err
	}

	return &urls[0], nil
}

func (i *UrlInter) UpdateByID(id int, url *domain.Url,
	context usecases.QueryContext) (*domain.Url, error) {

	urls, err := i.BeforeUpdate([]domain.Url{*url})
	if err != nil {
		return nil, err
	}

	url = &urls[0]

	url, err = i.repo.UpdateByID(id, url, context)
	if err != nil {
		return nil, err
	}

	urls, err = i.AfterUpdate([]domain.Url{*url})
	if err != nil {
		return nil, err
	}

	return &urls[0], nil
}

func (i *UrlInter) DeleteAll(context usecases.QueryContext) error {
	urls, err := i.repo.Find(context)
	if err != nil {
		return err
	}

	urls, err = i.BeforeDelete(urls)
	if err != nil {
		return err
	}

	err = i.repo.DeleteAll(context)
	if err != nil {
		return err
	}

	_, err = i.AfterDelete(urls)
	if err != nil {
		return err
	}

	return nil
}

func (i *UrlInter) DeleteByID(id int, context usecases.QueryContext) error {
	url, err := i.repo.FindByID(id, context)
	if err != nil {
		return err
	}

	urls, err := i.BeforeDelete([]domain.Url{*url})
	if err != nil {
		return err
	}

	url = &urls[0]

	err = i.repo.DeleteByID(id, context)
	if err != nil {
		return err
	}

	_, err = i.AfterDelete([]domain.Url{*url})
	if err != nil {
		return err
	}

	return nil
}
