package authhandler

import (
	"gosampleapi/singletons"
	"net/http"

	jwtware "github.com/gofiber/jwt/v3"

	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) RegisterFiberRoutes(app *fiber.App) {
	unauthenticated := app.Group("/auth")

	unauthenticated.Post("/register", h.register)
	unauthenticated.Post("/login", h.login)
	unauthenticated.Get("/ok", h.ok)

	authenticated := app.Group("/auth", jwtware.New(jwtware.Config{
		SigningMethod: "EdDSA",
		SigningKey:    h.AuthService.GetPublicKey(),
	}))

	authenticated.Get("/no", h.ok)
}

func (h *AuthHandler) register(c *fiber.Ctx) error {
	var request struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	err = singletons.Validate.Struct(request)
	if err != nil {
		return err
	}

	user, err := h.AuthService.RegisterUser(request.Username, request.Password)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (h *AuthHandler) login(c *fiber.Ctx) error {
	var request struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	err = singletons.Validate.Struct(request)
	if err != nil {
		return err
	}

	accessToken, err := h.AuthService.GenerateAccessToken(request.Username, request.Password)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "invalid credential")
	}

	return c.JSON(accessToken)
}

func (h *AuthHandler) ok(c *fiber.Ctx) error {
	return c.JSON("ok")
}
