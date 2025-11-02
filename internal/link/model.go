package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	return &Link{
		Url: url,
		Hash: generateHash(6),
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateHash(n int) string {
	myHash := make([]rune, n)

	for i := range myHash {
		myHash[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(myHash)
}