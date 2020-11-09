package tokens

// Type is a token type.
type Type string

const (
	ECHO = "ECHO"
)

type token struct{
	Type Type
	Literal string
}

var keywords = map[string]Type{
	"echo": ECHO
}
