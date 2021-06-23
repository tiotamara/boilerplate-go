package helpers

func InArrayString(needle string, hayStack []string) bool {
	for _, v := range hayStack {
		if v == needle {
			return true
		}
	}
	return false
}

func InArrayInt(needle int, hayStack []int) bool {
	for _, v := range hayStack {
		if v == needle {
			return true
		}
	}
	return false
}

func InArrayInt64(needle int64, hayStack []int64) bool {
	for _, v := range hayStack {
		if v == needle {
			return true
		}
	}
	return false
}
