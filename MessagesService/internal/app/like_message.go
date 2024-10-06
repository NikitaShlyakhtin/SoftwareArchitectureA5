package app

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (app *Application) LikeMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req *likeMessageRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid request body, err: %w", err))
		}

		if err := validateLikeMessageRequest(req); err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid request body, err: %w", err))
		}

		likedMessage, err := app.Store.LikeMessage(c.Request().Context(), req.Username, uuid.MustParse(req.ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Errorf("failed to like message, err: %w", err))
		}

		return c.JSON(http.StatusOK, likedMessage)
	}
}

type likeMessageRequest struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

func validateLikeMessageRequest(req *likeMessageRequest) error {
	if req.Username == "" {
		return errors.New("username must be provided")
	}

	if req.ID == "" {
		return errors.New("message ID must be provided")
	}

	if _, err := uuid.Parse(req.ID); err != nil {
		return errors.New("message ID must be a valid UUID")
	}

	return nil
}
