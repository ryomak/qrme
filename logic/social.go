package logic

type Social struct {
	Name    string
	Account string
	URL     string
}

var socialMap = map[string]string{
	"twitter":   "https://twitter.com",
	"instagram": "https://instagram.com",
}
