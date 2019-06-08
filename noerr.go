package noerr

import "log"

//Log Returns wether or not the error is nil, logs the error if err != nil
func Log(err error) bool {
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

//Fatal Checks if the error is nil, if not log.Fatal(err)
func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Panic Checks if the error is nil, if not panic(err)
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
