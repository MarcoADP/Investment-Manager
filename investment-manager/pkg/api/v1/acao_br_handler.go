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

type AcaoBrHandler struct {
	acaoBrService *service.AcaoBrService
}

func NewAcaoBrHandler(acaoBrService *service.AcaoBrService) *AcaoBrHandler {
	return &AcaoBrHandler{acaoBrService: acaoBrService}
}

type ErrorResponse struct {
	Error string `json:"error" example:"Bad Request"`
}

// @Summary Get All Acão BR
// @Description Get All Ação BR
// @Tags acoes
// @Accept  json
// @Produce  json
// @Success 200 {object} []response.AcaoBrResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/acoes [get]
func (h *AcaoBrHandler) GetAcoesBr(c *gin.Context) {
	acoes, err := h.acaoBrService.GetAllAcaoBrs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, acoes)
}

// @Summary Get a Acão BR
// @Description Get a Ação BR by ID
// @Tags acoes
// @Accept  json
// @Produce  json
// @Param id path int true "AcaoBr ID"
// @Success 200 {object} response.AcaoBrResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/acoes/{id} [get]
func (h *AcaoBrHandler) GetAcaoBr(c *gin.Context) {
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var acao response.AcaoBrResponse
	acao, err = h.acaoBrService.GetAcaoBrByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, acao)
}

// @Summary Create a acao
// @Description Create a new acao
// @Tags acoes
// @Accept  json
// @Produce  json
// @Param acao body request.AcaoBrRequest true "AcaoBrRequest info"
// @Success 201 {object} response.AcaoBrResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/acoes [post]
func (h *AcaoBrHandler) CreateAcaoBr(c *gin.Context) {
	var acao request.AcaoBrRequest
	if err := c.ShouldBindJSON(&acao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAcaoBr, err := h.acaoBrService.CreateAcaoBr(acao)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdAcaoBr)
}

// @Summary Update a acao
// @Description Update a acao by ID
// @Tags acoes
// @Accept  json
// @Produce  json
// @Param id path int true "AcaoBr ID"
// @Param acao body request.AcaoBrRequest true "AcaoBrRequest info"
// @Success 200 {object} response.AcaoBrResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/acoes/{id} [put]
func (h *AcaoBrHandler) UpdateAcaoBr(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var acao request.AcaoBrRequest
	if err := c.ShouldBindJSON(&acao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAcaoBr, err := h.acaoBrService.UpdateAcaoBr(uint(id), acao)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAcaoBr)
}

// @Summary Delete a acao
// @Description Delete a acao by ID
// @Tags acoes
// @Accept  json
// @Produce  json
// @Param id path int true "AcaoBr ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/acoes/{id} [delete]
func (h *AcaoBrHandler) DeleteAcaoBr(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.acaoBrService.DeleteAcaoBr(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}