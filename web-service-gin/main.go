package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	docs "example/web-service-gin/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// album represents data about a record album.
type album struct {
	ID     int  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


// + ------------------ HANDLERS ------------------ + //

// getAlbums godoc
// @Summary gets all albums
// @Schemes
// @Description gets all albums from the database.
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {object} []album "Successfully retrieved the albums"
// @Router /albums [get]
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbum godoc
// @Summary gets all albums
// @Schemes
// @Description gets one (1) album using a query.
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {object} album "Successfully fetched album."
// @Router /album [get]
func getAlbum(c *gin.Context) {
	var id, found = c.GetQuery("id")

	// Ensure ID is found on the context.
	if !found {
		c.IndentedJSON(http.StatusBadRequest, "😭 ID not found in context!")
		return
	}

	// Convert the ID to correct type (e.g., int)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "😭 Invalid ID type in context!")
		return
	}

	// Search for album by ID
	var foundAlbum *album = nil
	for i := 0; i < len(albums); i++ {
		if (idInt == albums[i].ID) {
			foundAlbum = &albums[i]
			break
		}
	}

	// Check if an album was found
	if foundAlbum != nil {
		c.IndentedJSON(http.StatusOK, foundAlbum)
	} else {
		c.IndentedJSON(http.StatusNotFound, "Album not found!")
	}
}

// postAlbums adds a new album from JSON received in the request body.
//
// postAlbums godoc
// @Summary adds a new album
// @Schemes
// @Description postAlbums adds a new album from JSON received in the request body.
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {object} album "Successfully added album."
// @Router /albums [post]
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "Cannot bind request body to album")
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	// Convert id (string) to int.
	intId, err := strconv.Atoi(id);
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "😭 Invalid ID (not an int)!")
		return
	}

	// Loop over the list of albums, looking for an
	// album whose ID value matches the paramter.
	for _, a := range albums {
		if a.ID == intId {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found."})
}

// ! ------------------ HANDLERS ------------------ ! //

// @title           Carlo's Album API
// @version         1.0
// @description     This is a generated swagger API from the Album API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Create a channel to receive signals.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	
	docs.SwaggerInfo.Title = "Swagger Example API"

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById) // uses /albums/<id>
	router.GET("/album", getAlbum) // uses /album?id=<id>
	router.POST("/albums", postAlbums)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Printf("Running Application on http://localhost:8080\n")

	router.Run("127.0.0.1:8080")

	// Block until a signal is received.
	<-quit
	
	// Perform graceful shutdown here.
}