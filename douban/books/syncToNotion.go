package main

import (
	"encoding/json"
	"fmt"
	"goSpider/handler"
	"strconv"
)

func SyncToNotion(data BookData)  {
	Config:=handler.ViperHelper()
	dbId:=fmt.Sprintf("%v",Config.Get("notion.bookTrackerDatabaseId"))
	url:=fmt.Sprintf("%v",Config.Get("notion.url"))

	// 写入新pages
	res, err := CreateBookTrackerPages(dbId, url,data)
	if err != nil {
		fmt.Println(err)
	}

	var f interface{}
	json.Unmarshal(res, &f)
	status := f.(map[string]interface{})["status"]
	fmt.Println(status)
/*	if !IsNil(&status){
		fmt.Println("已生成")
	}else if status.(string)=="400"{
		fmt.Println("生成失败！")
	}*/
}

func CreateBookTrackerPages(databaseid,url string,model BookData) ([]byte, error) {
	data := make(map[string]interface{})
	properties := make(map[string]interface{})

	cover := make(map[string]interface{})
	cover["type"]="external"
	pictureUrl:=make(map[string]interface{})
	pictureUrl["url"]=model.Picture
	cover["external"]=pictureUrl

	icon := make(map[string]interface{})
	icon["type"]="emoji"
	// todo: emoji
	icon["emoji"]="📚"

	parent := make(map[string]interface{})
	parent["type"] = "database_id"
	parent["database_id"] = databaseid

	pgsRead:=make(map[string]int)
	pgsRead["number"],_=strconv.Atoi(model.Page)

	rating:=make(map[string]interface{})
	rating["id"]="HhrT"
	rating["type"]="select"
	selects :=make(map[string]interface{})
	selects["id"]="cf1e9719-e70d-402c-bba7-b78e54aa804e"
	selects["color"]="yellow"
	stars:=""
	star, _ :=strconv.Atoi(model.Rating)
	for i:=1;i<=star;i++{
		stars+="⭐️"
	}
	selects["name"]=stars
	rating["select"]= selects

	relations:=	make(map[string]interface{})
	relations["id"]="XWCQ"
	relations["has_more"]=false
	relation:=make([]map[string]interface{},1)
	relation1:=make(map[string]interface{})
	relation1["id"]="95b98d81-6456-4084-bd68-d54aec89a483"
	relation[0]=relation1
	relations["type"]="relation"
	relations["relation"]=relation

	totalPgs:=make(map[string]int)
	totalPgs["number"],_=strconv.Atoi(model.Page)

	status:=make(map[string]interface{})
	status["id"]="itTY"
	status["type"]="select"
	statusSelect :=make(map[string]interface{})
	statusSelect["id"]="1b85a5a5-9001-4390-a576-12c7e269bb11"
	statusSelect["name"]="🔘To Summary"
	statusSelect["color"]="purple"
	status["select"]= statusSelect

	dateRead:=make(map[string]interface{})
	dateRead["id"]="lXcT"
	dateRead["type"]="date"
	d :=make(map[string]interface{})
	d["start"]=model.ReadDate
	dateRead["date"]= d

	//author
	author:=make(map[string]interface{})
	author["type"]="rich_text"
	authorRichText:=make([]map[string]interface{},1)
	authorRichText1:=make(map[string]interface{})
	authorRichText1["type"]="text"
	authorText:=make(map[string]interface{})
	authorText["content"]=model.Author
	authorRichText1["text"]=authorText
	authorRichText1["plain_text"]=model.Author
	authorRichText[0]=authorRichText1
	author["rich_text"]=authorRichText

	priority:=make(map[string]interface{})
	priority["id"]="seAH"
	priority["type"]="select"
	prioritySelect :=make(map[string]interface{})
	prioritySelect["id"]="mnCc"
	prioritySelect["name"]="📘"
	prioritySelect["color"]="default"
	priority["select"]= prioritySelect

	year:=make(map[string]int)
	year["number"],_=strconv.Atoi(model.Year)

	// Title
	title:=make(map[string]interface{})
	title["id"]="title"
	title["type"]="title"
	titleTitle:=make([]map[string]interface{},1)
	titleTitle1:=make(map[string]interface{})
	titleTitle1["type"]="text"
	titleText:=make(map[string]interface{})
	titleText["content"]=model.Title
	titleTitle1["text"]=titleText
	titleTitle1["plain_text"]=model.Title
	titleTitle[0]=titleTitle1
	title["title"]=titleTitle

	properties["Pgs Read"] = pgsRead
	properties["Rating"] = rating
	properties["To: Total Summary [MUST TAG EVERY BOOK WITH THIS]"] = relations
	properties["Total Pgs"] = totalPgs
	properties["Status"] = status
	properties["Date Read"] = dateRead
	properties["Author"] = author
	properties["Priority"] = priority
	properties["Year"] = year
	properties["Title"] = title

	data["cover"] = cover
	data["icon"] = icon
	data["parent"] = parent
	data["properties"] = properties

	bytesData, _ := json.Marshal(data)
	fmt.Println(string(bytesData))
	return handler.PostHttpsSkip(url+"/pages/", bytesData)
}

