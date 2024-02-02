package handlers

import (
	"app/templates/components"
	"app/views"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func ContactHandler(c echo.Context) error {
	// Get form data
	contact := make(map[string]string)
	if err := c.Bind(&contact); err != nil {
		return err
	}

	// Get Fields
	email := contact["email"]
	message := contact["message"]

	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	to := "to@example.com"
	body := []byte("To: " + to + "\r\n" +

		"Subject: This was sent using HTMX - Go - Templ?\r\n" +

		"\r\n" +

		message + "\r\n")

	auth := smtp.PlainAuth("", username, password, "sandbox.smtp.mailtrap.io")

	err := smtp.SendMail("sandbox.smtp.mailtrap.io:587", auth, email, []string{to}, body)
	if err != nil {
		return components.Toast(false).Render(c.Request().Context(), c.Response())
	}

	// Send response
	return components.Toast(true).Render(c.Request().Context(), c.Response())
}

func HomeHandler(c echo.Context) error {
	components.Layout(views.Index()).Render(c.Request().Context(), c.Response())
	return nil
}

func ToastRemover(c echo.Context) error {
	return c.HTML(http.StatusOK, "")
}
