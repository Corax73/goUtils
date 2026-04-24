package curl

type Request struct {
	Method, Url, Data string
	Headers           map[string]string
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
