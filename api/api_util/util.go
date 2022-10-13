package apiutil

import "log"

func FailOnError(err error, msg string) {
	if err != nil {
		log.Printf("%v ---> %s", err, msg)
	}
}
