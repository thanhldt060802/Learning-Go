package routes

import (
	"go-api-with-gin/handler"
	"go-api-with-gin/repository"
	"go-api-with-gin/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// Init repository
	userRepo := repository.NewUserRepository(db)
	invoiceRepo := repository.NewInvoiceRepository(db)

	// Init service
	userService := service.NewUserService(userRepo)
	invoiceService := service.NewInvoiceService(invoiceRepo)

	// Init handler
	userHandler := handler.NewUserHandler(userService)
	invoiceHandler := handler.NewInvoiceHandler(invoiceService)

	// Routes
	api := r.Group("/api")
	{
		userRoute := api.Group("/users")
		{
			// User
			userRoute.GET("/", userHandler.GetAllUsers)
			userRoute.GET("/id/:id", userHandler.GetUserById)
			userRoute.GET("/username/:username", userHandler.GetUserByUsername)
			userRoute.POST("/", userHandler.CreateUser)
			userRoute.PUT("/id/:id", userHandler.UpdateUser)
			userRoute.DELETE("/id/:id", userHandler.DeleteUserById)
		}

		invoiceRoute := api.Group("/invoices")
		{
			// Invoice
			invoiceRoute.GET("/", invoiceHandler.GetAllInvoices)
			invoiceRoute.GET("/id/:id", invoiceHandler.GetInvoiceById)
			invoiceRoute.GET("/user-id/:userId", invoiceHandler.GetAllInvoicesByUserId)
			invoiceRoute.GET("/id/:id/user-id/:userId", invoiceHandler.GetInvoiceByIdAndUserId)
			invoiceRoute.POST("/", invoiceHandler.CreateInvoice)
			invoiceRoute.PUT("/id/:id", invoiceHandler.UpdateInvoice)
			invoiceRoute.DELETE("/id/:id", invoiceHandler.DeleteInvoiceById)
		}
	}
}
