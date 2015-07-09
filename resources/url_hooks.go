package resources

import (
	"github.com/solher/wus/domain"
	zestdomain "github.com/solher/zest/domain"
	"time"
)

func (i *UrlInter) scopeModel(url *domain.Url) error {
	url.CreatedAt = time.Time{}
	url.UpdatedAt = time.Time{}
	url.Account = zestdomain.Account{}

	return nil
}

func (i *UrlInter) BeforeCreate(urls []domain.Url) ([]domain.Url, error) {
	for k := range urls {
		urls[k].ID = 0
		err := i.scopeModel(&urls[k])
		if err != nil {
			return nil, err
		}
	}
	return urls, nil
}

func (i *UrlInter) AfterCreate(urls []domain.Url) ([]domain.Url, error) {
	return urls, nil
}

func (i *UrlInter) BeforeUpdate(urls []domain.Url) ([]domain.Url, error) {
	for k := range urls {
		err := i.scopeModel(&urls[k])
		if err != nil {
			return nil, err
		}
	}
	return urls, nil
}

func (i *UrlInter) AfterUpdate(urls []domain.Url) ([]domain.Url, error) {
	return urls, nil
}

func (i *UrlInter) BeforeDelete(urls []domain.Url) ([]domain.Url, error) {
	return urls, nil
}

func (i *UrlInter) AfterDelete(urls []domain.Url) ([]domain.Url, error) {
	return urls, nil
}
