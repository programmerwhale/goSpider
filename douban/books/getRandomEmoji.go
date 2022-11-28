package main

import (
	"github.com/kyokomi/emoji/v2"
	"math/rand"
	"time"
)

func GetRandomEmoji() string {
	arr:=[]string{":two_hearts:",":love_letter:",":tulip:",":bouquet:",":closed_book:",":green_book:",":notebook:",":books:",":notebook_with_decorative_cover:",":bookmark:",":newspaper:"}
	i:=getRandomWithAll(0,len(arr)-1)
	e:=emoji.Sprint(arr[i])
	/*emoji.Println(":beer")*/
	return e
}

func getRandomWithAll(min,max int)int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1+min)
}