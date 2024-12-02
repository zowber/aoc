package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var file, err = os.Open("inputs")
	if err != nil {
		fmt.Println("No me gusta read file...", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		// do shit on each line omg i can't remember what little go i knew
	}

}
