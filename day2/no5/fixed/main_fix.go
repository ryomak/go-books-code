package main

import (
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/olekukonko/tablewriter"
	"github.com/sclevine/agouti"
)

var baseURL = "https://www.amazon.co.jp"

//"https://www.amazon.co.jp/gp/top-sellers/books"

type Book struct {
	Rank  int
	Name  string
	URL   string
	Price string
}
type BookList []Book

func main() {
	bookList, err := getList(baseURL + "/gp/top-sellers/books")
	if err != nil {
		panic(err)
	}
	bookList.exportTable()
}

func getList(url string) (BookList, error) {
	bookList := make([]Book, 0)
	//
	// brew install chromedriver
	// ブラウザで表示されたページから取得
	//
	driver := agouti.ChromeDriver()
	err := driver.Start()
	if err != nil {
		return nil, err
	}
	defer driver.Stop()
	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		return nil, err
	}

	err = page.Navigate(url)
	if err != nil {
		return nil, err
	}

	page.FindByButton("Next").Click()

	content, err := page.HTML()
	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(content)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	doc.Find("div.a-fixed-left-grid-col.a-col-right").Each(func(index int, s *goquery.Selection) {
		book := Book{Rank: (index + 1)}
		alink := s.Find("a").First()
		url, _ := alink.Attr("href")
		book.URL = baseURL + url
		title, ok := alink.Find(".p13n-sc-truncated").Attr("title")
		if !ok {
			book.Name = alink.Find(".p13n-sc-truncated").Text()
		} else {
			book.Name = title
		}
		book.Price = strings.Replace(s.Find(".p13n-sc-price").Text(), "￥ ", "", -1)
		bookList = append(bookList, book)
	})
	return bookList, nil
}

func (bookList BookList) exportTable() error {
	table := tablewriter.NewWriter(os.Stdout)
	/*
		table.SetHeader([]string{"順位", "名前", "値段", "URL"})
		for _, v := range bookList {
			table.Append([]string{
				string(v.Rank),
				v.Name,
				v.Price,
				v.URL,
			})
		}
	*/
	table.SetHeader([]string{"名前", "値段"})
	for _, v := range bookList {
		table.Append([]string{
			v.Name,
			v.Price,
		})
	}
	table.Render()
	return nil
}
