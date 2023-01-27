package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func CheckPasswordHash(senha string, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(senha))
}
