package commonutil

import "log"

func IndexOf(arr []string, e string) int {
	for k, v := range arr {
		if e == v {
			return k
		}
	}
	return -1 //not found.
}

func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
