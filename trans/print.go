package trans

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

const templDict = `
{{range .Variants}}{{.Pos | printf "[%s]" | printf "%-20s"}} {{.Def}}
{{end}}
美	{{.PhoneticUS}}
英	{{.PhoneticUK}}

{{range .Trans}}{{.Pos | printf "[%s]" | printf "%-7s"}} {{.Def}}
{{end}}
`

const templDictMore = templDict + `
	---------- 搭配 ----------
{{range .Colls}}{{.Pos | printf "[%s]" | printf "%-7s"}} {{StringsJoin .Def ", "}}
{{end}}
	---------- 同义词 ----------
{{range .Synonyms}}{{.Pos | printf "[%s]" | printf "%-7s"}} {{StringsJoin .Def ", "}}
{{end}}
	---------- 反义词 ----------
{{range .Antonyms}}{{.Pos | printf "[%s]" | printf "%-7s"}} {{StringsJoin .Def ", "}}
{{end}}
	---------- 英汉双解 ----------
{{range .Bilinguals}}{{.Pos | printf "[%s]"}}
{{range .Def}}	{{.Val}}
	{{.Bil}}
{{end}}{{end}}
	---------- 英英 ----------
{{range .Ees}}{{.Pos | printf "[%s]"}}
{{range .Def}}	{{.}}
{{end}}{{end}}
	---------- 例句 ----------
{{range .Sentences}}
	{{.En}}
	{{.Cn}}
{{end}}
`

var reportDict = template.Must(template.New("dict").
	Funcs(template.FuncMap{"StringsJoin": strings.Join}).
	Parse(templDict))

var reportDictMore = template.Must(template.New("dict").
	Funcs(template.FuncMap{"StringsJoin": strings.Join}).
	Parse(templDictMore))

// PrintGoogleTrans 打印谷歌翻译结果
func PrintGoogleTrans(res GoogleRes) {
	// fmt.Println("[谷歌翻译]")
	for _, tr := range res.Trans {
		fmt.Printf("\n%s\n", tr)
	}
}

// PrintDict 打印词典结果
func PrintDict(res DictRes, showMore bool) {
	// fmt.Println("[" + res.BotName + "]")
	templ := reportDict
	if showMore {
		templ = reportDictMore
	}
	if err := templ.Execute(os.Stdout, res); err != nil {
		log.Fatal(err)
	}
}

// PrintJSON 转为json，并打印
func PrintJSON(data interface{}) {
	json, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	fmt.Printf("%s\n", json)
}
