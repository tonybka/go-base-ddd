package account

import (
	"github.com/google/uuid"
	customgorm "github.com/tonybka/go-base-ddd/infrastructure/custom_gorm"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (repo *AccountRepository) Create(dataModel AccountModel) error {
	if result := repo.db.Create(&dataModel); result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *AccountRepository) FindById(id uuid.UUID) (AccountModel, error) {
	var dataModel AccountModel

	if result := repo.db.Where("id = ?", customgorm.CustomTypeUUIDv1FromString(id.String())).First(&dataModel); result.Error != nil {
		return AccountModel{}, result.Error
	}

	return dataModel, nil
}

func (repo *AccountRepository) GetAll() ([]AccountModel, error) {
	var dataModels []AccountModel

	if result := repo.db.Find(&dataModels); result.Error != nil {
		return nil, result.Error
	}

	return dataModels, nil
}
