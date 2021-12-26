package constant

type Lang struct {
	Name string
}

var Langs = map[string]*Lang{
	"zh_CN": {
		Name: "Chinese Simplified",
	},
	"en_US": {
		Name: "English USA",
	},
}
