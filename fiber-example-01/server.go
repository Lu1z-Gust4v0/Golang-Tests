package main

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func handleCreateAlbum(context *fiber.Ctx) error {
  album := new(Album)

  if error := context.BodyParser(album); error != nil {
    return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message": "Invalid body request",
      "album": nil,
      "error": error,
    })
  }

  if error := createAlbum(connection, album); error != nil {
    return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "message": "Could not create new album",
      "album": nil,
      "error": error,
    })
  }

  return context.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "Album created successfully",
    "album": album,
    "error": nil,
  })
}

func handleGetAllAlbums(context *fiber.Ctx) error {
  albums, error := getAllAlbums(connection)
  
  if error != nil {
    return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "message": "Could not get the albums",
      "albums": nil,
      "error": error,
    })
  }

  return context.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "Albums fetched successfully",
    "albums": albums,
    "error": nil,
  })
}

func handleGetAlbumById(context *fiber.Ctx) error {
  albumId, error := strconv.Atoi(context.Params("id"))
  
  if error != nil {
    return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message": "Invalid id",
      "album": nil,
      "error": error,
    }) 
  }
  
  album, error := getAlbumById(connection, uint(albumId))

  if error != nil {
    return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "message": "album not found",
      "album": nil,
      "error": error,
    })
  }
  
  return context.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "album found successfully",
    "album": album,
    "error": nil,
  })
}

func main() {
  app := fiber.New()
  
  connectToDB()
  
  app.Get("/albums", handleGetAllAlbums)
  app.Get("/albums/:id", handleGetAlbumById)
  app.Post("/album/create", handleCreateAlbum)

  app.Listen(":8000")
}
