package mails

import (
	"log"

	"github.com/KadirOzerOzturk/url-shortener/app/entities"
	"github.com/KadirOzerOzturk/url-shortener/app/helpers"
	"github.com/gofiber/fiber/v2"
)

func SendMail(c *fiber.Ctx) error {
	mail := entities.Mail{}
	if err := c.BodyParser(&mail); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	if mail.Subject == "" {
		return c.Status(400).SendString("Subject is required")
	}
	if mail.To == "" {
		return c.Status(400).SendString("To is required")
	}
	if mail.Body == "" {
		return c.Status(400).SendString("Body is required")
	}
	err := helpers.SendMail(mail)
	if err != nil {
		log.Fatalf("Failed to send mail: %v", err)
		return c.Status(500).SendString("Failed to send mail")
	}
	return c.SendString("Mail sent successfully")
}
