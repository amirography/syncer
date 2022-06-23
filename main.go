package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	for {
		syncRepo()
		log.Println("sleeping for 5 minutes")
		time.Sleep(time.Minute * 5)
	}
}
func syncRepo() {
	var home, _ = os.UserHomeDir()
	out, _ := exec.Command("just", home+"/notes/").CombinedOutput()
	fmt.Println(string(out))
}
