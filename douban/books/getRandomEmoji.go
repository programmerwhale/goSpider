package main

import (
	"github.com/kyokomi/emoji/v2"
)

func GetRandomEmoji() string {
	e:=emoji.Sprint(":beer")
	emoji.Println(":beer")
	return e
}
