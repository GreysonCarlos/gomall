package utils

import "log"

func MustHandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}