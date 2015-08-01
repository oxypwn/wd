package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFromUrl(url, file string) {
	fileName := file
	fmt.Println("Downloading", url, "to", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage : %s url location \n ", os.Args[0]) // return the program name back to %s
		os.Exit(1)                                            // graceful exit
	}
	url := os.Args[1]
	file := os.Args[2]
	downloadFromUrl(url, file)
}
