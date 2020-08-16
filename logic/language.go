package logic

type Language struct {
	Name string
	Icon string
}

var languageMap = map[string]Language{
	"go": {
		Name: "Go",
		Icon: "localhost",
	},
	"none": {
		Name: "none",
		Icon: "localhost",
	},
}
