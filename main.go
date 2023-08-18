package main

import (
	"log"
	"os"

	"github.com/lvelvis/feishu/cmd/feishu"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	feishu.Execute()
}
