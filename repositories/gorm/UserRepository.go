package gorm

import (
	"github.com/jinzhu/gorm"
	"nikan.dev/pronto/exceptions"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/internals/dependencies"
)

type userRepository struct {
	pool *gorm.DB
	deps dependencies.CommonDependencies
}

func (repository userRepository) Update(user entities.User) (entities.User, error) {
	q := repository.pool.Model(&user).Updates(user)
	if q.Error != nil {
		return user, q.Error
	}
	return user, nil
}

func NewUserRepository(deps dependencies.CommonDependencies, pool interface{}) userRepository {
	return userRepository{pool: pool.(*gorm.DB), deps: deps}
}

func (repository userRepository) Get(ID uint) (entities.User, error) {
	var user entities.User
	if err := repository.pool.First(&user,ID).Error; err != nil {
		return user, exceptions.UserNotFound
	}
	return user, nil
}

func (repository userRepository) GetByEmail(Email string) (entities.User, error) {
	var user entities.User
	if err := repository.pool.Where("email = ?", Email).First(&user); err != nil {
		return user,err.Error
	}
	return user, nil

}

func (repository userRepository) EmailExists(Email string) (bool, error) {
	userCount := 0
	repository.pool.Model(&entities.User{}).Where("email = ?", Email).Count(&userCount)
	return userCount > 0, nil
}
func (repository userRepository) Create(user entities.User) (entities.User, error) {
	if err := repository.pool.Create(&user).Error; err != nil {
		return user, exceptions.ServerError
	}
	return user, nil
}

