package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	//        "strconv"
)

var (
	data           []uint64
	process_switch bool
	sum            uint64 = 0
	filename       string = "data/hello.txt"
	m                     = make(map[string]string)
)

func main() {
	file, _ := os.Open(filename)
	defer file.Close()
	r := bufio.NewReader(file)
	for {

		line, err := r.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}

		if strings.Contains(line, "Sample Name") {
			sample_name := strings.Split(line, "\t")
			fmt.Println(sample_name[1])
		}

		if strings.Contains(line, "Amount") {
			process_switch = true
		}

		if strings.Contains(line, "Page") {
			process_switch = false
		}

		if process_switch {
			vec := strings.Split(line, "\t")
			if len(vec) == 6 {
				m[vec[1]] = vec[5]
			}
		}

	}
	fmt.Println(m)

}
