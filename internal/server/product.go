package server

import (
	"github.com/amitesh080295/GoMicroserviceProject/internal/dberrors"
	"github.com/amitesh080295/GoMicroserviceProject/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *EchoServer) GetAllProducts(ctx echo.Context) error {
	vendorID := ctx.QueryParam("vendorId")
	products, err := s.DB.GetAllProducts(ctx.Request().Context(), vendorID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, products)
}

func (s *EchoServer) AddProduct(ctx echo.Context) error {
	product := new(models.Product)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	product, err := s.DB.AddProduct(ctx.Request().Context(), product)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, product)
}
