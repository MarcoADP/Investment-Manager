package v1

import (
	"net/http"

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

type ConsolidacaoHandler struct {
	consolidacaoService *service.ConsolidacaoService
}

func NewConsolidacaoHandler(consolidacaoService *service.ConsolidacaoService) *ConsolidacaoHandler {
	return &ConsolidacaoHandler{consolidacaoService: consolidacaoService}
}

// @Summary Get All consolidacao BR
// @Description Get All consolidacao BR
// @Tags consolidacoes
// @Accept  json
// @Produce  json
// @Success 200 {object} []response.ConsolidacaoResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/consolidacoes [get]
func (h *ConsolidacaoHandler) GetConsolidacoes(c *gin.Context) {
	var consolidacoes []response.ConsolidacaoResponse
	var err error
	consolidacoes, err = h.consolidacaoService.GetAllConsolidacoes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, consolidacoes)
}

// @Summary Get a consolidacao BR
// @Description Get a consolidacao BR by Codigo
// @Tags consolidacoes
// @Accept  json
// @Produce  json
// @Param id path int true "Consolidacao Codigo"
// @Success 200 {object} response.ConsolidacaoResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/consolidacoes/{codigo} [get]
func (h *ConsolidacaoHandler) GetConsolidacao(c *gin.Context) {
	var err error
	codigo := c.Param("codigo")

	var consolidacao response.ConsolidacaoResponse
	consolidacao, err = h.consolidacaoService.GetConsolidacaoByCodigo(codigo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, consolidacao)
}

// @Summary Calcular a consolidacao
// @Description  Calcular a consolidacao
// @Tags consolidacoes
// @Accept  json
// @Produce  json
// @Success 201 {object} []response.ConsolidacaoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/consolidacoes/calcular [post]
func (h *ConsolidacaoHandler) CalcularConsolidacoes(c *gin.Context) {
	consolidacoes, err := h.consolidacaoService.GenerateConsolidacoes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, consolidacoes)
}
