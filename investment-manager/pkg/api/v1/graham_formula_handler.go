package v1

import (
	"net/http"

	"strconv"

	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
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

type GrahamFormulaHandler struct {
	grahamFormulaService *service.GrahamFormulaService
}

func NewGrahamFormulaHandler(grahamFormulaService *service.GrahamFormulaService) *GrahamFormulaHandler {
	return &GrahamFormulaHandler{grahamFormulaService: grahamFormulaService}
}

// @Summary Get All Graham Formula for a code
// @Description Get All Graham Formula for a code
// @Tags graham-formula
// @Accept  json
// @Produce  json
// @Param codigo path int true "Codigo"
// @Success 200 {object} []response.GrahamFormulaResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/graham-formula/{codigo} [get]
func (h *GrahamFormulaHandler) GetGrahamFormulaByCodigo(c *gin.Context) {

	codigo := c.Param("codigo")
	data, err := h.grahamFormulaService.GetGrahamFormulaByCodigo(codigo)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// @Summary Get the most recente Graham Formula for a code
// @Description Get the most Graham Formula for a code
// @Tags graham-formula
// @Accept  json
// @Produce  json
// @Param codigo path int true "Codigo"
// @Success 200 {object} response.GrahamFormulaResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/graham-formula/{codigo}/last [get]
func (h *GrahamFormulaHandler) GetGrahamFormulaMoreRecentByCodigo(c *gin.Context) {
	var err error
	codigo := c.Param("codigo")

	data, err := h.grahamFormulaService.GetGrahamFormulaMoreRecentByCodigo(codigo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary Calculate graham formula
// @Description Calculate graham formula
// @Tags graham-formula
// @Accept  json
// @Produce  json
// @Param grahamFormula body request.GrahamFormulaRequest true "GrahamFormulaRequest info"
// @Success 201 {object} response.GrahamFormulaResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/graham-formula [post]
func (h *GrahamFormulaHandler) CreateGrahamFormula(c *gin.Context) {
	var request request.GrahamFormulaRequest
	var response response.GrahamFormulaResponse
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.grahamFormulaService.CreateGrahamFormula(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary Delete a Graham Formula
// @Description Delete a Graham Formula by ID
// @Tags graham-formula
// @Accept  json
// @Produce  json
// @Param id path int true "Graham Formula ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/graham-formula/{id} [delete]
func (h *GrahamFormulaHandler) DeleteGrahamFormula(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.grahamFormulaService.DeleteGrahamFormula(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
