package store

import (
	"MessagesService/gen/a5/public/model"
	"MessagesService/gen/a5/public/table"
	"MessagesService/internal/dependencies"
	"MessagesService/internal/pkg/types"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var ErrMessageDoesNotExist = errors.New("message does not exist")
var ErrUsernameDoesNotExist = errors.New("username does not exist")

// Store implementation of IStore interface
type Store struct {
	logger *zap.SugaredLogger
	db     *sql.DB
}

var _ dependencies.IStore = (*Store)(nil)

// NewStore creates a new instance of Store
func NewStore(l *zap.Logger) (dependencies.IStore, error) {
	if l == nil {
		return nil, errors.New("logger must be provided")
	}

	return &Store{
		logger: l.Sugar(),
	}, nil
}

func (s *Store) Start() error {
	dsn := "postgresql://postgres:password@db:5432/a5?sslmode=disable"
	//dsn := "postgresql://postgres:password@localhost:5433/a5?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		s.logger.Fatal(fmt.Sprintf("failed to connect to database, err: %v", err))
	}

	err = db.Ping()
	if err != nil {
		s.logger.Fatal(fmt.Sprintf("failed to connect to database, err: %v", err))
	}

	s.db = db

	return nil
}

func (s *Store) Stop() error {
	return s.db.Close()
}

func (s *Store) InsertMessage(msg *types.Message) (*types.Message, error) {
	stmt := table.Messages.
		INSERT(table.Messages.AllColumns).
		MODEL(msg).
		RETURNING(table.Messages.AllColumns)

	res := &model.Messages{}
	err := stmt.Query(s.db, res)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr); pqErr.Code == "23503" {
			return nil, ErrUsernameDoesNotExist
		}

		return nil, err
	}

	return types.NewMessage(res.ID, res.Username, res.Content, res.IsLiked), nil
}

func (s *Store) LikeMessage(id uuid.UUID) (*types.Message, error) {
	stmt := table.Messages.
		UPDATE(table.Messages.IsLiked).
		SET(table.Messages.IsLiked.SET(postgres.Bool(true))).
		WHERE(postgres.AND(
			table.Messages.ID.EQ(postgres.UUID(id)),
		)).
		RETURNING(table.Messages.AllColumns)

	res := &model.Messages{}
	err := stmt.Query(s.db, res)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrMessageDoesNotExist
		}

		var pqErr *pq.Error
		if errors.As(err, &pqErr); pqErr.Code == "23503" {
			return nil, ErrUsernameDoesNotExist
		}

		return nil, err
	}

	return types.NewMessage(res.ID, res.Username, res.Content, res.IsLiked), nil
}
