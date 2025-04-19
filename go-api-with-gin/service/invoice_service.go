package service

import (
	"go-api-with-gin/dto"
	"go-api-with-gin/model"
	"go-api-with-gin/repository"
)

type invoiceService struct {
	invoiceRepository repository.InvoiceRepository
}

type InvoiceService interface {
	FindAllInvoices() ([]model.Invoice, error)
	FindInvoiceById(id int64) (*model.Invoice, error)
	FindAllInvoicesByUserId(userId int64) ([]model.Invoice, error)
	FindInvoiceByIdAndUserId(id int64, userId int64) (*model.Invoice, error)
	CreateInvoice(invoice dto.CreateInvoiceDTO) error
	UpdateInvoice(id int64, invoice dto.UpdateInvoiceDTO) error
	DeleteInvoiceById(id int64) error
}

func NewInvoiceService(invoiceRepository repository.InvoiceRepository) InvoiceService {
	return &invoiceService{invoiceRepository}
}

func (invoiceRepository *invoiceService) FindAllInvoices() ([]model.Invoice, error) {
	return invoiceRepository.invoiceRepository.FindAll()
}

func (invoiceRepository *invoiceService) FindInvoiceById(id int64) (*model.Invoice, error) {
	return invoiceRepository.invoiceRepository.FindById(id)
}

func (invoiceRepository *invoiceService) FindAllInvoicesByUserId(userId int64) ([]model.Invoice, error) {
	return invoiceRepository.invoiceRepository.FindAllByUserId(userId)
}

func (invoiceRepository *invoiceService) FindInvoiceByIdAndUserId(id int64, userId int64) (*model.Invoice, error) {
	return invoiceRepository.invoiceRepository.FindByIdAndInvoiceId(id, userId)
}

func (invoiceRepository *invoiceService) CreateInvoice(newInvoice dto.CreateInvoiceDTO) error {
	invoice := model.Invoice{
		UserId:      newInvoice.UserId,
		TotalAmount: newInvoice.TotalAmount,
		Status:      newInvoice.Status,
	}
	return invoiceRepository.invoiceRepository.Create(&invoice)
}

func (invoiceRepository *invoiceService) UpdateInvoice(id int64, updatingInvoice dto.UpdateInvoiceDTO) error {
	invoice := model.Invoice{
		Status: updatingInvoice.Status,
	}
	return invoiceRepository.invoiceRepository.Update(id, &invoice)
}

func (invoiceRepository *invoiceService) DeleteInvoiceById(id int64) error {
	return invoiceRepository.invoiceRepository.DeleteById(id)
}
