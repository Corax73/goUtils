package curl

type Request struct {
	Method, Url   string
	Headers, Data map[string]string
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
