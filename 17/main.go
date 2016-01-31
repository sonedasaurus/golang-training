package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	err := fetch(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func fetch(args []string) error {
	for _, url := range args {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
