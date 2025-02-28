package postgreSQL

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/RamzassH/LeadIt/authService/internal/domain/models"
	"github.com/RamzassH/LeadIt/authService/internal/storage"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

type Storage struct {
	db *sql.DB
}

func (s *Storage) UserById(ctx context.Context, id int64) (user models.User, err error) {
	const op = "storage.UserByID"

	stmt, err := s.db.Prepare(`SELECT id, email, is_admin FROM users WHERE id = $1`)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(&user.ID, &user.Email, &user.PassHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, err)
		}
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	return user, nil
}

func New(db *sql.DB) (*Storage, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUser(ctx context.Context, user models.User) (int64, error) {
	const op = "storage.SaveUser"

	query := `
        INSERT INTO users(name, surname, email, password_hash)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `

	var id int64
	err := s.db.QueryRowContext(ctx, query, user.Name, user.Surname, user.Email, user.PassHash).Scan(&id)
	if err != nil {
		var postgresError *pq.Error
		if errors.As(err, &postgresError) {
			switch postgresError.Code.Name() {
			case "unique_violation":
				return 0, storage.ErrUserExists
			}
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) UpdateUser(ctx context.Context, updatePayload models.User) error {
	const op = "storage.updateUser"
	query := `UPDATE users 
			SET
			    email = $1,
			    password_hash = $2,
			    name = $3,
			    surname = $4,
			    middle_name = $5,
			    birth_date = $6
				about_me = $7,
			    messengers=$8,
			WHERE id = $9`

	res, err := s.db.ExecContext(ctx, query,
		updatePayload.Email,
		updatePayload.PassHash,
		updatePayload.Name,
		updatePayload.Surname,
		updatePayload.MiddleName,
		updatePayload.BirthDate,
		updatePayload.AboutMe,
		updatePayload.Messengers,
		updatePayload.ID)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("%s: user not found, id=%d", op, updatePayload.ID)
	}

	return nil
}
func (s *Storage) VerifyUser(ctx context.Context, userID int64) error {
	const op = "storage.verifyUser"

	query := `UPDATE users 
				SET 
				    is_verified = $1
				    WHERE id = $2`

	res, err := s.db.ExecContext(ctx, query, true, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("%s: user not found", op)
	}
	return nil
}

func (s *Storage) UserByEmail(ctx context.Context, email string) (models.User, error) {
	const op = "storage.UserByEmail"

	stmt, err := s.db.Prepare("SELECT id, email, password_hash FROM users WHERE email = $1")
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, email)

	var user models.User
	err = row.Scan(&user.ID, &user.Email, &user.PassHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, err)
		}

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (s *Storage) IsAdmin(ctx context.Context, uid int64) (bool, error) {
	const op = "storage.IsAdmin"

	stmt, err := s.db.Prepare("SELECT isAdmin FROM users WHERE id = $1")
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, uid)
	var isAdmin bool

	err = row.Scan(&isAdmin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("%s: %w", op, err)
		}
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return isAdmin, nil
}

// SaveRefreshToken сохраняет новый refreshToken
func (s *Storage) SaveRefreshToken(ctx context.Context, token models.RefreshToken) error {
	const op = "storage.SaveRefreshToken"

	_, err := s.db.ExecContext(ctx, `
		INSERT INTO refresh_tokens (token, user_id, expires_at, revoked, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`, token.Token, token.UserID, token.ExpiresAt, token.Revoked, token.CreatedAt)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// GetRefreshToken получает refreshToken из базы данных
func (s *Storage) GetRefreshToken(ctx context.Context, token string) (models.RefreshToken, error) {
	const op = "storage.GetRefreshToken"

	var rt models.RefreshToken
	err := s.db.QueryRowContext(ctx, `
		SELECT token, user_id, expires_at, revoked, created_at
		FROM refresh_tokens
		WHERE token = $1
	`, token).Scan(&rt.Token, &rt.UserID, &rt.ExpiresAt, &rt.Revoked, &rt.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.RefreshToken{}, fmt.Errorf("token not found")
		}
		return models.RefreshToken{}, fmt.Errorf("%s: %w", op, err)
	}

	return rt, nil
}

// RevokeRefreshToken отзывает refreshToken
func (s *Storage) RevokeRefreshToken(ctx context.Context, token string) error {
	const op = "storage.RevokeRefreshToken"

	_, err := s.db.ExecContext(ctx, `
		UPDATE refresh_tokens
		SET revoked = TRUE
		WHERE token = $1
	`, token)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
