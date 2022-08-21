package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	flag.Parse()
	var dePoch string

	if flag.NArg() == 0 { // from stdin/pipe
		reader := bufio.NewReader(os.Stdin)
		var err error
		dePoch, err = reader.ReadString('\n')
		if err != nil {
			log.Fatalln("failed to read input")
		}
		dePoch = strings.TrimSpace(dePoch) // otherwise, we would have a blank line
	} else { // from argument
		if flag.NArg() > 1 {
			log.Fatalln("takes at most one input")
		}
		dePoch = os.Args[1]
	}

	i, err := strconv.ParseInt(dePoch, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	fmt.Println(tm)
}
