package main

import (
	"link-shorter/configs"
	"link-shorter/internal/link"
	"link-shorter/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
  db := db.NewDb(conf)

	db.AutoMigrate(&link.Link{})
}
