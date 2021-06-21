package main

import (
	// "fmt"
	"github.com/anatolethien/forum/pkg/forum"
)

func main() {
    forum.CreateDatabase()
	forum.Server()
}
