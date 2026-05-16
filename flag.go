package noregex

type Flag interface{ Token }

type flag struct{ value string }

func (f flag) Token() string { return f.value }

func CaseIgnored() Flag { return flag{"(?i)"} }
func Multilined() Flag  { return flag{"(?m)"} }
func Singlelined() Flag { return flag{"(?s)"} }
