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

type AtivoInformacaoHandler struct {
	ativoInformacaoService *service.AtivoInformacaoService
}

func NewAtivoInformacaoHandler(ativoInformacaoService *service.AtivoInformacaoService) *AtivoInformacaoHandler {
	return &AtivoInformacaoHandler{ativoInformacaoService: ativoInformacaoService}
}

// @Summary Get a informacao
// @Description Get a informacao by Codigo
// @Tags informacoes
// @Accept  json
// @Produce  json
// @Param codigo path int true "Informacao Codigo"
// @Success 200 {object} []response.AtivoInformacaoResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/informacoes/{codigo} [get]
func (h *AtivoInformacaoHandler) GetInformacoesByCodigo(c *gin.Context) {

	codigo := c.Param("codigo")
	informacoes, err := h.ativoInformacaoService.GetInformacoesByCodigo(codigo)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, informacoes)
}

// @Summary Get a informacao
// @Description Get a informacao by Codigo
// @Tags informacoes
// @Accept  json
// @Produce  json
// @Param codigo path int true "Informacao Codigo"
// @Success 200 {object} response.AtivoInformacaoResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/informacoes/{codigo}/last [get]
func (h *AtivoInformacaoHandler) GetInformacaoMoreRecentByCodigo(c *gin.Context) {
	var err error
	codigo := c.Param("codigo")

	informacao, err := h.ativoInformacaoService.GetInformacaoMoreRecentByCodigo(codigo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, informacao)
}

// @Summary Criar a informacao
// @Description Criar a informacao
// @Tags informacoes
// @Accept  json
// @Produce  json
// @Success 201 {object} response.AtivoInformacaoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/informacoes [post]
func (h *AtivoInformacaoHandler) CreateInformacao(c *gin.Context) {
	var informacaoRequest request.AtivoInformacaoRequest
	var informacaoResponse response.AtivoInformacaoResponse
	if err := c.ShouldBindJSON(&informacaoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	informacaoResponse, err := h.ativoInformacaoService.CreateInformacao(informacaoRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, informacaoResponse)
}

// @Summary Delete a Informacao
// @Description Delete a Informacao by ID
// @Tags informacoes
// @Accept  json
// @Produce  json
// @Param id path int true "Informacao ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/informacoes/{id} [delete]
func (h *AtivoInformacaoHandler) DeleteInformacao(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.ativoInformacaoService.DeleteInformacao(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
