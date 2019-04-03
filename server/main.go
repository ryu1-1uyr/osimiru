package main

import (
    "github.com/PuerkitoBio/goquery"
    "fmt"
)

func GetPage(url string) {
     doc, _ := goquery.NewDocument(url)
     doc.Find("img").Each(func(_ int, s *goquery.Selection) {
          url, _ := s.Attr("src")
          fmt.Println(url)
     })
}

func main() {
	url := "https://twitter.com/search?q=%23%E3%83%AD%E3%82%A2%E3%83%BC%E3%83%88&src=typeahead_click"
	GetPage(url)
	fmt.Println("ここから俺のツイーと \n")
	url = "https://www.google.co.jp/search?q=%E5%A4%A2%E8%A6%8B%E3%82%8A%E3%81%82%E3%82%80&rlz=1C5CHFA_enJP788JP790&source=lnms&tbm=isch&sa=X&ved=0ahUKEwiruY_a487gAhVOMt4KHbc6DzgQ_AUIDigB&biw=1440&bih=765"
    GetPage(url)
}
