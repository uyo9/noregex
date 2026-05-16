package noregex

func LineStart() Pattern       { return pattern{"^"} }
func LineEnd() Pattern         { return pattern{"$"} }
func WordBoundary() Pattern    { return pattern{`\b`} }
func NotWordBoundary() Pattern { return pattern{`\B`} }
