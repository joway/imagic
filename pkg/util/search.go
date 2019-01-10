package util

// IndexOf get the index of elem in target
func IndexOf(target string, elem string, reverse bool) int {
	length := len(target)
	for index := range target {
		var pos = index
		var val uint8
		if reverse {
			pos = length - index - 1
			val = target[pos]
		} else {
			val = target[index]
		}
		if Equal(string(val), elem) {
			return pos
		}
	}
	return -1
}
