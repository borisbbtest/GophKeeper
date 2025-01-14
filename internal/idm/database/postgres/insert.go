package postgres

import (
	"context"
	"embed"
	"time"

	"github.com/jackc/pgconn"
	"golang.org/x/crypto/bcrypt"

	"github.com/borisbbtest/GoMon/internal/idm/configs"
	"github.com/borisbbtest/GoMon/internal/idm/service"
	pb "github.com/borisbbtest/GoMon/internal/models/idm"
)

// Файлы SQL для вставки записей в таблицы хранятся в директории migrations/insert/
//
//go:embed migrations/insert/*.sql
var SQLInsert embed.FS

// CreateUser - функция создает пользователя, полученного по gRPC
func (r *IdmRepo) CreateUser(ctx context.Context, cfg *configs.AppConfig, user *pb.User) error {
	sqlBytes, err := SQLInsert.ReadFile("migrations/insert/SQLInsertNewUser.sql")
	if err != nil {
		return err
	}
	sqlQuery := string(sqlBytes)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}
	_, err = r.Pool.Exec(ctx, sqlQuery, user.Login, user.Firstname, user.Lastname, string(hashedPassword), user.Source, time.Now())
	if err != nil {
		pgerr := err.(*pgconn.PgError)
		if pgerr.Code == "23505" {
			return service.ErrUserExists
		}
		return err
	}
	return nil
}

// CreateSession - функция создает сессию, полученную по gRPC
func (r *IdmRepo) CreateSession(ctx context.Context, cfg *configs.AppConfig, session *pb.Session) error {
	sqlBytes, err := SQLInsert.ReadFile("migrations/insert/SQLInsertSession.sql")
	if err != nil {
		return err
	}
	sqlQuery := string(sqlBytes)
	if err != nil {
		return err
	}
	_, err = r.Pool.Exec(ctx, sqlQuery, session.Id, session.Login, session.Created.AsTime(), session.Duration.AsTime())
	if err != nil {
		return err
	}
	return nil
}
