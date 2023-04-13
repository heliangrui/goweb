package stringUtils

// EscapeStr 字符转义
func EscapeStr(param string) (result string) {
	if param == "" {
		return
	}
	for _, c := range param {
		if c == '\\' || c == '+' || c == '-' || c == '!' || c == '(' || c == ')' || c == ':' ||
			c == '_' || c == '^' || c == '[' || c == ']' || c == '{' || c == '}' || c == '~' ||
			c == '*' || c == '?' || c == '|' || c == '&' || c == ';' || c == '/' || c == '.' || c == '$' || c == '%' {
			result += "\\"
		}
		result += string(c)
	}
	return
}
