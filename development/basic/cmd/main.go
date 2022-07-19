package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

const NTPhost = "3.ru.pool.ntp.org"

func main() {
	t, err := ntp.Time(NTPhost)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	fmt.Println(t)
}
