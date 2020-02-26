package global

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type global struct {
	Path string
}

var Global = &global{}

func init() {
	var err error
	Global.Path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("init global path error:", err)
	}
	fmt.Println("init global path: ", Global.Path)
}
