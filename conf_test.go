package conf

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestInit(t *testing.T) {

	Init("demo")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dir)

	str, _ := os.Getwd()
	fmt.Println(str)

}
