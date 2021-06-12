package main

import (
	"fmt"
	"os"
)

func main() {

	dirPath := [4]string{"/Users/den/Desktop/Git/integration-tests/reports/",
		"/Users/den/Desktop/Git/integration-tests/tests/ui/_out/allure_reports",
		"/Users/den/Desktop/Git/integration-tests/tests/_out/allure_reports",
		"/Users/den/Desktop/Git/integration-tests/tests/backend/_out/allure_reports"}

	for _, value := range dirPath {
		err := os.RemoveAll(value)
		if err != nil {
			return
		}
		err2 := os.MkdirAll(value, os.FileMode(0777))
		if err2 != nil {
			return
		}
		fmt.Println("Успех!")
	}
}
