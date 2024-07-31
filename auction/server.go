package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"auction/dto"
	"auction/entity"
	"auction/scheduler"
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

	valid := validator.New()
	valid.RegisterValidation("dateFormatCheck", func(fl validator.FieldLevel) bool {
		_, err := time.Parse("2006-01-02", fl.Field().String())

		return err == nil
	})

	if err != nil {
		log.Fatalf("Error initialize database: %s", err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			response := &dto.ErrorResponse{
				Message: err.Error(),
			}

			return c.Status(code).JSON(response)
		},
	})

	svc := &service.Service{
		DB: db,
	}

	scheduler := &scheduler.Scheduler{
		Service: svc,
	}

	scheduler.InitiateTask()

	app.Post("/auction", func(c *fiber.Ctx) error {
		request := &dto.CreateAuctionRequest{}

		if err := c.BodyParser(request); err != nil {
			return err
		}

		if errs := valid.Struct(request); errs != nil {

			errorsString := ""

			for _, err := range errs.(validator.ValidationErrors) {
				errorsString += err.Field()
				errorsString += " "
			}

			return &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: fmt.Sprintf("Validation Failed: %s", errorsString),
			}
		}

		auction, err := svc.CreateAuction(request.Name, time.Time(request.EndDate))

		if err != nil {
			return err
		}

		response := &dto.CreateAuctionResponse{
			ID: auction.ID,
		}

		return c.Status(http.StatusCreated).JSON(response)
	})

	app.Post("/auction/:id/request-bid", func(c *fiber.Ctx) error {
		request := &dto.CreateBiddingRequest{}

		if err := c.BodyParser(request); err != nil {
			return err
		}

		if errs := valid.Struct(request); errs != nil {

			errorsString := ""

			for _, err := range errs.(validator.ValidationErrors) {
				errorsString += err.Field()
				errorsString += " "
			}

			return &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: fmt.Sprintf("Validation Failed: %s", errorsString),
			}
		}

		auctionID, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return err
		}

		_, err = svc.CreateBidding(auctionID, request.UserID, request.RequestPrice)

		if err != nil {
			return err
		}

		response := &dto.CreateBiddingResponse{
			IsSuccess: true,
		}

		return c.Status(http.StatusCreated).JSON(response)
	})

	log.Fatal(app.Listen(":3000"))
}
