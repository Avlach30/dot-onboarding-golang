package repository

import (
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gorm.io/gorm"
)

type RoleRepository struct {
	model *gorm.DB
}

func NewRoleRepository(db *gorm.DB) domain.RoleRepository {
	return &RoleRepository{
		model: db.Model(&domain.RoleEntity{}),
	}
}

// FindByKey implements domain.RoleRepository.
func (role *RoleRepository) FindByKey(key string, trashed bool) (*domain.RoleEntity, error) {
	roleEntity := &domain.RoleEntity{}
	if trashed {
		role.model = role.model.Unscoped()
	}

	role.model.Where("key = ?", key).First(&roleEntity)

	return roleEntity, nil
}

func (role *RoleRepository) FindById(id uuid.UUID, trashed bool) (*domain.RoleEntity, error) {
	roleEntity := &domain.RoleEntity{}
	if trashed {
		role.model = role.model.Unscoped()
	}

	role.model.Where("id = ?", id).First(&roleEntity)

	return roleEntity, nil
}

func (role *RoleRepository) FindByNameAndKey(name string, key string) (*domain.RoleEntity, error) {

	roleEntity := &domain.RoleEntity{}
	role.model.First(&roleEntity, "name = ? and key = ?", name, key)

	return roleEntity, nil
}

func (role *RoleRepository) Delete(id uuid.UUID) {
	role.model.Where("id = ?", id).Delete(&domain.RoleEntity{})
}

func (role *RoleRepository) ForceDelete(id uuid.UUID) {
	roleEntity := &domain.RoleEntity{}
	role.model.Unscoped().Where("id = ?", id).Find(&roleEntity)
	role.model.Unscoped().Delete(&roleEntity)
}

func (role *RoleRepository) Update(id uuid.UUID, payload *domain.RoleEntity) {
	role.model.Where("id = ?", id).Updates(&payload)
}

func (role *RoleRepository) Create(payload *domain.RoleEntity) error {
	role.model.Create(payload)

	return nil
}
