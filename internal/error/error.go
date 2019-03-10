package error

// Check is a very simple error handler
// which uses the built-in panic function
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
