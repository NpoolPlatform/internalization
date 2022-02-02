package constant

type Lang struct {
	Name string
}

var Langs = map[string]*Lang{
	"zh-CN": {
		Name: "Chinese Simplified",
	},
	"en-US": {
		Name: "English USA",
	},
}
