package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	//        "strconv"
)

var (
	process_switch bool
	filename       string = "data/hello.txt"
)

func main() {
	files, _ := ioutil.ReadDir("./data")
	aminos := []string{"ASP", "GLU", "ASN", "SER", "GLN", "HIS", "GLY", "THR", "CIT", "ARG", "b-ALA", "TAU", "ALA", "TYR", "TRP", "MET", "VAL", "PHE", "ILE", "LEU", "ORN", "LYS"}
	fmt.Print(" ,")
	fmt.Println(strings.Join(aminos, ","))

	for _, f := range files {
		sample_name, sample_data := get_data_from_file("./data/" + f.Name())
		fmt.Printf("%s,", sample_name)
		for key, amino := range aminos {
			fmt.Println(key)
			fmt.Println(amino)
			fmt.Println(sample_data[amino])
			//fmt.Println(sample_data)
		}
	}
}

func get_data_from_file(filename string) (sample_name string, sample_data map[string]string) {
	sample_data = make(map[string]string)
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
				sample_data[vec[1]] = vec[5]
			}
		}

	}
	return sample_name, sample_data

}
