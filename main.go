package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

const location = "Asia/Tokyo"

func main() {
	flag.Parse()
	args := flag.Args()

	t, err := getTime(args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t.UnixNano() / 1000000)
}

func getTime(args []string) (time.Time, error) {
	if len(args) <= 0 {
		return time.Now(), nil
	}

	format := "20060102150405"
	jst := time.FixedZone(location, 9*60*60)
	t, err := time.ParseInLocation(format, args[0], jst)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
