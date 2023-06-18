package algo

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/letusgogo/nopass/config"
)

type Algorithm interface {
	Generate(input string) string
}

func NewAlgorithmFromConfig(c *config.AlgoConfig) (Algorithm, error) {
	if c == nil {
		return nil, errors.New("nil config")
	}
	return &SHA256Algo{Slat: c.Salt}, nil
}

type SHA256Algo struct {
	Slat string
}

func (s *SHA256Algo) Generate(input string) string {
	hash := sha256.Sum256([]byte(input + s.Slat))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}
