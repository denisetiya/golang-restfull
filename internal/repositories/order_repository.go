package repositories

import (
	"rest-api/internal/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("User").Preload("OrderItems").Preload("OrderItems.Product").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) GetAll(offset, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	err := r.db.Model(&models.Order{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("User").Preload("OrderItems").Preload("OrderItems.Product").Offset(offset).Limit(limit).Find(&orders).Error
	return orders, total, err
}

func (r *OrderRepository) GetByUserID(userID uint, offset, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	err := r.db.Model(&models.Order{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("User").Preload("OrderItems").Preload("OrderItems.Product").Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&orders).Error
	return orders, total, err
}

func (r *OrderRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&models.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (r *OrderRepository) Delete(id uint) error {
	return r.db.Delete(&models.Order{}, id).Error
}
