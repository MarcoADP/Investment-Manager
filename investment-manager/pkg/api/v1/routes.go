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

func createMovimentacaoHandler(
	db *gorm.DB,
) MovimentacaoHandler {
	repo := repository.NewMovimentacaoRepository(db)
	service := service.NewMovimentacaoService(repo)
	return *NewMovimentacaoHandler(service)
}

func createConsolidacaoHandler(
	db *gorm.DB,
) ConsolidacaoHandler {
	repo := repository.NewConsolidacaoRepository(db)
	movRepo := repository.NewMovimentacaoRepository(db)
	movService := service.NewMovimentacaoService(movRepo)
	service := service.NewConsolidacaoService(repo, movService)
	return *NewConsolidacaoHandler(service)
}

func createCotacaoHistoricoHandler(
	db *gorm.DB,
) CotacaoHistoricoHandler {
	repo := repository.NewCotacaoHistoricoRepository(db)
	service := service.NewCotacaoHistoricoService(repo)
	return *NewCotacaoHistoricoHandler(service)
}

func createCarteiraHandler(
	db *gorm.DB,
) CarteiraHandler {
	repo := repository.NewCarteiraRepository(db)
	ativoRepo := repository.NewCarteiraAtivoRepository(db)
	consolidacaoRepo := repository.NewConsolidacaoRepository(db)
	cotacaoRepo := repository.NewCotacaoHistoricoRepository(db)
	service := service.NewCarteiraService(repo, ativoRepo, consolidacaoRepo, cotacaoRepo)
	return *NewCarteiraHandler(service)
}

func CreateRoutes(db *gorm.DB) *gin.Engine {

	router := gin.Default()

	acaoBrHandler := createAcaoBrHandler(db)
	fundoImobiliarioBrHandler := createFundoImobiliarioHandler(db)
	bdrHandler := createBdrHandler(db)
	movimentacaoHandler := createMovimentacaoHandler(db)
	consolidacaoHandler := createConsolidacaoHandler(db)
	cotacaoHistoricoHandler := createCotacaoHistoricoHandler(db)
	carteiraHandler := createCarteiraHandler(db)

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

		api.GET("/movimentacoes", movimentacaoHandler.GetMovimentacoes)
		api.GET("/movimentacoes/:id", movimentacaoHandler.GetMovimentacao)
		api.POST("/movimentacoes/entrada", movimentacaoHandler.CreateMovimentacaoEntrada)
		api.POST("/movimentacoes/saida", movimentacaoHandler.CreateMovimentacaoSaida)
		api.DELETE("/movimentacoes/:id", movimentacaoHandler.DeleteMovimentacao)

		api.GET("/consolidacoes", consolidacaoHandler.GetConsolidacoes)
		api.GET("/consolidacoes/:codigo", consolidacaoHandler.GetConsolidacao)
		api.POST("/consolidacoes/calcular", consolidacaoHandler.CalcularConsolidacoes)

		api.GET("/cotacoes", cotacaoHistoricoHandler.GetAllCotacoes)
		api.GET("/cotacoes/:codigo", cotacaoHistoricoHandler.GetCotacoesByCodigo)
		api.GET("/cotacoes/:codigo/last", cotacaoHistoricoHandler.GetCotacaoMoreRecentByCodigo)
		api.POST("/cotacoes", cotacaoHistoricoHandler.CreateCotacao)
		api.DELETE("/cotacoes/:id", cotacaoHistoricoHandler.DeleteCotacao)

		api.GET("/carteiras", carteiraHandler.GetCarteiras)
		api.GET("/carteiras/:id", carteiraHandler.GetCarteira)
		api.POST("/carteiras", carteiraHandler.CreateCarteira)
		api.PUT("/carteiras/:id", carteiraHandler.UpdateCarteira)
		api.DELETE("/carteiras/:id", carteiraHandler.DeleteCarteira)
		api.POST("/carteiras/:id/ativo", carteiraHandler.AddAtivoCarteira)
		api.DELETE("/carteiras/:id/ativo/:codigo", carteiraHandler.DeleteAtivo)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router

}
