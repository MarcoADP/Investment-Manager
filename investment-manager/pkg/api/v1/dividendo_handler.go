package v1

import (
	"net/http"
	"strconv"

	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/service"
	"github.com/gin-gonic/gin"
)

// @title Your Project API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @basepath /api/v1

type DividendoHandler struct {
	dividendoService *service.DividendoService
}

func NewDividendoHandler(dividendoService *service.DividendoService) *DividendoHandler {
	return &DividendoHandler{dividendoService: dividendoService}
}

// @Summary Get All Dividendo
// @Description Get All Dividendo
// @Tags dividendos
// @Accept  json
// @Produce  json
// @Success 200 {object} []response.DividendoResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/dividendos [get]
func (h *DividendoHandler) GetDividendos(c *gin.Context) {
	dividendos, err := h.dividendoService.GetAllDividendos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dividendos)
}

// @Summary Get dividendos pelo codigo
// @Description Get a dividendos by Codigo
// @Tags dividendos
// @Accept  json
// @Produce  json
// @Param codigo path int true "Dividendo Codigo"
// @Success 200 {object} []response.CotacaoHistoricoResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/dividendos/{codigo} [get]
func (h *DividendoHandler) GetDividendosByCodigo(c *gin.Context) {
	codigo := c.Param("codigo")

	dividendos, err := h.dividendoService.GetDividendosByCodigo(codigo)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dividendos)
}

// @Summary Create a dividendo
// @Description Create a new dividendo
// @Tags dividendos
// @Accept  json
// @Produce  json
// @Param dividendo body request.DividendoRequest true "DividendoRequest info"
// @Success 201 {object} response.DividendoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/dividendos [post]
func (h *DividendoHandler) CreateDividendo(c *gin.Context) {
	var dividendo request.DividendoRequest
	if err := c.ShouldBindJSON(&dividendo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdDividendo, err := h.dividendoService.CreateDividendo(dividendo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdDividendo)
}

// @Summary Delete a dividendo
// @Description Delete a dividendo by ID
// @Tags dividendos
// @Accept  json
// @Produce  json
// @Param id path int true "Dividendo ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/dividendos/{id} [delete]
func (h *DividendoHandler) DeleteDividendo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.dividendoService.DeleteDividendo(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
