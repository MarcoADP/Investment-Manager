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

type AcaoBrComparadorHandler struct {
	acaoBrComparadorService *service.AcaoBrComparadorService
}

func NewAcaoBrComparadorHandler(acaoBrComparadorService *service.AcaoBrComparadorService) *AcaoBrComparadorHandler {
	return &AcaoBrComparadorHandler{acaoBrComparadorService: acaoBrComparadorService}
}

// @Summary Compare List of Acoes
// @Description Compare List of Acoes
// @Tags comparador
// @Accept  json
// @Produce  json
// @Param codigos query []string true "Lista de Códigos das Ações"
// @Success 200 {object} response.AcaoBrComparadorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/comparador [get]
func (h *AcaoBrComparadorHandler) CompareAcoesBr(c *gin.Context) {

	codigos := c.QueryArray("codigos")

	var comparacao response.AcaoBrComparadorResponse
	var err error
	comparacao, err = h.acaoBrComparadorService.CompareAcoes(codigos)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comparacao)
}
