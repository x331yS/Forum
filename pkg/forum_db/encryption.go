package forum_db

import (
  "crypto/sha256"
)

func EncryptPassword(clear string) []byte {
  var hash = sha256.Sum256([]byte(clear))
  return hash[:]
}
