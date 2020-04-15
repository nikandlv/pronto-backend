package mysql

import (
	"github.com/jinzhu/gorm"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/exceptions"
	"nikan.dev/pronto/internals/dependencies"
)

type categoryRepository struct {
	pool *gorm.DB
	deps dependencies.CommonDependencies
}

func NewCategoryRepository(deps dependencies.CommonDependencies, pool interface{}) categoryRepository {
	return categoryRepository{pool: pool.(*gorm.DB), deps: deps}
}

func (repository categoryRepository) List() ([]entities.Category, error) {
	var categories []entities.Category
	repository.pool.Find(&categories)
	return categories, nil
}

func (repository categoryRepository) Create(category entities.Category) (entities.Category, error) {
	if err := repository.pool.Create(&category).Error; err != nil {
		return category, exceptions.ServerError
	}
	return category, nil

}


func (repository categoryRepository) Get(ID uint) (entities.Category, error) {
	var category entities.Category
	if err := repository.pool.First(&category,ID).Error; err != nil {
		return category, exceptions.UserNotFound
	}
	return category, nil
}

func (repository categoryRepository) SlugExists(Slug string) (bool, error) {
	categoryCount := 0
	repository.pool.Model(&entities.Category{}).Where("slug = ?", Slug).Count(&categoryCount)
	return categoryCount > 0, nil
}

func (repository categoryRepository) Exists(ID uint) (bool, error) {
	categoryCount := 0
	repository.pool.Model(&entities.Category{}).Where("id = ?", ID).Count(&categoryCount)
	return categoryCount > 0, nil
}

