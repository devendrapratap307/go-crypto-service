package handler

import (
	"github.com/devendrapratap307/go-crypto-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type RestHandler struct{ svc *service.CryptoService }

func NewRestHandler(svc *service.CryptoService) *RestHandler {
	return &RestHandler{svc: svc}
}
func (h *RestHandler) Register(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/encrypt", h.encrypt)
	api.Post("/decrypt", h.decrypt)
}

type encReq struct {
	Plaintext string `json:"plaintext"`
}

func (h *RestHandler) encrypt(c *fiber.Ctx) error {
	var r encReq
	if err := c.BodyParser(&r); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	ct, nonce, err := h.svc.Encrypt([]byte(r.Plaintext))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{"ciphertext": ct, "nonce": nonce})
}

type decReq struct {
	Ciphertext, Nonce string `json:"ciphertext"`
}

func (h *RestHandler) decrypt(c *fiber.Ctx) error {
	var r decReq
	if err := c.BodyParser(&r); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	pt, err := h.svc.Decrypt(r.Ciphertext, r.Nonce)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{"plaintext": string(pt)})
}
