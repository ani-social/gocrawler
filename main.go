package main

import (
  "fmt"
  "github.com/gocolly/colly"
  "github.com/fatih/color"
)

var (
  root = "https://gogoanime.llc"
)

func main() {
  c := colly.NewCollector()

  c.OnRequest(func(r *colly.Request) {
    color.Green("[gocrawler/info]: Fetching %s", r.URL)
  })

  for i := 0; i <= 91; i++ {
    url := fmt.Sprintf("%s/anime-list.html?page=%d", root, i)
    if err := c.Visit(url); err != nil {
      color.Red("[gocrawler/error]: Error visiting %s", root)
    }
  }
}
