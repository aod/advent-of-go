package error

// Check : very basic error handling
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
