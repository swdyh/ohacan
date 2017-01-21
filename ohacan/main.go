package main

import (
	"fmt"
	"github.com/swdyh/ohacan"
	"os"
)

func main() {
	groupID := os.Getenv("JOBCAN_GROUP_ID")
	code := os.Getenv("JOBCAN_CODE")
	devServer := os.Getenv("OHACAN_DEV_SERVER")

	if groupID == "" || code == "" {
		fmt.Println("Error: please set JOBCAN_GROUP_ID/JOBCAN_CODE")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage:\n ohacan oha\n ohacan otsu")
		os.Exit(1)
	}

	c := ohacan.Client{GroupID: groupID, Code: code}
	if devServer != "" {
		c.URL = devServer
		fmt.Println("DEBUG: use", c.URL)
	}

	if strContains([]string{"oha", "clockin", "clock-in", "in"}, os.Args[1]) {
		_, err := c.ClockIn()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("\U0001f646") // ok_woman
		}
	}

	if strContains([]string{"otsu", "otu", "clockout", "clock-out", "out"}, os.Args[1]) {
		_, err := c.ClockOut()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("\U0001f37a") // beer
		}
	}
}

func strContains(ss []string, s string) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}
	return false
}
