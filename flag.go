package noregex

type Flag interface {
	Token
	Raw() string
}

type flag struct {
	value string
	raw   string
}

func (f flag) Token() string { return f.value }
func (f flag) Raw() string   { return f.raw }

func CaseIgnored() Flag { return flag{"(?i)", "i"} }
func Multilined() Flag  { return flag{"(?m)", "m"} }
func Singlelined() Flag { return flag{"(?s)", "s"} }
