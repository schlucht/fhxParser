package models

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/google/uuid"
)

const (
	ScopeAuthentication = "authentication"
	ScopePasswordUpdate = "password_update"
	ScopeDeletion       = "deletion"
)

type Token struct {
	Plaintext string    `json:"token"`
	Hash      []byte    `json:"-"`
	UserID    string    `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Scope     string    `json:"-"`
}

// GenerateToken generates a new token with a specific userID and duration
func GenerateToken(userId string, ttl time.Duration, scope string) (*Token, error) {

	token := &Token{
		UserID: userId,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	token.Plaintext = base64.URLEncoding.EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]
	return token, nil
}

func (t *Token) IsExpired() bool {
	return time.Now().After(t.Expiry)
}

func (m *DBModel) InsertToken(token *Token, u User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.deleteToken(u.ID)
	if err != nil {
		return err
	}
	id := uuid.New()

	stmt := `INSERT INTO tokens (id, user_id, name, email, token_hash) VALUES(?,?,?,?,?)`
	_, err = m.DB.ExecContext(ctx, stmt, id.String(), u.ID, u.Name, u.Email, token.Hash)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) deleteToken(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `DELETE FROM tokens WHERE user_id = ?`
	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}
