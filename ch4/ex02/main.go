package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	usage := "Usage: echo -n hello | ./ex02 [256|384|512]"
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	n := 256
	if len(os.Args) == 2 {
		n, _ = strconv.Atoi(os.Args[1])
	}

	switch n {
	case 256:
		fmt.Printf("%x\n", sha256.Sum256(bytes))
		os.Exit(0)
		return
	case 384:
		fmt.Printf("%x\n", sha512.Sum384(bytes))
		os.Exit(0)
		return
	case 512:
		fmt.Printf("%x\n", sha512.Sum512(bytes))
		os.Exit(0)
		return
	default:
		fmt.Println(usage)
		os.Exit(1)
		return
	}
}
