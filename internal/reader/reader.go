package reader

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ReadFile(name string) {
	f, err = os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	readToConsoleLike()

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}

func readToConsoleLike() {
	bdata := make([]byte, 1024)
	for {
		n, err := f.Read(bdata)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		if n > 0 {
			fmt.Println(string(bdata[:n]))
		}
	}
}

func ReadToConsole(name string) {
	data, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(data))
}
