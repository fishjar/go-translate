package trans

// Def 翻译内容
type Def struct {
	Pos string `json:"pos"`
	Def string `json:"def"`
}

// DefList 翻译内容列表形式
type DefList struct {
	Pos string   `json:"pos"`
	Def []string `json:"def"`
}

// BilingualsDef 英汉双解内容
type BilingualsDef struct {
	Bil string `json:"bil"`
	Val string `json:"val"`
}

// Bilingual 英汉双解
type Bilingual struct {
	Pos string          `json:"pos"`
	Def []BilingualsDef `json:"def"`
}

// Sentence 例句
type Sentence struct {
	En string `json:"sen_en"`
	Cn string `json:"sen_cn"`
}

// DictRes 词典结果
type DictRes struct {
	Bot        string      `json:"bot"`        //翻译来源
	BotName    string      `json:"botName"`    //翻译来源名称
	PhoneticUS string      `json:"phoneticUS"` //音标
	PhoneticUK string      `json:"phoneticUK"` //音标
	AudioUS    string      `json:"audioUS"`    //读音
	AudioUK    string      `json:"audioUK"`    //读音
	Trans      []Def       `json:"trans"`      //翻译
	Variants   []Def       `json:"variants"`   //相关词
	Colls      []DefList   `json:"colls"`      //搭配
	Synonyms   []DefList   `json:"synonyms"`   //同义词
	Antonyms   []DefList   `json:"antonyms"`   //反义词
	Bilinguals []Bilingual `json:"bilinguals"` //英汉双解
	Ees        []DefList   `json:"ees"`        //英英
	Sentences  []Sentence  `json:"sentences"`  //例句
}

// GoogleRes 翻译结果
type GoogleRes struct {
	Q      string   `json:"q"`
	Sl     string   `json:"sl"`
	Tl     string   `json:"tl"`
	Trans  []string `json:"trans"`
	IsWord bool     `json:"isWord"`
}
