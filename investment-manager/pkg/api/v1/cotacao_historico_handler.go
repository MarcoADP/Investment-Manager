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

type CotacaoHistoricoHandler struct {
	cotacaoHistoricoService *service.CotacaoHistoricoService
}

func NewCotacaoHistoricoHandler(cotacaoHistoricoService *service.CotacaoHistoricoService) *CotacaoHistoricoHandler {
	return &CotacaoHistoricoHandler{cotacaoHistoricoService: cotacaoHistoricoService}
}

// @Summary Get All cotacao
// @Description Get All cotacao
// @Tags cotacoes
// @Accept  json
// @Produce  json
// @Success 200 {object} []response.CotacaoHistoricoResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/cotacoes [get]
func (h *CotacaoHistoricoHandler) GetAllCotacoes(c *gin.Context) {
	var cotacoes []response.CotacaoHistoricoResponse
	var err error
	cotacoes, err = h.cotacaoHistoricoService.GetAllCotacoes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cotacoes)
}

// @Summary Get a cotacao
// @Description Get a cotacao by Codigo
// @Tags cotacoes
// @Accept  json
// @Produce  json
// @Param codigo path int true "Cotacao Codigo"
// @Param data query string false "data"
// @Success 200 {object} []response.CotacaoHistoricoResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/cotacoes/{codigo} [get]
func (h *CotacaoHistoricoHandler) GetCotacoesByCodigo(c *gin.Context) {
	var err error
	codigo := c.Param("codigo")
	data := c.Query("data")

	var cotacoes []response.CotacaoHistoricoResponse
	var cotacao response.CotacaoHistoricoResponse
	if data == "" {
		cotacoes, err = h.cotacaoHistoricoService.GetCotacoesByCodigo(codigo)
	} else {
		cotacao, err = h.cotacaoHistoricoService.GetCotacaoByCodigoAndData(codigo, data)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		cotacoes = append(cotacoes, cotacao)
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cotacoes)
}

// @Summary Get a cotacao
// @Description Get a cotacao by Codigo
// @Tags cotacoes
// @Accept  json
// @Produce  json
// @Param codigo path int true "Cotacao Codigo"
// @Success 200 {object} response.CotacaoHistoricoResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/cotacoes/{codigo}/last [get]
func (h *CotacaoHistoricoHandler) GetCotacaoMoreRecentByCodigo(c *gin.Context) {
	var err error
	codigo := c.Param("codigo")

	cotacao, err := h.cotacaoHistoricoService.GetCotacaoMoreRecentByCodigo(codigo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cotacao)
}

// @Summary Criar a cotacao
// @Description Criar a cotacao
// @Tags cotacoes
// @Accept  json
// @Produce  json
// @Success 201 {object} response.CotacaoHistoricoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/cotacoes [post]
func (h *CotacaoHistoricoHandler) CreateCotacao(c *gin.Context) {
	var cotacaoRequest request.CotacaoHistoricoRequest
	if err := c.ShouldBindJSON(&cotacaoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCotacao, err := h.cotacaoHistoricoService.CreateCotacao(cotacaoRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdCotacao)
}

// @Summary Delete a Cotacao
// @Description Delete a Cotacao by ID
// @Tags cotacoes
// @Accept  json
// @Produce  json
// @Param id path int true "Cotacao ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/cotacoes/{id} [delete]
func (h *CotacaoHistoricoHandler) DeleteCotacao(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.cotacaoHistoricoService.DeleteCotacao(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
