package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)


// album represents data about a record album.
type album struct {
  ID        string  `json:"id"`
  Title     string  `json:"title"`
  Artist    string  `json:"artist"`
  Price     float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album {
  { ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 59.99 },
  { ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99 },
  { ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99 },
}

// getAlbums responds with the list of all albums as JSON 
func getAlbums(context *gin.Context) {
  context.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body
func postAlbums(context *gin.Context) {
  var newAlbum album

  // Call BindJSON to bind the recieved JSON to newAlbum
  if error := context.BindJSON(&newAlbum); error != nil {
    return 
  }

  // Add new album to the slice
  albums = append(albums, newAlbum)
  context.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id 
// parameter sent by the client, then returns that album as a response
func getAlbumByID(context *gin.Context) {
  id := context.Param("id")

  // Loop over the list of albums, looking for an album that matches 
  // the id provided in the parameter
  for _, album := range albums {
    if album.ID == id {
      context.IndentedJSON(http.StatusOK, album)
      return 
    }
  }
  
  context.IndentedJSON(http.StatusNotFound, gin.H {
    "message": "album not found",
  })
}

func main() {
  router := gin.Default()
  router.GET("/albums", getAlbums)
  router.GET("/albums/:id", getAlbumByID)
  router.POST("/albums", postAlbums)

  router.Run("localhost:8000")
}

