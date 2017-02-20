package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func fetchHtml(url string, outputName string) {
	p, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	p.Find("script").Remove()
	p.Find("div.manual-left").Remove()
	p.Find("div.think-loading").Remove()
	p.Find("div.article-head").Remove()
	p.Find("div.manual-head").Remove()
	p.Find("div.view-foot").Remove()
	content, err := p.Html()
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(outputName, []byte(content), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func translateFromHtmlToMarkdown(input string, output string) {
	args := make([]string, 3)
	args[0] = input
	args[1] = "-o"
	args[2] = output
	cmd := exec.Command("pandoc", args...)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	//      https://github.com/jacpy/GoDemo
	url := os.Args[1]
	p, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	_, err = ioutil.ReadDir("html")
	if os.IsNotExist(err) {
		os.Mkdir("html", os.ModePerm)
	}
	_, err = ioutil.ReadDir("md")
	if os.IsNotExist(err) {
		os.Mkdir("md", os.ModePerm)
	}
	current, _ := os.Getwd()
	summary := "# 目录"
	summary = summary + "\n\n"
	p.Find("div.tab-content").Find("ul.directory-list li").Each(func(idx int, selection *goquery.Selection) {
		title := strings.TrimSpace(selection.Text())
		href, _ := selection.Find("a").Attr("href")
		href = "http://www.kancloud.cn" + href
		strIdx := fmt.Sprint(idx)
		chapHtml := current + "/html/chap" + strIdx + ".html"
		chapMD := current + "/md/chap" + strIdx + ".md"
		fetchHtml(href, chapHtml)
		translateFromHtmlToMarkdown(chapHtml, chapMD)
		summary = summary + "* [" + title + "](chap" + strIdx + ".md)\n"
	})
	err = ioutil.WriteFile(current+"/md/"+"SUMMARY.md", []byte(summary), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
