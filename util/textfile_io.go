package util

import "io/ioutil"

/*
	Takes a path to a text file as an argument and returns the file's contents
	as a single string. If no file could be found under the given path, an error is thrown.
 */
func ContentAsString(filePath string) string {
	data, err := ioutil.ReadFile(filePath)

	if (err != nil) {
		panic(err)
	}

	return string(data)
}