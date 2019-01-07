package util

func Insert(target []interface{}, elem interface{}, pos int) []interface{} {
	return append(target[:pos], elem, target[pos:])
}
