package api

import (
	"github.com/gin-gonic/gin"
	"github.com/valverdethiago/volume-take-home-assignment/domain"
	"github.com/valverdethiago/volume-take-home-assignment/server"
	"net/http"
)

const defaultCalculateRoute = "/calculate"

type FlightMapApiHandlerImpl struct {
	service domain.FlightPathService
}

func NewFlightMapApiHandlerImpl(service domain.FlightPathService) server.Controller {
	return &FlightMapApiHandlerImpl{
		service: service,
	}
}

func (f FlightMapApiHandlerImpl) ConfigureRoutes(router *gin.Engine) {
	router.POST(defaultCalculateRoute, f.Calculate)
}

func (f FlightMapApiHandlerImpl) Calculate(ctx *gin.Context) {
	payload, err := parsePayload(ctx)
	if err != nil {
		return
	}
	result, err := f.service.CalculateFlightPath(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func parsePayload(ctx *gin.Context) (*[]*domain.Flight, error) {
	var req *[]*domain.Flight
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return nil, err
	}
	return req, err
}
