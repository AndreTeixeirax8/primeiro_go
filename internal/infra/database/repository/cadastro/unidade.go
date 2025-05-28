package repository

import (
	"github.com/primeiro/internal/modules/cadastro/domain/entity"
	"gorm.io/gorm"
)

type UnidadeRepository struct {
	Db *gorm.DB
}

func NewUnidadeRepository(db *gorm.DB) *UnidadeRepository {
	return &UnidadeRepository{
		Db: db,
	}
}

func (r *UnidadeRepository) Create(unidade *entity.Unidade) (*entity.Unidade, error) {
	err := r.Db.Create(unidade).Error
	if err != nil {
		return nil, err
	}
	return unidade, nil
}

func (r *UnidadeRepository) GetByID(id string) (*entity.Unidade, error) {
	var unidade entity.Unidade
	err := r.Db.First(&unidade, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &unidade, nil
}

func (r *UnidadeRepository) GetAll() ([]entity.Unidade, error) {
	var unidades []entity.Unidade
	err := r.Db.Find(&unidades).Error
	if err != nil {
		return nil, err
	}
	return unidades, nil
}

func (r *UnidadeRepository) Update(unidade *entity.Unidade) (*entity.Unidade, error) {
	err := r.Db.Updates(unidade).Error
	if err != nil {
		return nil, err
	}
	return unidade, nil
}

func (r *UnidadeRepository) Delete(id string) error {

	return r.Db.Delete(&entity.Unidade{}, "id = ?", id).Error
}
