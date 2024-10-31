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
	setorRepo := repository.NewSetorRepository(db)
	service := service.NewAcaoBrService(repo, setorRepo)
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
	carteiraAtivoRep := repository.NewCarteiraAtivoRepository(db)
	service := service.NewCotacaoHistoricoService(repo, carteiraAtivoRep)
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

func createAtivoValuationService(db *gorm.DB) *service.AtivoValuationService {
	valuationRepo := repository.NewAtivoValuationRepository(db)
	return service.NewAtivoValuationService(valuationRepo)
}

func createAtivoEndividamentoService(db *gorm.DB) *service.AtivoEndividamentoService {
	endividamentoRepo := repository.NewAtivoEndividamentoRepository(db)
	return service.NewAtivoEndividamentoService(endividamentoRepo)
}

func createAtivoEficienciaService(db *gorm.DB) *service.AtivoEficienciaService {
	eficienciaRepo := repository.NewAtivoEficienciaRepository(db)
	return service.NewAtivoEficienciaService(eficienciaRepo)
}

func createAtivoRentabilidadeService(db *gorm.DB) *service.AtivoRentabilidadeService {
	rentabilidadeRepo := repository.NewAtivoRentabilidadeRepository(db)
	return service.NewAtivoRentabilidadeService(rentabilidadeRepo)
}

func createAtivoDividendoService(db *gorm.DB) *service.AtivoDividendoService {
	ativoDividendoRep := repository.NewAtivoDividendoRepository(db)
	dividendosRep := repository.NewDividendoRepository(db)
	consolidadeRep := repository.NewConsolidacaoRepository(db)
	return service.NewAtivoDividendoService(ativoDividendoRep, dividendosRep, consolidadeRep)
}

func createAtivoInformacaoService(db *gorm.DB) *service.AtivoInformacaoService {
	repo := repository.NewAtivoInformacaoRepository(db)
	cotacaoRepo := repository.NewCotacaoHistoricoRepository(db)
	ativoValuationService := createAtivoValuationService(db)
	ativoEndividamentoService := createAtivoEndividamentoService(db)
	ativoEficienciaService := createAtivoEficienciaService(db)
	ativoRentabilidadeService := createAtivoRentabilidadeService(db)
	ativoDividendoService := createAtivoDividendoService(db)
	return service.NewAtivoInformacaoService(repo, cotacaoRepo, ativoValuationService, ativoEndividamentoService, ativoEficienciaService, ativoRentabilidadeService,
		ativoDividendoService)
}

func createAtivoHandlerHandler(
	db *gorm.DB,
) AtivoInformacaoHandler {
	service := createAtivoInformacaoService(db)
	return *NewAtivoInformacaoHandler(service)
}

func createDividendoHandler(
	db *gorm.DB,
) DividendoHandler {
	repo := repository.NewDividendoRepository(db)
	service := service.NewDividendoService(repo)
	return *NewDividendoHandler(service)
}

func createGrahamFormulaHandler(
	db *gorm.DB,
) GrahamFormulaHandler {
	repo := repository.NewGrahamFormulaRepository(db)
	service := service.NewGrahamFormulaService(repo)
	return *NewGrahamFormulaHandler(service)
}

func createAcaoBrComparadorHandler(
	db *gorm.DB,
) AcaoBrComparadorHandler {
	ativoInformacaoService := createAtivoInformacaoService(db)
	service := service.NewAcaoBrComparadorService(ativoInformacaoService)
	return *NewAcaoBrComparadorHandler(service)
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
	ativoInformacaoHandler := createAtivoHandlerHandler(db)
	dividendoHandler := createDividendoHandler(db)
	grahamFormulaHandler := createGrahamFormulaHandler(db)
	AcaoBrComparadorHandler := createAcaoBrComparadorHandler(db)

	api := router.Group("/api/v1")
	{
		api.GET("/acoes", acaoBrHandler.GetAcoesBr)
		api.GET("/acoes/:codigo", acaoBrHandler.GetAcaoBrByCodigo)
		api.GET("/acoes/setor/:setor", acaoBrHandler.GetAcoesBySetor)
		api.POST("/acoes", acaoBrHandler.CreateAcaoBr)
		api.PUT("/acoes/:codigo", acaoBrHandler.UpdateAcaoBr)
		api.DELETE("/acoes/:codigo", acaoBrHandler.DeleteAcaoBr)

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
		api.POST("/cotacoes/brapi", cotacaoHistoricoHandler.CreateCotacaoBrapi)
		api.POST("/cotacoes/carteira/:carteiraId", cotacaoHistoricoHandler.CreateCotacoesCarteira)
		api.DELETE("/cotacoes/:id", cotacaoHistoricoHandler.DeleteCotacao)

		api.GET("/carteiras", carteiraHandler.GetCarteiras)
		api.GET("/carteiras/:id", carteiraHandler.GetCarteira)
		api.POST("/carteiras", carteiraHandler.CreateCarteira)
		api.PUT("/carteiras/:id", carteiraHandler.UpdateCarteira)
		api.DELETE("/carteiras/:id", carteiraHandler.DeleteCarteira)
		api.POST("/carteiras/:id/ativo", carteiraHandler.AddAtivoCarteira)
		api.DELETE("/carteiras/:id/ativo/:codigo", carteiraHandler.DeleteAtivo)

		api.GET("/informacoes/:codigo", ativoInformacaoHandler.GetInformacoesByCodigo)
		api.GET("/informacoes/:codigo/last", ativoInformacaoHandler.GetInformacaoMoreRecentByCodigo)
		api.POST("/informacoes", ativoInformacaoHandler.CreateInformacao)
		api.DELETE("/informacoes/:id", ativoInformacaoHandler.DeleteInformacao)

		api.GET("/dividendos", dividendoHandler.GetDividendos)
		api.GET("/dividendos/:codigo", dividendoHandler.GetDividendosByCodigo)
		api.POST("/dividendos", dividendoHandler.CreateDividendo)
		api.DELETE("/dividendos/:id", dividendoHandler.DeleteDividendo)

		api.GET("/graham-formula/:codigo", grahamFormulaHandler.GetGrahamFormulaByCodigo)
		api.GET("/graham-formula/:codigo/last", grahamFormulaHandler.GetGrahamFormulaMoreRecentByCodigo)
		api.POST("/graham-formula", grahamFormulaHandler.CreateGrahamFormula)
		api.DELETE("/graham-formula/:id", grahamFormulaHandler.DeleteGrahamFormula)

		api.GET("/comparador", AcaoBrComparadorHandler.CompareAcoesBr)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router

}
