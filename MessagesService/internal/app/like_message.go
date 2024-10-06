package app

import (
	"MessagesService/internal/pkg/services/store"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (app *Application) LikeMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req *likeMessageRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		if err := validateLikeMessageRequest(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		likedMessage, err := app.Store.LikeMessage(uuid.MustParse(req.ID))
		if err != nil {
			app.Logger.Errorf("failed to like message, err: %v", err)
			if errors.Is(err, store.ErrUsernameDoesNotExist) {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
		}

		return c.JSON(http.StatusOK, likedMessage)
	}
}

type likeMessageRequest struct {
	ID string `json:"id"`
}

func validateLikeMessageRequest(req *likeMessageRequest) error {
	if req.ID == "" {
		return errors.New("message ID must be provided")
	}

	if _, err := uuid.Parse(req.ID); err != nil {
		return errors.New("message ID must be a valid UUID")
	}

	return nil
}
