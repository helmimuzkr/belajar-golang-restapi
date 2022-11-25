package service

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/helmimuzkr/belajar-golang-restapi/model"
	"github.com/helmimuzkr/belajar-golang-restapi/repository"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, request *model.CreateCategoryRequest) (*model.CategoryResponse, error)
	UpdateCategory(ctx context.Context, request *model.UpdateCategoryRequest) (*model.CategoryResponse, error)
	DeleteCategory(ctx context.Context, requestID int) error
	GetAllCategory(ctx context.Context) ([]*model.CategoryResponse, error)
	GetCategory(ctx context.Context, requestID int) (*model.CategoryResponse, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
	validator    *validator.Validate
}

func NewCategoryService(validator *validator.Validate, categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		validator:    validator,
		categoryRepo: categoryRepo,
	}
}

// Create category
func (service *categoryService) CreateCategory(ctx context.Context, request *model.CreateCategoryRequest) (*model.CategoryResponse, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, fmt.Errorf("validation error - %v", err)
	}

	category := &model.Category{Name: request.Name}
	category, err = service.categoryRepo.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	categoryResponse := &model.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return categoryResponse, nil
}

// Update category
func (service *categoryService) UpdateCategory(ctx context.Context, request *model.UpdateCategoryRequest) (*model.CategoryResponse, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, fmt.Errorf("validation error - %v", err)
	}

	category := &model.Category{
		ID:   request.ID,
		Name: request.Name,
	}

	category, err = service.categoryRepo.GetCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	category, err = service.categoryRepo.UpdateCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	categoryResponse := &model.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return categoryResponse, nil
}

// Delete Category and return error
func (service *categoryService) DeleteCategory(ctx context.Context, requestID int) error {
	category := &model.Category{ID: requestID}
	err := service.categoryRepo.DeleteCategory(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

// Read all category, return category responses and an error
func (service *categoryService) GetAllCategory(ctx context.Context) ([]*model.CategoryResponse, error) {
	categories, err := service.categoryRepo.GetAllCategory(ctx)
	if err != nil {
		return nil, err
	}

	categoriesResponse := []*model.CategoryResponse{}
	for _, category := range categories {
		categoryResponse := &model.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		}
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return categoriesResponse, nil
}

// Read category by id
func (service *categoryService) GetCategory(ctx context.Context, requestID int) (*model.CategoryResponse, error) {
	category := &model.Category{ID: requestID}
	category, err := service.categoryRepo.GetCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	categoryResposne := &model.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return categoryResposne, nil
}
