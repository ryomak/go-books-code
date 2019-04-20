package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

var baseURL = "https://avex.jp/nissy/news/"

type Article struct {
	Title string
	URL   string
	Date  time.Time
}

type ArticleList []Article

func main() {
	articleList, err := getList(baseURL)
	if err != nil {
		panic(err)
	}
	err = articleList.exportCSV("output.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("finish")
}
func getList(url string) (ArticleList, error) {
	articleList := make([]Article, 0)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	doc.Find("dd").Each(func(_ int, s *goquery.Selection) {
		article := Article{}
		article.Title = s.Find("a").Text()
		URI, _ := s.Find("a").Attr("href")
		article.URL = baseURL + URI
		articleList = append(articleList, article)
	})
	doc.Find("dt").Each(func(index int, s *goquery.Selection) {
		date, _ := time.Parse("2006.01.02", s.Find("time").Text())
		articleList[index].Date = date
	})
	return articleList, nil
}

func (articleList ArticleList) exportCSV(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	converter, err := iconv.NewWriter(file, "utf-8", "sjis")
	if err != nil {
		return err
	}
	writer := csv.NewWriter(converter)
	header := []string{"日付", "タイトル", "URL"}
	writer.Write(header)
	for _, v := range articleList {
		content := []string{
			v.Date.Format("2006/01/02"),
			v.Title,
			v.URL,
		}
		writer.Write(content)
	}
	writer.Flush()
	return nil
}
