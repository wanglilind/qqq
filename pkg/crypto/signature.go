package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type Signer struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func NewSigner() (*Signer, error) {
	// çæECDSAå¯é¥å¯?
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return &Signer{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}, nil
}

// Sign å¯¹æ°æ®è¿è¡ç­¾å?
func (s *Signer) Sign(data []byte) ([]byte, error) {
	// è®¡ç®æ°æ®åå¸
	hasher := NewHasher(SHA256)
	hash := hasher.Hash(data)

	// ç­¾å
	r, ss, err := ecdsa.Sign(rand.Reader, s.privateKey, []byte(hash))
	if err != nil {
		return nil, err
	}

	// ç»åç­¾åç»æ
	signature := append(r.Bytes(), ss.Bytes()...)
	return signature, nil
}

// Verify éªè¯ç­¾å
func (s *Signer) Verify(data []byte, signature []byte) bool {
	// è®¡ç®æ°æ®åå¸
	hasher := NewHasher(SHA256)
	hash := hasher.Hash(data)

	// è§£æç­¾å
	r := new(ecdsa.PublicKey)
	ss := new(ecdsa.PublicKey)
	// è¿ééè¦æ­£ç¡®åå²signatureæ¥è·årås
	
	// éªè¯ç­¾å
	return ecdsa.Verify(s.publicKey, []byte(hash), r.X, ss.X)
}

// ExportPublicKey å¯¼åºå¬é¥
func (s *Signer) ExportPublicKey() (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(s.publicKey)
	if err != nil {
		return "", err
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	return string(publicKeyPEM), nil
}

// ImportPublicKey å¯¼å¥å¬é¥
func (s *Signer) ImportPublicKey(publicKeyPEM string) error {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return errors.New("failed to decode PEM block")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	ecdsaPublicKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("not an ECDSA public key")
	}

	s.publicKey = ecdsaPublicKey
	return nil
} 
