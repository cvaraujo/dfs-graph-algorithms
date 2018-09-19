package Functions

import (
	"bufio"
	"log"
	"os"
)

func read_archive(name string) {
	file, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	
}
