package responder

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type response struct {
	StatusCode   int              `json:"status_code"`
	StatusString string           `json:"status_string"`
	Message      *string          `json:"message,omitempty"`
	Data         *json.RawMessage `json:"data,omitempty"`
}

// 200
func OkWithData(c *fiber.Ctx, data any, message ...any) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return InternalServerError(c)
	}

	raw := json.RawMessage(jsonBytes)

	return respond(c, fiber.StatusOK, &raw, message...)
}

func Ok(c *fiber.Ctx, message ...any) error {
	return respond(c, fiber.StatusOK, nil, message...)
}

// 400
func BadRequest(c *fiber.Ctx, message ...any) error {
	return respond(c, fiber.StatusBadRequest, nil, message...)
}

// 401
func Unauthorized(c *fiber.Ctx, message ...any) error {
	return respond(c, fiber.StatusUnauthorized, nil, message...)
}

// 403
func Forbidden(c *fiber.Ctx, message ...any) error {
	return respond(c, fiber.StatusForbidden, nil, message...)
}

func InternalServerError(c *fiber.Ctx, message ...any) error {
	return respond(c, fiber.StatusInternalServerError, nil, message...)
}

func NotYetImplemented(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func respond(c *fiber.Ctx, statusCode int, data *json.RawMessage, message ...any) error {
	msg := utils.StatusMessage(statusCode)
	if len(message) > 0 {
		format, args := message[0].(string), message[1:]
		msg = fmt.Sprintf(format, args...)
	}

	res := response{
		StatusCode:   statusCode,
		StatusString: utils.StatusMessage(statusCode),
		Message:      &msg,
		Data:         data,
	}

	return c.Status(statusCode).JSON(res)
}
