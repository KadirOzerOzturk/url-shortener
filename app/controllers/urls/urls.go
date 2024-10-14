package urls

import (
	"fmt"
	"net/http"
	"time"

	"github.com/KadirOzerOzturk/url-shortener/app/entities"
	"github.com/KadirOzerOzturk/url-shortener/app/helpers"
	"github.com/KadirOzerOzturk/url-shortener/internal/database"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	db := database.Connection()
	if db == nil {
		return c.Status(500).SendString("Failed to connect to database")
	}

	result, err := helpers.AllShortUrls()
	if err != nil {
		return c.Status(500).SendString("Failed to fetch URLs")
	}

	if len(result) == 0 {
		return c.Status(404).SendString("No records found")
	}

	return c.JSON(result)
}

func Shorten(c *fiber.Ctx) error {
	shortenRequest := new(entities.ShortenRequest)
	if err := c.BodyParser(shortenRequest); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	/*
		ip := c.IP()
		current, err := helpers.IncrementRateLimit(ip)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		if current >= 5 {
			return c.Status(429).SendString("Rate limit exceeded. Please try again later.")
		}
	*/
	shortUrlStr := helpers.GenerateShortUrl()

	urlNew := entities.Url{
		OriginalUrl:  shortenRequest.OriginalUrl,
		ShortenedUrl: "http://10.150.238.245:3000/url/" + shortUrlStr,
		UsageCount:   0,
		ExpiresAt:    time.Now().Add(24 * time.Hour),
	}

	err := database.Connection().Model(&entities.Url{}).Create(&urlNew).Error
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(urlNew)
}

func Redirect(c *fiber.Ctx) error {
	short_url := "http://10.150.238.245:3000/url/" + c.Params("short_url")
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

	return c.Redirect(url.OriginalUrl, http.StatusMovedPermanently)
}

func Delete(c *fiber.Ctx) error {

	short_url := "http://10.150.238.245:3000/url/" + c.Params("short_url")

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
	short_url := "http://10.150.238.245:3000/url/" + c.Params("short_url")
	result := database.Connection().First(&url, "shortened_url = ?", short_url)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Status(200).JSON(url)
}
