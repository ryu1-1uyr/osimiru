package main

import (
  "github.com/PuerkitoBio/goquery"
  "fmt"
  "net/http"
  "encoding/json"
)

type Pagedata struct { //jsonの構造
	URL   []string
}

func GetPage(url string) ([]string) {
  var array []string
  doc, _ := goquery.NewDocument(url)
  doc.Find("img").Each(func(_ int, s *goquery.Selection) {
    url, _ := s.Attr("src")
    array = append(array,url)
  })

  return array
}

func handler(w http.ResponseWriter, r *http.Request) {
  url := "https://twitter.com/search?q=%23%E3%83%AD%E3%82%A2%E3%83%BC%E3%83%88&src=typeahead_click"//任意のurl取れるように改造したい
  pagedata := GetPage(url)

  pages := Pagedata{pagedata}

	res, err := json.Marshal(pages)

  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(res)

}

func main() {

  http.HandleFunc("/", handler)       // http://localhost:8080/にアクセスしてきた人はhandlerを実行するよ！
  fmt.Printf("server is running\n")
  http.ListenAndServe(":8080", nil)   // サーバーを起動するよ！
}
