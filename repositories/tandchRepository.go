package repositories

import (
	"errors"
	"gorm.io/gorm"
	"server/dtos/tandch"
	"server/interface/Repository"
	"server/models"
	"time"
)

type TAndChRepository struct {
	DB *gorm.DB
}

func (e *TAndChRepository) CreateTAndCh(createTAndChDto *tandch.CreateTandChDto) (*models.TandCh, error) {
	if err := e.DB.Table("tandch").Create(createTAndChDto).Error; err != nil {
		return nil, err
	}

	m := &models.TandCh{
		Name:   createTAndChDto.Name,
		RoomID: createTAndChDto.RoomID,
		Status: models.EquipmentStatus(createTAndChDto.Status),
	}
	return m, nil
}

func (e *TAndChRepository) UpdateTAndCh(TAndChId int, dto tandch.UpdateTandChDto) (*models.TandCh, error) {
	var existingTAnd models.TandCh
	if err := e.DB.Table("tandch").Where("id = ?", TAndChId).First(&existingTAnd).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"name":    dto.Name,
		"room_id": dto.RoomID,
		"status":  dto.Status,
	}

	if err := e.DB.Table("tandch").Where("id = ?", TAndChId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := e.DB.First(&existingTAnd, TAndChId).Error; err != nil {
		return nil, err
	}
	return &existingTAnd, nil
}

func (e *TAndChRepository) DeleteTAndCh(TAndChId int) error {
	result := e.DB.Table("tandch").Where("id = ?", TAndChId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (e *TAndChRepository) GetAllTAndChs() ([]*models.TandCh, error) {
	var tandchs []*models.TandCh
	if err := e.DB.Find(&tandchs).Error; err != nil {
		return nil, err
	}
	return tandchs, nil
}

func NewTAndChRepository(db *gorm.DB) Repository.TAndChRepositoryInterface {
	return &TAndChRepository{
		DB: db,
	}
}
