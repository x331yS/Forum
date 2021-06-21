package forum

import (
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(clear string) string {
	var hash, _ = bcrypt.GenerateFromPassword([]byte(clear), 2)
	return string(hash)
}

func CheckPasswordHash(hash string, clear string) bool {
    var err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(clear))
    if err == nil {
        return true
    } else {
        return false
    }
}
