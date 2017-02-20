package main

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
        "os"
        "log"
        "os/exec"
)

func main() {
//      https://github.com/jacpy/GoDemo
//	url := "http://www.kancloud.cn/digest/hevc-fred/"
	url := "http://www.kancloud.cn/digest/hevc-fred/181912"
        p, err := goquery.NewDocument(url)
        if err != nil {
                panic(err)
        }
        fileInfoArr, err := ioutil.ReadDir("html")
        log.Println(fileInfoArr, err)
        if len(fileInfoArr) == 0 || err != nil && os.IsNotExist(err) {
                os.Mkdir("html", os.ModePerm)
                log.Println(err)
        }
        current, _ := os.Getwd()
        dir := current + "/html"
//	p.Find("div.tab-content").Find("ul.directory-list li").Each(func(idx int, selection *goquery.Selection) {
//                href, exists := selection.Find("a").Attr("href")
//		href = "http://www.kancloud.cn" + href
//                log.Println("exists: ", exists, "href: ", href, "idx: ", idx)
//        })
        p.Find("script").Remove()
	p.Find("div.manual-left").Remove()
        p.Find("div.think-loading").Remove()
	p.Find("div.article-head").Remove()
        p.Find("div.manual-head").Remove()
        p.Find("div.view-foot").Remove()
        content, err := p.Html()
	savefile := dir + "/output.html"
        err = ioutil.WriteFile(savefile, []byte(content), os.ModePerm)
        if err != nil {
                panic(err)
        }
        args := make([]string, 3)
        args[0] = "html/output.html"
        args[1] = "-o"
        args[2] = "output.md"
        cmd := exec.Command("pandoc", args...)
        err = cmd.Run()
        if err != nil {
                panic(err)
        }
	
}
