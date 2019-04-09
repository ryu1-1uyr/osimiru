package main

import (
  "github.com/PuerkitoBio/goquery"
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "github.com/bitly/go-simplejson"
)

func main () {

//jsonファイルの読み込み
bytes, err := ioutil.ReadFile("./simple.json")
if err != nil {
  return
}

jsondata,err := simplejson.NewJson(bytes)
http.HandleFunc("/test", func (w http.ResponseWriter, r *http.Request) {

  res, err := json.Marshal(jsondata)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(res)

})

}
