package v1

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gorm.io/gorm"
)

func createAcaoBrHandler(
	db *gorm.DB,
) AcaoBrHandler {
	repo := repository.NewAcaoBrRepository(db)
	service := service.NewAcaoBrService(repo)
	return *NewAcaoBrHandler(service)
}

func createFundoImobiliarioHandler(
	db *gorm.DB,
) FundoImobiliarioHandler {
	repo := repository.NewFundoImobiliarioRepository(db)
	service := service.NewFundoImobiliarioService(repo)
	return *NewFundoImobiliarioHandler(service)
}

func createBdrHandler(
	db *gorm.DB,
) BdrHandler {
	repo := repository.NewBdrRepository(db)
	service := service.NewBdrService(repo)
	return *NewBdrHandler(service)
}

func CreateRoutes(db *gorm.DB) *gin.Engine {

	router := gin.Default()

	acaoBrHandler := createAcaoBrHandler(db)
	fundoImobiliarioBrHandler := createFundoImobiliarioHandler(db)
	bdrHandler := createBdrHandler(db)

	api := router.Group("/api/v1")
	{
		api.GET("/acoes", acaoBrHandler.GetAcoesBr)
		api.GET("/acoes/:id", acaoBrHandler.GetAcaoBr)
		api.POST("/acoes", acaoBrHandler.CreateAcaoBr)
		api.PUT("/acoes/:id", acaoBrHandler.UpdateAcaoBr)
		api.DELETE("/acoes/:id", acaoBrHandler.DeleteAcaoBr)

		api.GET("/fundos-imobiliarios", fundoImobiliarioBrHandler.GetFundosImobiliarios)
		api.GET("/fundos-imobiliarios/:id", fundoImobiliarioBrHandler.GetFundoImobiliario)
		api.POST("/fundos-imobiliarios", fundoImobiliarioBrHandler.CreateFundoImobiliario)
		api.PUT("/fundos-imobiliarios/:id", fundoImobiliarioBrHandler.UpdateFundoImobiliario)
		api.DELETE("/fundos-imobiliarios/:id", fundoImobiliarioBrHandler.DeleteFundoImobiliario)

		api.GET("/bdrs", bdrHandler.GetBdrs)
		api.GET("/bdrs/:id", bdrHandler.GetBdr)
		api.POST("/bdrs", bdrHandler.CreateBdr)
		api.PUT("/bdrs/:id", bdrHandler.UpdateBdr)
		api.DELETE("/bdrs/:id", bdrHandler.DeleteBdr)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router

}