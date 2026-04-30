package curl

type Request struct {
	Method, Url, Data string
	Headers, Cookies  map[string]string
}

var validKeys = []string{
	"-G",
	"--get",
	"-I",
	"--head",
	"-X",
	"--request",
	"-H",
	"--header",
	"-d",
	"--data",
	"-b",
	"--cookie",
}

var requestMethods = []string{
	"-G",
	"--get",
	"-I",
	"--head",
}

var headerKeys = []string{
	"-H",
	"--header",
}

var dataKeys = []string{
	"-d",
	"--data",
}

var cookieKeys = []string{
	"-b",
	"--cookie",
}
