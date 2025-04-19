package repository

import (
	"go-api-with-gin/model"

	"gorm.io/gorm"
)

type invoiceRepository struct {
	db *gorm.DB
}

type InvoiceRepository interface {
	FindAll() ([]model.Invoice, error)
	FindById(id int64) (*model.Invoice, error)
	FindAllByUserId(userId int64) ([]model.Invoice, error)
	FindByIdAndInvoiceId(id int64, invoiceId int64) (*model.Invoice, error)
	Create(invoice *model.Invoice) error
	Update(id int64, invoice *model.Invoice) error
	DeleteById(id int64) error
}

func NewInvoiceRepository(db *gorm.DB) InvoiceRepository {
	return &invoiceRepository{db}
}

func (invoiceRepository *invoiceRepository) FindAll() ([]model.Invoice, error) {
	var invoices []model.Invoice
	err := invoiceRepository.db.Find(&invoices).Error
	return invoices, err
}

func (invoiceRepository *invoiceRepository) FindById(id int64) (*model.Invoice, error) {
	var invoice model.Invoice
	err := invoiceRepository.db.Where("invoice_id = ?", id).First(&invoice).Error
	return &invoice, err
}

func (invoiceRepository *invoiceRepository) FindAllByUserId(userId int64) ([]model.Invoice, error) {
	var invoices []model.Invoice
	err := invoiceRepository.db.Where("user_id = ?", userId).Find(&invoices).Error
	return invoices, err
}

func (invoiceRepository *invoiceRepository) FindByIdAndInvoiceId(id int64, userId int64) (*model.Invoice, error) {
	var invoice model.Invoice
	err := invoiceRepository.db.Where("invoice_id = ?", id).Where("user_id = ?", userId).First(&invoice).Error
	return &invoice, err
}

func (invoiceRepository *invoiceRepository) Create(invoice *model.Invoice) error {
	return invoiceRepository.db.Create(invoice).Error
}

func (invoiceRepository *invoiceRepository) Update(id int64, invoice *model.Invoice) error {
	return invoiceRepository.db.Where("invoice_id = ?", id).Updates(invoice).Error
}

func (invoiceRepository *invoiceRepository) DeleteById(id int64) error {
	return invoiceRepository.db.Where("invoice_id = ?", id).Delete(&model.Invoice{}, id).Error
}
