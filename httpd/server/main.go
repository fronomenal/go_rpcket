package main

import "log"

func Start() error {
	return nil
}

func main() {
	if err := Start(); err != nil {
		log.Fatal(err)
	}
}
