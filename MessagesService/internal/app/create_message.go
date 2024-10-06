package app

import (
	"MessagesService/internal/pkg/services/store"
	"MessagesService/internal/pkg/types"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (app *Application) CreateMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req *createMessageRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		if err := validateCreateMessageRequest(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		msg := bindMessage(req)

		insertedMsg, err := app.Store.InsertMessage(msg)
		if err != nil {
			app.Logger.Errorf("failed to create message, err: %v", err)
			if errors.Is(err, store.ErrUsernameDoesNotExist) {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
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

func bindMessage(req *createMessageRequest) *types.Message {
	return types.NewMessage(uuid.Must(uuid.NewV7()), req.Username, req.Content, false)
}
