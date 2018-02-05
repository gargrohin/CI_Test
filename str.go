package aca

import (
	"strconv"
	"strings"
	"unicode"
)

func check_ip(ip string) string {

	var full []string = strings.Split(ip, ".")

	if len(full) != 4 {
		return "invalid"
	}

	//var part int
	for _, i := range full {
		part, err := strconv.Atoi(i)
		if err != nil {
			return "invalid"
		} else if part <= 0 || part >= 255 {
			return "invalid"
		}
	}

	return "valid"

}

func rm_alnum(raw string) string {
	var str string = ""
	for _, r := range raw {
		if unicode.IsLetter(r) {
			str = str + string(r)
			continue
		}

		_, err := strconv.Atoi(string(r))
		if err != nil {
			str = str + string(r)
		}

	}

	return str
}
