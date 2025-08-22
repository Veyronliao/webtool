package crypto

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
	"hash"
)

const (
	defaultSaltLen    = 64
	defaultIterations = 10000
	defaultKeyLen     = 128
)

var defaultHashFunction = sha512.New

type Options struct {
	SaltLen      int
	Iterations   int
	KeyLen       int
	HashFunction func() hash.Hash
}

// 生产salt
func generateSalt(length int) []byte {
	const alphaNum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	salt := make([]byte, length)
	rand.Read(salt)
	for key, val := range salt {
		salt[key] = alphaNum[val%byte(len(alphaNum))]
	}
	return salt
}

// Encode
func Encode(pwd string, options *Options) (string, string) {
	if options == nil {
		salt := generateSalt(defaultSaltLen)
		encodepwd := pbkdf2.Key([]byte(pwd), salt, defaultIterations, defaultKeyLen, defaultHashFunction)
		return hex.EncodeToString(encodepwd), string(salt)
	} else {
		salt := generateSalt(options.SaltLen)
		encodepwd := pbkdf2.Key([]byte(pwd), salt, options.Iterations, options.KeyLen, options.HashFunction)
		return hex.EncodeToString(encodepwd), string(salt)
	}
}

// verify
func Verify(rawpwd string, encodepwd string, salt string, options *Options) bool {
	if options == nil || options.HashFunction == nil {
		return encodepwd == hex.EncodeToString(pbkdf2.Key([]byte(rawpwd), []byte(salt), defaultIterations, defaultKeyLen, defaultHashFunction))
	} else {
		return encodepwd == hex.EncodeToString(pbkdf2.Key([]byte(rawpwd), []byte(salt), options.Iterations, options.KeyLen, options.HashFunction))
	}
}
