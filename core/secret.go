package core

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"

	bkgw "github.com/moby/buildkit/frontend/gateway/client"
)

// Secret is a content-addressed secret.
type Secret struct {
	ID SecretID `json:"id"`
}

func NewSecret(id SecretID) *Secret {
	return &Secret{ID: id}
}

func NewSecretFromFile(fileID FileID) (*Secret, error) {
	id, err := (&secretIDPayload{FromFile: fileID}).Encode()
	if err != nil {
		return nil, err
	}

	return NewSecret(id), nil
}

func NewSecretFromHostEnv(name string) (*Secret, error) {
	id, err := (&secretIDPayload{FromHostEnv: name}).Encode()
	if err != nil {
		return nil, err
	}

	return NewSecret(id), nil
}

// SecretID is an opaque value representing a content-addressed secret.
type SecretID string

func (id SecretID) String() string { return string(id) }

func (id SecretID) IsOldFormat() (bool, error) {
	payload, err := id.decode()
	if err != nil {
		return false, err
	}

	if payload.FromFile == "" && payload.FromHostEnv == "" {
		return false, nil
	}

	return true, nil
}

func NewSecretID(name, plaintext string) (SecretID, error) {
	digestBytes := sha256.Sum256([]byte(plaintext))

	id, err := (&secretIDPayload{Name: name, Digest: string(digestBytes[:])}).Encode()
	if err != nil {
		return "", err
	}
	return id, nil
}

// secretIDPayload is the inner content of a SecretID.
type secretIDPayload struct {
	// FromFile specifies the FileID it is based off.
	//
	// Deprecated: this shouldn't be used as it can leak secrets in the cache.
	// Use the setSecret API instead.
	FromFile FileID `json:"file,omitempty"`

	// FromHostEnv specifies the FileID it is based off.
	//
	// Deprecated: use the setSecret API instead.
	FromHostEnv string `json:"host_env,omitempty"`

	// Name specifies the arbitrary name/id of the secret.
	Name string `json:"name,omitempty"`

	// Digest represents a digest of the plaintext of the secret.
	Digest string `json:"digest,omitempty"`
}

// Encode returns the opaque string ID representation of the secret.
func (payload *secretIDPayload) Encode() (SecretID, error) {
	id, err := encodeID(payload)
	if err != nil {
		return "", err
	}

	return SecretID(id), nil
}

func (id SecretID) decode() (*secretIDPayload, error) {
	var payload secretIDPayload
	if err := decodeID(&payload, id); err != nil {
		return nil, err
	}

	return &payload, nil
}

func (secret *Secret) Plaintext(ctx context.Context, gw bkgw.Client) ([]byte, error) {
	payload, err := secret.ID.decode()
	if err != nil {
		return nil, err
	}

	if payload.FromFile != "" {
		file := &File{ID: payload.FromFile}
		return file.Contents(ctx, gw)
	}

	if payload.FromHostEnv != "" {
		return []byte(os.Getenv(payload.FromHostEnv)), nil
	}

	return nil, fmt.Errorf("plaintext: empty secret?")
}
