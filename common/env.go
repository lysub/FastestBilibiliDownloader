package common

var CD int

// B站监测封禁,需更换进行伪装 log: wrong status code: 412
var UserAgent string
var UserAgentIndex int

var DefaultUserAgent = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:69.0) Gecko/20100101 Firefox/69.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0",
}

func Init() {
	if UserAgent == "" {
		UserAgent = DefaultUserAgent[UserAgentIndex]
	}
}
