package noerr

import "log"

//Test Returns err == nil
func Test(err error) bool {
	return err == nil
}

//Log Returns wether or not the error is nil, logs the error if err != nil
func Log(err error) bool {
	t := !Test(err)
	if t {
		log.Println(err)
	}

	return t
}

//Fatal Checks if the error is nil, if not log.Fatal(err)
func Fatal(err error) {
	if !Test(err) {
		log.Fatal(err)
	}
}

//Panic Checks if the error is nil, if not panic(err)
func Panic(err error) {
	if !Test(err) {
		panic(err)
	}
}
