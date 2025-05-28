package pkg

import "gorm.io/gorm"

type RepositoryBaseInterface[E interface{}] interface {
	Create(unidade *E) error
	GetByID(id string) (*E, error)
	GetAll() ([]E, error)
	Update(unidade *E) error
	Delete(id string) error
}

type RepositoryBase[E interface{}] struct {
	Db *gorm.DB
}

func (r *RepositoryBase[E]) Create(entity *E) error {
	return r.Db.Create(entity).Error

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

func (r *RepositoryBase[E]) Update(entity *E) error {
	return r.Db.Updates(entity).Error
}

func (r *RepositoryBase[E]) Delete(id string) error {
	var entity E
	if err := r.Db.First(&entity, "id = ?", id).Error; err != nil {
		return err
	}
	return r.Db.Delete(&entity).Error
}
