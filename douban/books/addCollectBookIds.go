package main

import (
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goSpider/handler"
	"golang.org/x/net/context"
	"log"
	"strconv"
	"time"
)

var DB *sql.DB
var ctx context.Context

type BookData struct {
	BookId   string `json:"bookId"`
	Title    string `json:"title"`
	Author 	 string `json:"author"`
	Picture  string `json:"picture"`
	Year     string `json:"year"`
	Rating   string `json:"rating"`
	RateWord   string `json:"rateWord"`
	Page	 string `json:"page"`
	Quote    string `json:"quote"`
}

func main() {
	DB=handler.InitDB()
	ch := make(chan bool)
	go getCollectBookIds(strconv.Itoa(1*15), ch)
	bookIds:=selectSql()
	updateBookDetail(bookIds)
	<-ch
	DB.Close()
}

func getCollectBookIds(paramStr string, ch chan bool) {

	resp:=handler.WebRequest("GET","https://book.douban.com/people/153216787/collect?start=",paramStr)
	defer resp.Body.Close()
	docDetail, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("fatal err")
		log.Fatal(err)
	}

	// document.querySelector("#content > div.grid-16-8.clearfix > div.article > ul > li:nth-child(1) > div.info > div.ft > div > span.cart-info > span > a")
	docDetail.Find("#content > div.grid-16-8.clearfix > div.article > ul > li"). //定位到html页面指定元素
		Each(func(i int, s *goquery.Selection) { //循环遍历每一个指定元素
			var bookData BookData //实例化结构体
			info:=s.Find("div.info > div.ft > div > span.cart-info > span > a")
			bookId,ok := info.Attr("name")
			if ok{
				cnt:=0
				if cnt==0{
					bookData.BookId=bookId
					insertSql(bookData)
				}
			}
		})

	if ch != nil {
		ch <- true
	}
}

func insertSql(bookData BookData) bool {

	fmt.Println("InsertSql")
	tx, err := DB.Begin()

	if err != nil {
		fmt.Println("tx fail", err)
		return false
	}
	stmt, err := tx.Prepare("INSERT INTO books(`BookId`,`CreatedAt`) VALUES (?,?)")
	if err != nil {
		fmt.Println("Prepare fail", err)
		return false
	}
	fmt.Println(time.Now())
	_, err = stmt.Exec(bookData.BookId,time.Now())
	if err != nil {
		fmt.Println("Exec fail", err)
		return false
	}
	_ = tx.Commit()
	return true

}

func selectSql() (res []string){
	rows, err :=DB.Query("select BookId from books where IsAddToNotion=0")
	if err != nil {
		fmt.Println("Exec fail", err)
	}
	for rows.Next(){
		var bookId string
		if err:=rows.Scan(&bookId);err!=nil{
			fmt.Println(err)
		}
		fmt.Println(bookId)
		res=append(res,bookId)
	}
	return res
}
func updateBookDetail(ids []string) bool{
	return false
}