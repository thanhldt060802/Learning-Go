package handler

import (
	"go-api-with-gin/dto"
	"go-api-with-gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	invoiceService service.InvoiceService
}

func NewInvoiceHandler(invoiceService service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{invoiceService}
}

func (invoiceHandler *InvoiceHandler) GetAllInvoices(context *gin.Context) {
	invoices, err := invoiceHandler.invoiceService.FindAllInvoices()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, invoices)
}

func (invoiceHandler *InvoiceHandler) GetInvoiceById(context *gin.Context) {
	invoiceId, _ := strconv.Atoi(context.Param("id"))

	invoice, err := invoiceHandler.invoiceService.FindInvoiceById(int64(invoiceId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, invoice)
}

func (invoiceHandler *InvoiceHandler) GetAllInvoicesByUserId(context *gin.Context) {
	userId, _ := strconv.Atoi(context.Param("userId"))

	invoices, err := invoiceHandler.invoiceService.FindAllInvoicesByUserId(int64(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, invoices)
}

func (invoiceHandler *InvoiceHandler) GetInvoiceByIdAndUserId(context *gin.Context) {
	invoiceId, _ := strconv.Atoi(context.Param("id"))
	userId, _ := strconv.Atoi(context.Param("userId"))

	invoice, err := invoiceHandler.invoiceService.FindInvoiceByIdAndUserId(int64(invoiceId), int64(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, invoice)
}

func (invoiceHandler *InvoiceHandler) CreateInvoice(context *gin.Context) {
	var newInvoice dto.CreateInvoiceDTO

	if err := context.ShouldBind(&newInvoice); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := invoiceHandler.invoiceService.CreateInvoice(newInvoice); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, "Create invoice success!")
}

func (invoiceHandler *InvoiceHandler) UpdateInvoice(context *gin.Context) {
	invoiceId, _ := strconv.Atoi(context.Param("id"))

	var updatingInvoice dto.UpdateInvoiceDTO

	if err := context.ShouldBind(&updatingInvoice); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := invoiceHandler.invoiceService.UpdateInvoice(int64(invoiceId), updatingInvoice); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, "Update invoice success!")
}

func (invoiceHandler *InvoiceHandler) DeleteInvoiceById(context *gin.Context) {
	invoiceId, _ := strconv.Atoi(context.Param("id"))

	if err := invoiceHandler.invoiceService.DeleteInvoiceById(int64(invoiceId)); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, "Delete invoice success!")
}
