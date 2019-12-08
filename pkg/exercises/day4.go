package exercises

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func Day4a(r io.Reader) int {
	var c int
	input, _ := ioutil.ReadAll(r)
	rng := strings.Split(strings.TrimSpace(string(input)), "-")
	start, _ := strconv.Atoi(rng[0])
	end, _ := strconv.Atoi(rng[1])
	for i := start; i < end+1; i++ {
		if valid(strconv.Itoa(i), false) {
			c++
		}
	}
	return c
}

func Day4b(r io.Reader) int {
	var c int
	input, _ := ioutil.ReadAll(r)
	rng := strings.Split(strings.TrimSpace(string(input)), "-")
	start, _ := strconv.Atoi(rng[0])
	end, _ := strconv.Atoi(rng[1])
	for i := start; i < end+1; i++ {
		if valid(strconv.Itoa(i), true) {
			c++
		}
	}
	return c
}

func valid(password string, strict bool) bool {
	var b bool
	for i := 1; i < len(password); i++ {
		d := password[i]
		if d < password[i-1] {
			return false
		}
		if d == password[i-1] {
			if strict {
				if !strings.Contains(password, strings.Repeat(string(d), 3)) {
					b = true
				}
			} else {
				b = true
			}
		}
	}
	return b
}
