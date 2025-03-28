package urls

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/KadirOzerOzturk/url-shortener/app/entities"
	"github.com/KadirOzerOzturk/url-shortener/app/helpers"
	"github.com/KadirOzerOzturk/url-shortener/internal/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	result, err := helpers.AllShortUrls()
	if err != nil {
		log.Println("Failed to fetch URLs:", err)
		return c.Status(500).SendString("Failed to fetch URLs")
	}

	if len(result) == 0 {
		return c.Status(404).SendString("No records found")
	}

	return c.JSON(result)
}

func Shorten(c *fiber.Ctx) error {
	shortenRequest := entities.ShortenRequest{}
	if err := c.BodyParser(&shortenRequest); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	if shortenRequest.OriginalUrl == "" {
		return c.Status(400).SendString("Original URL is required")
	}
	shortUrlStr := helpers.GenerateShortUrl()
	expireHours, err := strconv.Atoi(os.Getenv("URL_EXPIRE"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	urlNew := entities.Url{
		OriginalUrl:  shortenRequest.OriginalUrl,
		ShortenedUrl: os.Getenv("BASE_URL") + shortUrlStr,
		UsageCount:   0,
		ExpiresAt:    time.Now().Add(time.Duration(expireHours) * time.Hour),
		UserEmail:    shortenRequest.UserEmail,
	}

	err = database.Connection().Model(&entities.Url{}).Create(&urlNew).Error
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(urlNew)
}

func GetUrlsByUser(c *fiber.Ctx) error {
	userEmail := strings.TrimSpace(c.Params("email")) // Trim any spaces

	// Optionally, convert to lowercase if needed (emails are case-insensitive)
	userEmail = strings.ToLower(userEmail)

	var urls []entities.Url

	// Perform the query with a consistent structure
	result := database.Connection().Where("user_email = ?", userEmail).Find(&urls)

	if result.Error != nil {
		// Log the error for better debugging
		log.Println("Error fetching URLs by user:", result.Error)
		return c.Status(500).SendString(result.Error.Error())
	}

	if len(urls) == 0 {
		return c.Status(404).SendString("No records found")
	}

	return c.JSON(urls)
}
func Redirect(c *fiber.Ctx) error {
	short_url := os.Getenv("BASE_URL") + c.Params("short_url")
	//short_url := c.Params("short_url")
	fmt.Println("short_url : ", short_url)

	var url entities.Url
	result := database.Connection().First(&url, "shortened_url = ?", short_url)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	if url.ExpiresAt.Before(time.Now()) {
		return c.Status(500).SendString("Url has expired")
	}

	go helpers.IncClickCount(url)
	go helpers.SaveAccessDetails(url, c.IP())

	return c.Status(200).JSON(fiber.Map{
		"originalUrl": url.OriginalUrl,
	})
}
func Delete(c *fiber.Ctx) error {
	short_url := c.Params("short_url")

	var url entities.Url
	if err := database.Connection().Model(&entities.Url{}).Where("shortened_url = ?", short_url).First(&url).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return c.Status(404).JSON(fiber.Map{
				"error": "Record not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to find item",
		})
	}

	if err := database.Connection().Model(&entities.Url{}).Where("shortened_url = ?", short_url).Delete(&entities.Url{}).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete item",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Item deleted successfully",
	})
}

func UrlStats(c *fiber.Ctx) error {
	var url entities.Url
	short_url := os.Getenv("BASE_URL") + c.Params("short_url")
	result := database.Connection().First(&url, "shortened_url = ?", short_url)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Status(200).JSON(url)
}
