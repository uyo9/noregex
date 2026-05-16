package noregex

func LineStart() Pattern       { return pattern{"^"} }
func LineEnd() Pattern         { return pattern{"$"} }
func TextStart() Pattern       { return pattern{`\A`} }
func TextEnd() Pattern         { return pattern{`\z`} }
func WordBoundary() Pattern    { return pattern{`\b`} }
func NotWordBoundary() Pattern { return pattern{`\B`} }
