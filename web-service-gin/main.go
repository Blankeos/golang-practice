package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"
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
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumById responds with 1 album as JSON
func getAlbumById(c *gin.Context) {
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

// ! ------------------ HANDLERS ------------------ ! //

func main() {
	// Create a channel to receive signals.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album", getAlbumById)
	router.POST("/albums", postAlbums)

	fmt.Printf("Running Application on http://localhost:8080\n")

	router.Run("localhost:8080")

	// Block until a signal is received.
	<-quit
	
	// Perform graceful shutdown here.
}