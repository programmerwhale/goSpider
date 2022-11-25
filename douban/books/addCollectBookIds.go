package main

import (
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goSpider/handler"
	"golang.org/x/net/context"
	"log"
	"regexp"
	"strings"
	"time"
)

var DB *sql.DB
var ctx context.Context

type BookData struct {
	BookId   string `json:"bookId"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Picture  string `json:"picture"`
	Year     string `json:"year"`
	Rating   string `json:"rating"`
	Page     string `json:"page"`
	ReadDate string `json:"readDate"`
}

func main() {
	DB = handler.InitDB()
	ch := make(chan bool)
	// go getCollectBookIds(strconv.Itoa(5*15), ch)
	bookIds := selectSql()
		for _, book := range bookIds {
		spiderBookDetail(book)
	}
	//spiderBookDetail("27104959")
	/*var books = selectBookInfo()
	for _,book:=range books{
		SyncToNotion(book)
	}*/
	<-ch
	DB.Close()
	//fmt.Println(GetRandomEmoji())
}

func getCollectBookIds(paramStr string, ch chan bool) {

	resp := handler.WebRequest("GET", "https://book.douban.com/people/153216787/collect?start=", paramStr)
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
			info := s.Find("div.info > div.ft > div > span.cart-info > span > a")
			bookId, ok := info.Attr("name")
			if ok {
				cnt := 0
				if cnt == 0 {
					bookData.BookId = bookId
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
	_, err = stmt.Exec(bookData.BookId, time.Now())
	if err != nil {
		fmt.Println("Exec fail", err)
		return false
	}
	_ = tx.Commit()
	return true

}

func selectSql() (res []string) {
	rows, err := DB.Query("select BookId from books where IsCollect=0")
	if err != nil {
		fmt.Println("Exec fail", err)
	}
	for rows.Next() {
		var bookId string
		if err := rows.Scan(&bookId); err != nil {
			fmt.Println(err)
		}
		fmt.Println(bookId)
		res = append(res, bookId)
	}
	return res
}

func selectBookInfo() (res []BookData) {
	rows, err := DB.Query("select title,author,picture,year,rating,page,readDate from books where IsAddToNotion=0")
	if err != nil {
		fmt.Println("Exec fail", err)
	}
	for rows.Next() {
		var title,author,picture,year,rating,page,readDate string
		if err := rows.Scan(&title,&author,&picture,&year,&rating,&page,&readDate); err != nil {
			fmt.Println(err)
		}
		var book BookData
		book.Title=title
		book.Author=author
		book.Picture=picture
		book.Year=year
		book.Rating=rating
		book.Page=page
		book.ReadDate=readDate

		res = append(res, book)
	}
	return res
}

func spiderBookDetail(id string) (bookData BookData) {
	resp := handler.WebRequest("GET", "https://book.douban.com/subject/", id)
	defer resp.Body.Close()
	docDetail, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("fatal err")
		log.Fatal(err)
	}
	title := docDetail.Find("#wrapper > h1 > span").Text()
	author := docDetail.Find("#info > span:nth-child(1) > a").Text()
	if len(author)==0{
		author = docDetail.Find("#info > a:nth-child(2)").Text()
		author=strings.TrimSpace(author)
	}
	picture, _ := docDetail.Find("#mainpic > a").Attr("href")
	info := docDetail.Find("#info").Text()
	year, page := infoSpite(info)
	readDateInfo := docDetail.Find("#interest_sect_level > div > span.color_gray").Text()
	readDate := readDateInfoSpite(readDateInfo)
	rating, _ := docDetail.Find("#n_rating").Attr("value")

	bookData.BookId = id
	bookData.Title = title
	bookData.Author = author
	bookData.Picture = picture
	bookData.Year = year
	bookData.Rating = rating
	bookData.Page = page
	bookData.ReadDate = readDate
	updateSql(bookData)

	//SyncToNotion(bookData)
	return
}
func updateSql(bookData BookData) bool {

	fmt.Println("UpdateSql")
	tx, err := DB.Begin()

	if err != nil {
		fmt.Println("tx fail", err)
		return false
	}
	stmt, err := tx.Prepare("update books set Title=?,Author=?,Picture=?,Year=?,Rating=?,Page=?,isCollect=1,ReadDate=? where BookId=?")
	if err != nil {
		fmt.Println("Prepare fail", err)
		return false
	}
	fmt.Println(time.Now())
	_, err = stmt.Exec(bookData.Title, bookData.Author, bookData.Picture, bookData.Year, bookData.Rating, bookData.Page, bookData.ReadDate, bookData.BookId)
	if err != nil {
		fmt.Println("Exec fail", err)
		return false
	}
	_ = tx.Commit()
	return true

}
func infoSpite(info string) (year, page string) {
	yearRe, _ := regexp.Compile(`\d{4}`)
	year = string(yearRe.Find([]byte(info)))
	pageRe, _ := regexp.Compile(`页数: -?[1-9]\d*`)
	page = string(pageRe.Find([]byte(info)))
	page = strings.Trim(page, "页数: ")
	return
}
func readDateInfoSpite(info string) (readDate string) {
	readDateRe, _ := regexp.Compile(`\d{4}-\d{2}-\d{2}`)
	readDate = string(readDateRe.Find([]byte(info)))
	return
}
