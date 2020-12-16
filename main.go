package main

import (
	"flag"
	"fmt"
	"gotrans/trans"
	"strings"
)

func main() {
	const moreTip = "显示更详细的内容(仅英文单词有效)"
	var showMore bool
	flag.BoolVar(&showMore, "more", false, moreTip)
	flag.BoolVar(&showMore, "m", false, moreTip)
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("没有翻译内容")
	}
	word := strings.Join(args, " ")

	// 谷歌翻译
	googleRes, err := trans.FetchGoogle(word)
	if err != nil {
		fmt.Println("获取谷歌数据错误", err)
		return
	}
	// trans.PrintJSON(googleRes)
	trans.PrintGoogleTrans(googleRes)
	if !googleRes.IsWord {
		return
	}

	// bing词典
	bingRes, err := trans.FetchBing(word)
	if err != nil {
		fmt.Println("获取BING数据错误", err)
		return
	}
	// trans.PrintJSON(bingRes)
	trans.PrintDict(bingRes, showMore)
}
