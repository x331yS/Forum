package main

import (
	"github.com/anatolethien/forum/pkg/forum"
)

func main() {
	forum.CreateDatabase()
	forum.Static()
}
