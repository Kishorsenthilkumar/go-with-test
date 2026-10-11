package data

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"time"

	"greenlight.alexedwards.net/internal/data/validator"
)

const (
	ScopeActivation = "activation"
)

type Token struct {
	Plaintext string
	Hash      []byte
	UserID    int64
	Scope     string
	Expiry    time.Time
}

func generateToken(UserID int64, TTL time.Duration, Scope string) (*Token, error) {

	token := &Token{
		UserID: UserID,
		Scope:  Scope,
		Expiry: time.Now().Add(TTL),
	}

	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]

	return token, nil
}

func ValidateTokenPlaintext(v *validator.Validator, tokenPlaintext string) {
	v.Check(tokenPlaintext != "", "token", "must be provided")
	v.Check(len(tokenPlaintext) == 26, "token", "must be 26 bytes long")
}

type TokenModel struct {
	DB *sql.DB
}

func (t *TokenModel) New(UserID int64, TTL time.Duration, Scope string) (*Token, error) {
	token, err := generateToken(UserID, TTL, Scope)
	if err != nil {
		return nil, err
	}

	err = t.Insert(token)
	return token, err
}

func (t *TokenModel) Insert(token *Token) error {

	query := `
        INSERT INTO tokens (hash, user_id, expiry, scope) 
        VALUES ($1, $2, $3, $4)`

	args := []interface{}{token.Hash, token.UserID, token.Expiry, token.Scope}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := t.DB.ExecContext(ctx, query, args...)
	return err
}

func (t *TokenModel) DeleteAllForUser(Scope string, UserID int64) error {
	query := `
        DELETE FROM tokens 
        WHERE scope = $1 AND user_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := t.DB.ExecContext(ctx, query, Scope, UserID)
	return err
}
