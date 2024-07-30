package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"auction/dto"
	"auction/entity"
	"auction/service"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local", db_username, db_password)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&entity.Artwork{}, &entity.Buyer{}, &entity.Bidding{}, &entity.Auction{})

	if err != nil {
		log.Fatalf("Error initialize database: %s", err)
	}

	app := fiber.New()

	svc := &service.Service{
		DB: db,
	}

	app.Post("/auction", func(c *fiber.Ctx) error {
		request := &dto.CreateAuctionRequest{}

		if err := c.BodyParser(request); err != nil {
			return err
		}

		auction, err := svc.CreateAuction(request.Name, time.Time(request.EndDate))

		if err != nil {
			return err
		}

		response := &dto.CreateAuctionResponse{
			ID: auction.ID,
		}

		return c.Status(http.StatusOK).JSON(response)
	})

	log.Fatal(app.Listen(":3000"))
}
