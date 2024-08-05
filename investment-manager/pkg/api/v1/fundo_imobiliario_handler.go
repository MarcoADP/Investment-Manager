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

type FundoImobiliarioHandler struct {
	fundoImobiliarioService *service.FundoImobiliarioService
}

func NewFundoImobiliarioHandler(fundoImobiliarioService *service.FundoImobiliarioService) *FundoImobiliarioHandler {
	return &FundoImobiliarioHandler{fundoImobiliarioService: fundoImobiliarioService}
}

// @Summary Get All Fundos Imobiliarios
// @Description Get All Fundos Imobiliarios
// @Tags fundos-imobiliarios
// @Accept  json
// @Produce  json
// @Success 200 {object} []response.FundoImobiliarioResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/fundos-imobiliarios [get]
func (h *FundoImobiliarioHandler) GetFundosImobiliarios(c *gin.Context) {
	fundosImobiliarios, err := h.fundoImobiliarioService.GetAllFundosImobiliarios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fundosImobiliarios)
}

// @Summary Get a Fundo Imobiliario
// @Description Get a Fundo Imobiliario by ID
// @Tags fundos-imobiliarios
// @Accept  json
// @Produce  json
// @Param id path int true "FundoImobiliario ID"
// @Success 200 {object} response.FundoImobiliarioResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/fundos-imobiliarios/{id} [get]
func (h *FundoImobiliarioHandler) GetFundoImobiliario(c *gin.Context) {
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var fundoImobiliario response.FundoImobiliarioResponse
	fundoImobiliario, err = h.fundoImobiliarioService.GetFundoImobiliarioByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fundoImobiliario)
}

// @Summary Create a Fundo Imobiliario
// @Description Create a new Fundo Imobiliario
// @Tags fundos-imobiliarios
// @Accept  json
// @Produce  json
// @Param fundoImobiliario body request.FundoImobiliarioRequest true "FundoImobiliarioRequest info"
// @Success 201 {object} response.FundoImobiliarioResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/fundos-imobiliarios [post]
func (h *FundoImobiliarioHandler) CreateFundoImobiliario(c *gin.Context) {
	var fundoImobiliario request.FundoImobiliarioRequest
	if err := c.ShouldBindJSON(&fundoImobiliario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdFundoImobiliario, err := h.fundoImobiliarioService.CreateFundoImobiliario(fundoImobiliario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdFundoImobiliario)
}

// @Summary Update a Fundo Imobiliario
// @Description Update a Fundo Imobiliario by ID
// @Tags fundos-imobiliarios
// @Accept  json
// @Produce  json
// @Param id path int true "FundoImobiliario ID"
// @Param fundoImobiliario body request.FundoImobiliarioRequest true "FundoImobiliarioRequest info"
// @Success 200 {object} response.FundoImobiliarioResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/fundos-imobiliarios/{id} [put]
func (h *FundoImobiliarioHandler) UpdateFundoImobiliario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var fundoImobiliario request.FundoImobiliarioRequest
	if err := c.ShouldBindJSON(&fundoImobiliario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedFundoImobiliario, err := h.fundoImobiliarioService.UpdateFundoImobiliario(uint(id), fundoImobiliario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedFundoImobiliario)
}

// @Summary Delete a Fundo Imobiliario
// @Description Delete a Fundo Imobiliario by ID
// @Tags fundos-imobiliarios
// @Accept  json
// @Produce  json
// @Param id path int true "FundoImobiliario ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/fundos-imobiliarios/{id} [delete]
func (h *FundoImobiliarioHandler) DeleteFundoImobiliario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.fundoImobiliarioService.DeleteFundoImobiliario(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
