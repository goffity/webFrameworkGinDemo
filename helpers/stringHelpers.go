package helpers

func IsEmpty(data string) (ret bool) {
	if len(data) == 0 {
		ret = true
	} else {
		ret = false
	}

	return
}
