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
	link := &Link{
		Url: url,
	}
	link.GenerateHash()

	return link
}

func (link *Link) GenerateHash() {
	link.Hash = randString(6)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	myRunes := make([]rune, n)

	for i := range myRunes {
		myRunes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(myRunes)
}