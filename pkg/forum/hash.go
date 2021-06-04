package forum

import (
    "golang.org/x/crypto/bcrypt"
    "encoding/base64"
)

func HashPassword(clearPasswordString string) string {
    var clearPasswordBytes = []byte(clearPasswordString)
    var hashedPasswordBytes, _ = bcrypt.GenerateFromPassword(clearPasswordBytes, bcrypt.MinCost)
    var hashedPasswordString = base64.URLEncoding.EncodeToString(hashedPasswordBytes)
    return hashedPasswordString
}

// func HashPassword2(clear string) []byte {
//     var hash = sha256.Sum256([]byte(clear))
//     return hash[:]
// }
