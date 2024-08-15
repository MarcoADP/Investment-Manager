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

type CarteiraHandler struct {
	carteiraService *service.CarteiraService
}

func NewCarteiraHandler(carteiraService *service.CarteiraService) *CarteiraHandler {
	return &CarteiraHandler{carteiraService: carteiraService}
}

// @Summary Get All Carteiras
// @Description Get All Carteiras
// @Tags carteiras
// @Accept  json
// @Produce  json
// @Success 200 {object} []response.CarteiraResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/carteiras [get]
func (h *CarteiraHandler) GetCarteiras(c *gin.Context) {
	carteiras, err := h.carteiraService.GetAllCarteiras()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, carteiras)
}

// @Summary Get a Carteira
// @Description Get a Carteira by ID
// @Tags carteiras
// @Accept  json
// @Produce  json
// @Param id path int true "Carteira ID"
// @Success 200 {object} response.CarteiraResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/carteiras/{id} [get]
func (h *CarteiraHandler) GetCarteira(c *gin.Context) {
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var carteira response.CarteiraResponse
	carteira, err = h.carteiraService.GetCarteiraByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, carteira)
}

// @Summary Create a carteira
// @Description Create a new carteira
// @Tags carteiras
// @Accept  json
// @Produce  json
// @Param carteira body request.CarteiraRequest true "CarteiraRequest info"
// @Success 201 {object} response.CarteiraResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/carteiras [post]
func (h *CarteiraHandler) CreateCarteira(c *gin.Context) {
	var carteira request.CarteiraRequest
	if err := c.ShouldBindJSON(&carteira); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCarteira, err := h.carteiraService.CreateCarteira(carteira)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdCarteira)
}

// @Summary Update a carteira
// @Description Update a carteira by ID
// @Tags carteiras
// @Accept  json
// @Produce  json
// @Param id path int true "Carteira ID"
// @Param carteira body request.CarteiraRequest true "CarteiraRequest info"
// @Success 200 {object} response.CarteiraResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/carteiras/{id} [put]
func (h *CarteiraHandler) UpdateCarteira(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var carteira request.CarteiraRequest
	if err := c.ShouldBindJSON(&carteira); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCarteira, err := h.carteiraService.UpdateCarteira(uint(id), carteira)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCarteira)
}

// @Summary Delete a carteira
// @Description Delete a carteira by ID
// @Tags carteiras
// @Accept  json
// @Produce  json
// @Param id path int true "Carteira ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/carteiras/{id} [delete]
func (h *CarteiraHandler) DeleteCarteira(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.carteiraService.DeleteCarteira(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @Summary Add ativo na carteira
// @Description Add ativo na carteira
// @Tags carteiras
// @Accept  json
// @Produce  json
// @Param id path int true "Carteira ID"
// @Param carteira body request.CarteiraAtivoRequest true "CarteiraAtivoRequest info"
// @Success 200 {object} response.CarteiraAtivoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/carteiras/{id}/ativo [post]
func (h *CarteiraHandler) AddAtivoCarteira(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var ativo request.CarteiraAtivoRequest
	if err := c.ShouldBindJSON(&ativo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ativoAdded, err := h.carteiraService.AddAtivo(uint(id), ativo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ativoAdded)
}

// @Summary Delete a carteira
// @Description Delete a carteira by ID
// @Tags carteiras
// @Accept  json
// @Produce  json
// @Param id path int true "Carteira ID"
// @Param codigo path int true "Ativo Codigo"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/carteiras/{id}/ativo/{codigo} [delete]
func (h *CarteiraHandler) DeleteAtivo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.carteiraService.RemoverAtivo(uint(id), c.Param("codigo"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
