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

	err := role.model.Where("key = ?", key).First(&roleEntity).Error

	return roleEntity, err
}

func (role *RoleRepository) FindById(id uuid.UUID, trashed bool) (*domain.RoleEntity, error) {
	roleEntity := &domain.RoleEntity{}
	if trashed {
		role.model = role.model.Unscoped()
	}

	err := role.model.Where("id = ?", id).First(&roleEntity).Error

	return roleEntity, err
}

func (role *RoleRepository) FindByNameAndKey(name string, key string) (*domain.RoleEntity, error) {
	roleEntity := &domain.RoleEntity{}
	role.model.First(&roleEntity, "name = ? and key = ?", name, key)

	return roleEntity, nil
}

func (role *RoleRepository) Delete(id uuid.UUID) {
	role.model.Delete(&domain.RoleEntity{}, id)
}

func (role *RoleRepository) ForceDelete(id uuid.UUID) {
	roleEntity := &domain.RoleEntity{}
	role.model.Unscoped().Delete(&roleEntity, id)
}

func (role *RoleRepository) Update(id uuid.UUID, payload *domain.RoleEntity) error {
	err := role.model.Where("id = ?", id).Updates(&payload).Error
	return err
}

func (role *RoleRepository) Create(payload *domain.RoleEntity) error {
	err := role.model.Create(&payload).Error
	return err
}

func (user *RoleRepository) IsKeyExist(email string) bool {
	var count int64
	user.model.
		Where("key = ?", email).
		Count(&count)
	return count > 0
}

func (role *RoleRepository) IsKeyExistExceptRoleId(key string, id uuid.UUID) bool {
	var count int64
	role.model.
		Session(&gorm.Session{}).
		Where("key = ? AND id != ?", key, id).
		Count(&count)
	return count > 0
}
