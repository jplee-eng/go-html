package db

import (
	"errors"
	"fmt"
	"go-html/inits"
	"go-html/models"

	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

func AddUser(user models.User) error {
	db, err := inits.GetDB()
	if err != nil {
		return fmt.Errorf("database connection error: %w", err)
	}

	if err := validateUser(user); err != nil {
		return fmt.Errorf("user validation failed: %w", err)
	}

	result := db.Create(&user)
	if result.Error != nil {
		return fmt.Errorf("failed to create user: %w", result.Error)
	}

	return nil
}

func RemoveUser(userID uint) error {
	if userID == 0 {
		return fmt.Errorf("invalid user ID")
	}

	db, err := inits.GetDB()
	if err != nil {
		return fmt.Errorf("database connection error: %w", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to start transaction: %w", tx.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user models.User
	if err := tx.First(&user, userID).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return fmt.Errorf("failed to find user: %w", err)
	}

	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func GetUser(userID uint) (*models.User, error) {
	if userID == 0 {
		return nil, fmt.Errorf("invalid user ID")
	}

	db, err := inits.GetDB()
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

func validateUser(user models.User) error {
	if user.Name == "" {
		return fmt.Errorf("user name is required")
	}
	return nil
}
