package server

import (
	"github.com/amitesh080295/GoMicroserviceProject/internal/dberrors"
	"github.com/amitesh080295/GoMicroserviceProject/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *EchoServer) GetAllServices(ctx echo.Context) error {
	services, err := s.DB.GetAllServices(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, services)
}

func (s *EchoServer) AddService(ctx echo.Context) error {
	service := new(models.Service)
	if err := ctx.Bind(service); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	service, err := s.DB.AddService(ctx.Request().Context(), service)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, service)
}

func (s *EchoServer) GetServiceById(ctx echo.Context) error {
	ID := ctx.Param("id")
	service, err := s.DB.GetServiceById(ctx.Request().Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, service)
}

func (s *EchoServer) UpdateService(ctx echo.Context) error {
	ID := ctx.Param("id")
	service := new(models.Service)
	if err := ctx.Bind(service); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	if ID != service.ServiceID {
		return ctx.JSON(http.StatusBadRequest, "ID on path doesn't match ID on body")
	}
	service, err := s.DB.UpdateService(ctx.Request().Context(), service)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, service)
}

func (s *EchoServer) DeleteService(ctx echo.Context) error {
	ID := ctx.Param("id")
	err := s.DB.DeleteService(ctx.Request().Context(), ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusResetContent)
}
