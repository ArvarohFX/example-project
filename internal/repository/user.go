package repository

import (
	dao "github.com/ArvarohFX/example-project/internal/models/dao/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	Get() error
}

type userRepository struct {
	DB *gorm.DB
}

func (u *userRepository) Get() error {
	var users []*dao.User
	if result := u.DB.Debug().
		Where("last_name = ?", "Zotov").
		Find(&users); result.Error != nil {
		return result.Error
	}

	return nil
}

func isUnique(slice []string) {
	var result []string
	checker := map[string]struct{}{}
	for _, el := range slice {
		if _, ok := checker[el]; !ok {
			result = append(result, el)
			checker[el] = struct{}{}
		}
		continue
	}
}
