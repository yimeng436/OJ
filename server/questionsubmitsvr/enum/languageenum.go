package enum

import "errors"

const (
	Java   = "java"
	CPP    = "CPP"
	Golang = "Golang"
)

var languageMap = make(map[string]string)

func init() {
	languageMap[Java] = "java"
	languageMap[CPP] = "CPP"
	languageMap[Golang] = "Golang"
}

func GetLanguage(language string) (string, error) {
	res, err := languageMap[language]
	if err {
		return "不支持该语言", errors.New("不支持该语言")
	}
	return res, nil
}
