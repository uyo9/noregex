package noregex

func Character() Pattern     { return pattern{"."} }
func Word() Pattern          { return pattern{`\w`} }
func NotWord() Pattern       { return pattern{`\W`} }
func Digit() Pattern         { return pattern{`\d`} }
func NotDigit() Pattern      { return pattern{`\D`} }
func Letter() Pattern        { return pattern{"[a-zA-Z]"} }
func NotLetter() Pattern     { return pattern{`[^a-zA-Z]`} }
func Lowercase() Pattern     { return pattern{"[a-z]"} }
func NotLowercase() Pattern  { return pattern{`[^a-z]`} }
func Uppercase() Pattern     { return pattern{"[A-Z]"} }
func NotUppercase() Pattern  { return pattern{`[^A-Z]`} }
func Whitespace() Pattern    { return pattern{`\s`} }
func NotWhitespace() Pattern { return pattern{`\S`} }
func Tab() Pattern           { return pattern{`\t`} }
func Newline() Pattern       { return pattern{`\n`} }
