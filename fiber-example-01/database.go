package main

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Album struct {
  ID        uint      `gorm:"primaryKey" json:"id"`       
  Title     string    `gorm:"notNull" json:"title"`
  Artist    string    `gorm:"notNull" json:"artist"`
  Price     float64   `gorm:"notNull" json:"price"`
}

var connection *gorm.DB

func connectToDB() {
  dsn := "host=localhost user=ubuntu password=ubuntu dbname=test port=5432"
  var error error
  
  connection, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if error != nil {
    log.Panic("Failed to connect to database")
  }
  
  if !connection.Migrator().HasTable(&Album{}) {
    connection.Migrator().CreateTable(&Album{})
  }  
}

func createAlbum(db *gorm.DB, album *Album) error {
  result := db.Create(album)

  if result.Error != nil {
    return result.Error
  }

  return nil
}

func getAlbumById(db *gorm.DB, albumId uint) (*Album, error) {
  var album Album

  // SELECT * FROM albums WHERE id = albumId LIMIT 1
  result := db.First(&album, albumId)
  
  if result.Error != nil {
    return nil, result.Error 
  }
  
  return &album, nil
}

func getAllAlbums(db *gorm.DB) ([]Album, error) {
  var albums []Album 

  // SELECT * FROM albums
  result := db.Find(&albums)

  if result.Error != nil {
    return nil, result.Error
  }

  return albums, nil
}
