package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/MarcoADP/Investment-Manager/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Item representa um item na API
type Item struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"Item One"`
}

var items = []Item{
	{ID: 1, Name: "Item One"},
	{ID: 2, Name: "Item Two"},
}

type ErrorResponse struct {
	Error string `json:"error" example:"Bad Request"`
}

// @Summary Get all items
// @Description Get all items in the system
// @Tags items
// @Produce json
// @Success 200 {array} Item
// @Router /items [get]
func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

// @Summary Get item by ID
// @Description Get a single item by ID
// @Tags items
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} Item
// @Failure 404 {object} ErrorResponse
// @Router /items/{id} [get]
func getItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// @Summary Create new item
// @Description Create a new item in the system
// @Tags items
// @Accept json
// @Produce json
// @Param item body Item true "New Item"
// @Success 201 {object} Item
// @Failure 404 {object} ErrorResponse
// @Router /items [post]
func createItem(c *gin.Context) {
	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = len(items) + 1
	items = append(items, item)
	c.JSON(http.StatusCreated, item)
}

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Obter valores das variáveis de ambiente
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construir string de conexão
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Conectar ao banco de dados
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}
	defer db.Close()

	// Verificar se a conexão foi bem-sucedida
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	fmt.Println("Conexão com o banco de dados bem-sucedida!")

	router := gin.Default()

	// Rotas
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/items", getItems)
	router.GET("/items/:id", getItem)
	router.POST("/items", createItem)

	router.Run(":8080")

}
