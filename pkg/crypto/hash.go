package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

type HashAlgorithm string

const (
	SHA256 HashAlgorithm = "sha256"
	SHA512 HashAlgorithm = "sha512"
)

type Hasher struct {
	algorithm HashAlgorithm
	hash      hash.Hash
}

func NewHasher(algorithm HashAlgorithm) *Hasher {
	var h hash.Hash
	switch algorithm {
	case SHA256:
		h = sha256.New()
	case SHA512:
		h = sha512.New()
	default:
		h = sha256.New()
	}

	return &Hasher{
		algorithm: algorithm,
		hash:     h,
	}
}

// Hash è®¡ç®æ°æ®çåå¸å?
func (h *Hasher) Hash(data []byte) string {
	h.hash.Reset()
	h.hash.Write(data)
	return hex.EncodeToString(h.hash.Sum(nil))
}

// HashMultiple è®¡ç®å¤ä¸ªæ°æ®çç»ååå¸?
func (h *Hasher) HashMultiple(data ...[]byte) string {
	h.hash.Reset()
	for _, d := range data {
		h.hash.Write(d)
	}
	return hex.EncodeToString(h.hash.Sum(nil))
}

// VerifyHash éªè¯åå¸å?
func (h *Hasher) VerifyHash(data []byte, hash string) bool {
	computedHash := h.Hash(data)
	return computedHash == hash
} 
