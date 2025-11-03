package urlmasker

const Prefix = "http://"

func MaskURL(url string) string {
	buffer := []byte(url)
	prefix := []byte(Prefix)
	prefixLength := len(prefix)

	i := 0

	for i < len(buffer) {
		if containsPrefix(buffer, prefix, i) {
			i += prefixLength
			for i < len(buffer) && !isSpace(buffer[i]) {
				buffer[i] = byte('*')
				i++
			}
		} else {
			i++
		}
	}

	return string(buffer)
}

func containsPrefix(buffer []byte, prefix []byte, start int) bool {
	if start+len(prefix) > len(buffer) {
		return false
	}

	for i := 0; i < len(prefix); i++ {
		if buffer[start+i] != prefix[i] {
			return false
		}
	}

	return true
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}
