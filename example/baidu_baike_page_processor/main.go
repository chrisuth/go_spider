//
package main

import (
    "fmt"
    "github.com/hu17889/go_spider/core/common/page"
    "github.com/hu17889/go_spider/core/spider"
    "strings"
)

type MyPageProcesser struct {
}

func NewMyPageProcesser() *MyPageProcesser {
    return &MyPageProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (this *MyPageProcesser) Process(p *page.Page) {
    if !p.IsSucc() {
        println(p.Errormsg())
        return
    }

    query := p.GetHtmlParser()

    name := query.Find(".lemmaTitleH1").Text()
    name = strings.Trim(name, " \t\n")

    summary := query.Find(".card-summary-content .para").Text()
    summary = strings.Trim(summary, " \t\n")

    // the entity we want to save by Pipeline
    p.AddField("name", name)
    p.AddField("summary", summary)
}

func main() {
    // spider input:
    //  PageProcesser ;
    //  task name used in Pipeline for record;
    sp := spider.NewSpider(NewMyPageProcesser(), "TaskName")
    pageItems := sp.Get("http://baike.baidu.com/view/1628025.htm?fromtitle=http&fromid=243074&type=syn", "html") // url, html is the responce type ("html" or "json" or "jsonp" or "text")

    url := pageItems.GetRequest().GetUrl()
    println("-----------------------------------spider.Get---------------------------------")
    println("url\t:\t" + url)
    for name, value := range pageItems.GetAll() {
        println(name + "\t:\t" + value)
    }

    println("\n--------------------------------spider.GetAll---------------------------------")
    urls := []string{
        "http://baike.baidu.com/view/1628025.htm?fromtitle=http&fromid=243074&type=syn",
        "http://baike.baidu.com/view/383720.htm?fromtitle=html&fromid=97049&type=syn",
    }
    pageItemsArr := sp.SetThreadnum(2).GetAll(urls, "html")
    for _, item := range pageItemsArr {
        url = item.GetRequest().GetUrl()
        println("url\t:\t" + url)
        fmt.Printf("item\t:\t%s\n", item.GetAll())
    }
}
