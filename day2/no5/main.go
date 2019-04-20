package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

var baseURL = "https://www.amazon.co.jp/gp/top-sellers/books"

type Book struct {
	Rank int
	Name string
	URL  string
}

type BookList []Book

func main() {
	bookList, err := getList(baseURL)
	if err != nil {
		panic(err)
	}
	fmt.Println(bookList)
}

func getList(url string) ([]Book, error) {
	bookList := make([]Book, 0)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	fmt.Println(doc.Text())
	doc.Find("zg_itemRow").Each(func(_ int, s *goquery.Selection) {
		book := new(Book)
		fmt.Println(s.Text())
		s.Find("a-link-normal").Each(func(_ int, alink *goquery.Selection) {
			url, _ := alink.Attr("src")
			fmt.Println(url)
			book.URL = url
			book.Name = alink.Text()
		})
	})
	return bookList, nil
}

func (bookList BookList) exportCSV(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	converter, err := iconv.NewWriter(file, "utf-8", "sjis")
	if err != nil {
		return err
	}
	writer := csv.NewWriter(converter)
	header := []string{"順位", "名前", "URL"}
	writer.Write(header)
	for _, v := range bookList {
		content := []string{
			string(v.Rank),
			v.Name,
			v.URL,
		}
		writer.Write(content)
	}
	writer.Flush()
	return nil
}
