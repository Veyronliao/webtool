package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

func Md5(s string) string {
	return Hash(s, "md5", false)
}

func MD5(s string) string {
	return Hash(s, "MD5", true)
}

func Sha1(s string) string {
	return Hash(s, "sha1", false)
}

func SHA1(s string) string {
	return Hash(s, "SHA1", true)
}

func Sha224(s string) string {
	return Hash(s, "sha224", false)
}

func SHA224(s string) string {
	return Hash(s, "SHA224", true)
}
func Sha256(s string) string {
	return Hash(s, "sha256", false)
}

func SHA256(s string) string {
	return Hash(s, "SHA256", true)
}

func Sha384(s string) string {
	return Hash(s, "sha384", false)
}

func SHA384(s string) string {
	return Hash(s, "SHA384", true)
}

func Sha512(s string) string {
	return Hash(s, "sha512", false)
}

func SHA512(s string) string {
	return Hash(s, "SHA512", true)
}

func Hash(s string, name string, capital bool) string {
	var h hash.Hash
	switch name {
	case "md5", "MD5":
		h = md5.New()
	case "sha1", "SHA1":
		h = sha1.New()
	case "sha224", "SHA224":
		h = sha256.New224()
	case "sha256", "SHA256":
		h = sha256.New()
	case "sha384", "SHA384":
		h = sha512.New384()
	case "sha512", "SHA512":
		h = sha512.New()
	default:
		return ""
	}
	h.Write([]byte(s))
	if capital {
		return fmt.Sprintf("%X", h.Sum(nil))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
