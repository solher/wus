package ressources

import (
	"database/sql"
	"strings"

	"github.com/solher/wus/domain"
	"github.com/solher/zest/interfaces"
	"github.com/solher/zest/internalerrors"
	"github.com/solher/zest/usecases"
	"github.com/solher/zest/utils"
)

func init() {
	usecases.DependencyDirectory.Register(NewUrlRepo)
}

type UrlRepo struct {
	store interfaces.AbstractGormStore
}

func NewUrlRepo(store interfaces.AbstractGormStore) *UrlRepo {
	return &UrlRepo{store: store}
}

func (r *UrlRepo) Create(urls []domain.Url) ([]domain.Url, error) {
	db := r.store.GetDB()
	transaction := db.Begin()

	for i, url := range urls {
		err := db.Create(&url).Error
		if err != nil {
			transaction.Rollback()

			if strings.Contains(err.Error(), "constraint") {
				return nil, internalerrors.NewViolatedConstraint(err.Error())
			}

			return nil, internalerrors.DatabaseError
		}

		urls[i] = url
	}

	transaction.Commit()
	return urls, nil
}

func (r *UrlRepo) CreateOne(url *domain.Url) (*domain.Url, error) {
	urls, err := r.Create([]domain.Url{*url})
	if err != nil {
		return nil, err
	}

	return &urls[0], nil
}

func (r *UrlRepo) Find(context usecases.QueryContext) ([]domain.Url, error) {
	query, err := r.store.BuildQuery(context.Filter, context.OwnerRelations)
	if err != nil {
		return nil, internalerrors.DatabaseError
	}

	var urls []domain.Url

	err = query.Find(&urls).Error
	if err != nil {
		return nil, internalerrors.DatabaseError
	}

	if len(urls) == 0 {
		urls = []domain.Url{}
	}

	return urls, nil
}

func (r *UrlRepo) FindByID(id int, context usecases.QueryContext) (*domain.Url, error) {
	query, err := r.store.BuildQuery(context.Filter, context.OwnerRelations)
	if err != nil {
		return nil, internalerrors.DatabaseError
	}

	url := domain.Url{}

	err = query.Where(utils.ToDBName("urls")+".id = ?", id).First(&url).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, internalerrors.NotFound
		}

		return nil, internalerrors.DatabaseError
	}

	return &url, nil
}

func (r *UrlRepo) Update(urls []domain.Url, context usecases.QueryContext) ([]domain.Url, error) {
	db := r.store.GetDB()
	transaction := db.Begin()

	query, err := r.store.BuildQuery(context.Filter, context.OwnerRelations)
	if err != nil {
		return nil, internalerrors.DatabaseError
	}

	for _, url := range urls {
		queryCopy := *query

		dbName := utils.ToDBName("urls")

		err = queryCopy.Where(dbName+".id = ?", url.ID).First(&domain.Url{}).Error
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				return nil, internalerrors.NotFound
			}

			return nil, internalerrors.DatabaseError
		}

		err = r.store.GetDB().Where(dbName+".id = ?", url.ID).Model(&domain.Url{}).Updates(&url).Error
		if err != nil {
			if strings.Contains(err.Error(), "constraint") {
				return nil, internalerrors.NewViolatedConstraint(err.Error())
			}

			return nil, internalerrors.DatabaseError
		}
	}

	transaction.Commit()
	return urls, nil
}

func (r *UrlRepo) UpdateByID(id int, url *domain.Url,
	context usecases.QueryContext) (*domain.Url, error) {

	query, err := r.store.BuildQuery(context.Filter, context.OwnerRelations)
	if err != nil {
		return nil, internalerrors.DatabaseError
	}

	dbName := utils.ToDBName("urls")

	err = query.Where(dbName+".id = ?", id).First(&domain.Url{}).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, internalerrors.NotFound
		}

		return nil, internalerrors.DatabaseError
	}

	err = r.store.GetDB().Where(dbName+".id = ?", id).Model(&domain.Url{}).Updates(&url).Error
	if err != nil {
		if strings.Contains(err.Error(), "constraint") {
			return nil, internalerrors.NewViolatedConstraint(err.Error())
		}

		return nil, internalerrors.DatabaseError
	}

	return url, nil
}

func (r *UrlRepo) DeleteAll(context usecases.QueryContext) error {
	query, err := r.store.BuildQuery(context.Filter, context.OwnerRelations)
	if err != nil {
		return internalerrors.DatabaseError
	}

	urls := []domain.Url{}
	err = query.Find(&urls).Error
	if err != nil {
		return internalerrors.DatabaseError
	}

	urlIDs := []int{}
	for _, url := range urls {
		urlIDs = append(urlIDs, url.ID)
	}

	err = r.store.GetDB().Delete(&urls, utils.ToDBName("urls")+".id IN (?)", urlIDs).Error
	if err != nil {
		return internalerrors.DatabaseError
	}

	return nil
}

func (r *UrlRepo) DeleteByID(id int, context usecases.QueryContext) error {
	query, err := r.store.BuildQuery(context.Filter, context.OwnerRelations)
	if err != nil {
		return internalerrors.DatabaseError
	}

	url := &domain.Url{}

	err = query.Where(utils.ToDBName("urls")+".id = ?", id).First(&url).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return internalerrors.NotFound
		}

		return internalerrors.DatabaseError
	}

	err = r.store.GetDB().Where(utils.ToDBName("urls")+".id = ?", url.ID).Delete(domain.Url{}).Error
	if err != nil {
		return internalerrors.DatabaseError
	}

	return nil
}

func (r *UrlRepo) Raw(query string, values ...interface{}) (*sql.Rows, error) {
	db := r.store.GetDB()

	rows, err := db.Raw(query, values...).Rows()
	if err != nil {
		if strings.Contains(err.Error(), "constraint") {
			return nil, internalerrors.NewViolatedConstraint(err.Error())
		}

		return nil, internalerrors.DatabaseError
	}

	return rows, nil
}
