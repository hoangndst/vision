package models

import (
	"context"
	"github.com/hoangndst/vision/domain/entity"
	"github.com/hoangndst/vision/domain/repository"
	"gorm.io/gorm"
)

var _ repository.UserRepository = &userRepository{}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u userRepository) Create(ctx context.Context, dataEntity *entity.User) error {
	err := dataEntity.Validate()
	if err != nil {
		return err
	}

	var dataModel UserModel
	err = dataModel.FromEntity(dataEntity)
	if err != nil {
		return err
	}

	return u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&dataModel).Error; err != nil {
			return err
		}
		dataEntity.ID = dataModel.ID
		return nil
	})
}

func (u userRepository) Delete(ctx context.Context, id uint) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		var dataModel UserModel
		if err := tx.WithContext(ctx).First(&dataModel, id).Error; err != nil {
			return err
		}
		return tx.WithContext(ctx).Unscoped().Delete(&dataModel).Error
	})
}

func (u userRepository) Update(ctx context.Context, user *entity.User) error {
	var dataModel UserModel
	if err := dataModel.FromEntity(user); err != nil {
		return err
	}
	if err := u.db.WithContext(ctx).Updates(&dataModel).Error; err != nil {
		return err
	}
	return nil
}

func (u userRepository) Get(ctx context.Context, id uint) (*entity.User, error) {
	var dataModel UserModel
	if err := u.db.WithContext(ctx).First(&dataModel, id).Error; err != nil {
		return nil, err
	}
	return dataModel.ToEntity()
}

func (u userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var dataModel UserModel
	if err := u.db.WithContext(ctx).Where("username = ?", username).First(&dataModel).Error; err != nil {
		return nil, err
	}
	return dataModel.ToEntity()
}

func (u userRepository) List(ctx context.Context) ([]*entity.User, error) {
	var dataModels []UserModel
	userEntityList := make([]*entity.User, 0)
	if err := u.db.WithContext(ctx).Find(&dataModels).Error; err != nil {
		return nil, err
	}
	for _, user := range dataModels {
		userEntity, err := user.ToEntity()
		if err != nil {
			return nil, err
		}
		userEntityList = append(userEntityList, userEntity)
	}
	return userEntityList, nil
}
