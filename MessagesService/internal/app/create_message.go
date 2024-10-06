package app

import (
	"MessagesService/internal/pkg/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (app *Application) CreateMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req *createMessageRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid request body, err: %w", err))
		}

		if err := validateCreateMessageRequest(req); err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid request body, err: %w", err))
		}

		msg := bindMessage(req)

		insertedMsg, err := app.Store.InsertMessage(c.Request().Context(), msg)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Errorf("failed to insert new message, err: %w", err))
		}

		return c.JSON(http.StatusCreated, insertedMsg)
	}
}

type createMessageRequest struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func validateCreateMessageRequest(req *createMessageRequest) error {
	if req.Username == "" {
		return errors.New("username must be provided")
	}

	if req.Content == "" {
		return errors.New("content must be provided")
	}

	if len(req.Content) > 400 {
		return errors.New("content must not be longer than 400 characters")
	}

	return nil
}

func bindMessage(req *createMessageRequest) *models.Message {
	return models.NewMessage(uuid.Must(uuid.NewV7()), req.Username, req.Content, false)
}
