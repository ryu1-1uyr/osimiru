package main

import (
  "github.com/PuerkitoBio/goquery"
  "fmt"
  "net/http"

)

func GetPage(url string) ([]string) {
  var array []string
  doc, _ := goquery.NewDocument(url)
  doc.Find("img").Each(func(_ int, s *goquery.Selection) {
    url, _ := s.Attr("src")
    // fmt.Println(url)
    array = append(array,url)
  })

  return array
}

func handler(w http.ResponseWriter, r *http.Request) {
  url := "https://twitter.com/search?q=%23%E3%83%AD%E3%82%A2%E3%83%BC%E3%83%88&src=typeahead_click"
  pagedata := GetPage(url)
  fmt.Fprintf(w, "pagedata")      // Hello, Worldってアクセスした人に返すよ！
  fmt.Println(pagedata)
}

func main() {

  // url = "https://www.google.co.jp/search?q=%E5%A4%A2%E8%A6%8B%E3%82%8A%E3%81%82%E3%82%80&rlz=1C5CHFA_enJP788JP790&source=lnms&tbm=isch&sa=X&ved=0ahUKEwiruY_a487gAhVOMt4KHbc6DzgQ_AUIDigB&biw=1440&bih=765"
  // GetPage(url)

  http.HandleFunc("/", handler)       // http://localhost:8080/にアクセスしてきた人はhandlerを実行するよ！
  fmt.Printf("server is running\n")
  http.ListenAndServe(":8080", nil)   // サーバーを起動するよ！
}
