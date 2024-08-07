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

type MovimentacaoHandler struct {
	movimentacaoService *service.MovimentacaoService
}

func NewMovimentacaoHandler(movimentacaoService *service.MovimentacaoService) *MovimentacaoHandler {
	return &MovimentacaoHandler{movimentacaoService: movimentacaoService}
}

// @Summary Get All movimentacao BR
// @Description Get All movimentacao BR
// @Tags movimentacoes
// @Accept  json
// @Produce  json
// @Param codigo query string false "Codigo"
// @Success 200 {object} []response.MovimentacaoResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/movimentacoes [get]
func (h *MovimentacaoHandler) GetMovimentacoes(c *gin.Context) {
	codigo := c.Query("codigo")

	var movimentacoes []response.MovimentacaoResponse
	var err error
	if codigo == "" {
		movimentacoes, err = h.movimentacaoService.GetAllMovimentacoes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		movimentacoes, err = h.movimentacaoService.GetAllMovimentacaoByCodigo(codigo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, movimentacoes)
}

// @Summary Get a movimentacao BR
// @Description Get a movimentacao BR by ID
// @Tags movimentacoes
// @Accept  json
// @Produce  json
// @Param id path int true "Movimentacao ID"
// @Success 200 {object} response.MovimentacaoResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/movimentacoes/{id} [get]
func (h *MovimentacaoHandler) GetMovimentacao(c *gin.Context) {
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var movimentacao response.MovimentacaoResponse
	movimentacao, err = h.movimentacaoService.GetMovimentacaoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movimentacao)
}

// @Summary Create a movimentacao de entrada
// @Description Create a new movimentacao de entrada
// @Tags movimentacoes
// @Accept  json
// @Produce  json
// @Param movimentacao body request.MovimentacaoRequest true "MovimentacaoRequest info"
// @Success 201 {object} response.MovimentacaoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/movimentacoes/entrada [post]
func (h *MovimentacaoHandler) CreateMovimentacaoEntrada(c *gin.Context) {
	var movimentacao request.MovimentacaoRequest
	if err := c.ShouldBindJSON(&movimentacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMovimentacao, err := h.movimentacaoService.CreateMovimentacao(movimentacao, "ENTRADA")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdMovimentacao)
}

// @Summary Create a movimentacao de saida
// @Description Create a new movimentacao de saida
// @Tags movimentacoes
// @Accept  json
// @Produce  json
// @Param movimentacao body request.MovimentacaoRequest true "MovimentacaoRequest info"
// @Success 201 {object} response.MovimentacaoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/movimentacoes/saida [post]
func (h *MovimentacaoHandler) CreateMovimentacaoSaida(c *gin.Context) {
	var movimentacao request.MovimentacaoRequest
	if err := c.ShouldBindJSON(&movimentacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMovimentacao, err := h.movimentacaoService.CreateMovimentacao(movimentacao, "SAIDA")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdMovimentacao)
}

// @Summary Delete a movimentacao
// @Description Delete a movimentacao by ID
// @Tags movimentacoes
// @Accept  json
// @Produce  json
// @Param id path int true "Movimentacao ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/movimentacoes/{id} [delete]
func (h *MovimentacaoHandler) DeleteMovimentacao(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.movimentacaoService.DeleteMovimentacao(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
