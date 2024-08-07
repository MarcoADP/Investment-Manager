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

type BdrHandler struct {
	bdrService *service.BdrService
}

func NewBdrHandler(bdrService *service.BdrService) *BdrHandler {
	return &BdrHandler{bdrService: bdrService}
}

// @Summary Get All Acão BR
// @Description Get All Ação BR
// @Tags bdrs
// @Accept  json
// @Produce  json
// @Success 200 {object} []response.BdrResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/bdrs [get]
func (h *BdrHandler) GetBdrs(c *gin.Context) {
	bdrs, err := h.bdrService.GetAllBdrs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bdrs)
}

// @Summary Get a Acão BR
// @Description Get a Ação BR by ID
// @Tags bdrs
// @Accept  json
// @Produce  json
// @Param id path int true "Bdr ID"
// @Success 200 {object} response.BdrResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/bdrs/{id} [get]
func (h *BdrHandler) GetBdr(c *gin.Context) {
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var bdr response.BdrResponse
	bdr, err = h.bdrService.GetBdrByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bdr)
}

// @Summary Create a bdr
// @Description Create a new bdr
// @Tags bdrs
// @Accept  json
// @Produce  json
// @Param bdr body request.BdrRequest true "BdrRequest info"
// @Success 201 {object} response.BdrResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/bdrs [post]
func (h *BdrHandler) CreateBdr(c *gin.Context) {
	var bdr request.BdrRequest
	if err := c.ShouldBindJSON(&bdr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBdr, err := h.bdrService.CreateBdr(bdr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdBdr)
}

// @Summary Update a bdr
// @Description Update a bdr by ID
// @Tags bdrs
// @Accept  json
// @Produce  json
// @Param id path int true "Bdr ID"
// @Param bdr body request.BdrRequest true "BdrRequest info"
// @Success 200 {object} response.BdrResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/bdrs/{id} [put]
func (h *BdrHandler) UpdateBdr(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var bdr request.BdrRequest
	if err := c.ShouldBindJSON(&bdr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBdr, err := h.bdrService.UpdateBdr(uint(id), bdr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedBdr)
}

// @Summary Delete a bdr
// @Description Delete a bdr by ID
// @Tags bdrs
// @Accept  json
// @Produce  json
// @Param id path int true "Bdr ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/bdrs/{id} [delete]
func (h *BdrHandler) DeleteBdr(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.bdrService.DeleteBdr(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
