package noregex

func Character() Pattern  { return pattern{"."} }
func Word() Pattern       { return pattern{`\w`} }
func Digit() Pattern      { return pattern{`\d`} }
func Letter() Pattern     { return pattern{"[a-zA-Z]"} }
func Lowercase() Pattern  { return pattern{"[a-z]"} }
func Uppercase() Pattern  { return pattern{"[A-Z]"} }
func Whitespace() Pattern { return pattern{`\s`} }
func Tab() Pattern        { return pattern{`\t`} }
func Newline() Pattern    { return pattern{`\n`} }
