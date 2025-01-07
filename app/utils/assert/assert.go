package assert

func HasError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func IsTrue() bool {
	return true
}

func IsEmpty() bool {
	return false
}
