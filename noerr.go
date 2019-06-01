package noerr

import "log"

//NoErr Returns wether or not the error is nil, logs the error if err != nil
func NoErr(err error) bool {
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

//NoErrF Checks if the error is nil, if not log.Fatal(err)
func NoErrF(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//NoErrP Checks if the error is nil, if not panic(err)
func NoErrP(err error) {
	if err != nil {
		panic(err)
	}
}
