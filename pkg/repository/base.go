package pkg

import (
	"math"

	"github.com/primeiro/pkg/pagination"
	"gorm.io/gorm"
)

type RepositoryBaseInterface[E interface{}] interface {
	BeginTx() *gorm.DB
	Create(unidade *E) (*E, error)
	CreateTx(unidade *E, tx *gorm.DB) error
	GetByID(id string) (*E, error)
	GetAll() ([]E, error)
	Update(unidade *E) (*E, error)
	Delete(id string) error
	GetPaginated(query *pagination.PaginationQuery) (*pagination.PaginationResponse[E], error)
}

type RepositoryBase[E interface{}] struct {
	Db               *gorm.DB
	SearchExpression string
}

func (r *RepositoryBase[E]) Create(entity *E) (*E, error) {
	err := r.Db.Create(entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *RepositoryBase[E]) GetByID(id string) (*E, error) {
	var entity E
	if err := r.Db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *RepositoryBase[E]) GetAll() ([]E, error) {
	var entities []E
	if err := r.Db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *RepositoryBase[E]) Update(entity *E) (*E, error) {
	err := r.Db.Updates(entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *RepositoryBase[E]) Delete(id string) error {
	var entity E
	if err := r.Db.First(&entity, "id = ?", id).Error; err != nil {
		return err
	}
	return r.Db.Delete(&entity).Error
}

func (r *RepositoryBase[E]) GetPaginated(query *pagination.PaginationQuery) (*pagination.PaginationResponse[E], error) {
	var entities []E

	find := r.Db.Order(query.Sort)

	filters := query.Filters
	for _, filter := range filters {
		switch filter.Action {
		case "eq":
			find = find.Where(filter.Field+" = ?", filter.Value)
		case "like":
			find = find.Where(filter.Field+" LIKE ?", "%"+filter.Value+"%")
		case "gt":
			find = find.Where(filter.Field+" > ?", filter.Value)
		case "lt":
			find = find.Where(filter.Field+" < ?", filter.Value)
		case "gte":
			find = find.Where(filter.Field+" >= ?", filter.Value)
		case "lte":
			find = find.Where(filter.Field+" <= ?", filter.Value)
		case "ne":
			find = find.Where(filter.Field+" != ?", filter.Value)
		}
	}

	if (r.SearchExpression != "") && (query.Search != "") {
		find = find.Where(r.SearchExpression, query.Search)

	}

	var totalRows int64
	findCount := find.Find(&entities)
	count := findCount.Count(&totalRows)
	if count.Error != nil {
		return nil, count.Error
	}

	find = find.Limit(query.Limit).Offset((query.Page - 1) * query.Limit)

	find = find.Find(&entities)
	if find.Error != nil {
		return nil, find.Error
	}

	paginationResponse := &pagination.PaginationResponse[E]{
		Rows: entities,
		Meta: pagination.PaginationMeta{
			TotalRows:   int(totalRows),
			FromRow:     ((query.Page - 1) * query.Limit) + 1,
			ToRow:       (query.Page * query.Limit),
			TotalPages:  int(math.Ceil(float64(totalRows) / float64(query.Limit-1))),
			PerPage:     query.Limit,
			CurrentPage: query.Page,
		},
	}
	return paginationResponse, nil

	/*find := r.Db.Where("nome = ?", "nome")



	if err := find.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil*/
}

func (r *RepositoryBase[E]) CreateTx(entity *E, tx *gorm.DB) error {
	return tx.Create(entity).Error
}

func (r *RepositoryBase[E]) BeginTx() *gorm.DB {
	return r.Db.Begin()
}
