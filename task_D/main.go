package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkLogin(login string) bool {
	lenLogin := len(login)

	if login[0] == '-' || lenLogin < 2 || lenLogin > 24 {
		return false
	} else {
		for i := 0; i < lenLogin; i++ {
			if (login[i] < 'a' || login[i] > 'z') && (login[i] < '0' || login[i] > '9') && login[i] != '-' && login[i] != '_' {
				return false
			}
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		logins := make(map[string]bool)
		var ats int
		fmt.Fscan(in, &ats)
		for j := 0; j < ats; j++ {
			var nextLogin string
			fmt.Fscan(in, &nextLogin)
			nextLoginLower := strings.ToLower(nextLogin)
			if checkLogin(nextLoginLower) {
				if _, ok := logins[nextLoginLower]; ok {
					fmt.Fprintln(out, "NO")
				} else {
					logins[nextLoginLower] = true
					fmt.Fprintln(out, "YES")
				}
			} else {
				fmt.Fprintln(out, "NO")
			}
		}
		fmt.Fprint(out, "\n")
	}
}
