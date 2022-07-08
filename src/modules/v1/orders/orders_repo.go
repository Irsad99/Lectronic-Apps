package orders

import (
	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindAll() (*models.Orders, error) {
	var orders *models.Orders

	// err := r.db.Order("id desc").Preload("Images", "vehicle_images.is_primary = true").Find(&orders).Error
	err := r.db.Order("id desc").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *repository) FindByID(id int) (*models.Order, error) {
	var order *models.Order

	err := r.db.First(&order, id).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *repository) FindByUserID(id int) (*models.Orders, error) {
	var order *models.Orders

	err := r.db.Where("user_id = ?", id).Find(&order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *repository) Save(order *models.Order) (*models.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *repository) Update(order *models.Order) (*models.Order, error) {
	err := r.db.Save(&order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *repository) Delete(id int) error {
	var order *models.Order

	err := r.db.Where("id = ?", id).Delete(&order).Error
	if err != nil {
		return err
	}

	return nil
}
