// Copyright 2020 Nigel Tao.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build ignore

package main

// test-go-strconv.go tests that the strconv package from the Go standard
// library correctly parses the generated data files in this repository.
//
// This seems somewhat circular at first glance, since those data files were
// generated by the extract-numbery-strings.go program, which also calls
// strconv.ParseFloat. Nonetheless, this program can still be useful when
// modifying the strconv.ParseFloat implementation, verifying that none of the
// strconv.ParseFloat results have changed for the hundreds of thousands of
// test cases in this repository's data directory.
//
// To test the current version of the Go standard library:
//   go run script/test-go-strconv.go data/*.txt
//
// To test version 1.4 of the Go standard library:
//   GOROOT=~/go1.4 ~/go1.4/bin/go run script/test-go-strconv.go data/*.txt
//
// Be aware that version 1.4 has a bug (https://golang.org/issue/15364 fixed in
// 2016) where strconv.ParseFloat("0e+308", 64) incorrectly returns +Infinity
// instead of 0. You will need to manually trim the data/*.txt files for this
// program to pass with Go 1.4.

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 0 {
		for _, arg := range os.Args[1:] {
			if err := do(arg); err != nil {
				fmt.Fprintf(os.Stderr, "file: %s\n%v\n", arg, err)
				os.Exit(1)
			}
		}
	}
}

func do(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	n, s := 0, bufio.NewScanner(f)
	for s.Scan() {
		line := s.Bytes()
		if len(line) < 18 {
			continue
		}

		want := uint64(0)
		for _, c := range line[:16] {
			want = want<<4 | uint64(hex[c])
		}
		src := string(line[17:])

		haveF64, err := strconv.ParseFloat(src, 64)
		if e, ok := err.(*strconv.NumError); ok && (e.Err == strconv.ErrRange) {
			// No-op. ParseFloat returns "Inf, ErrRange" for large inputs.
		} else if err != nil {
			return err
		}
		have := math.Float64bits(haveF64)

		if have != want {
			return fmt.Errorf("src:  %s\nhave: %016X\nwant: %016X", src, have, want)
		}
		n++
	}
	if err := s.Err(); err != nil {
		return err
	}
	fmt.Printf("%8d OK in %s\n", n, filename)
	return nil
}

var hex = [256]uint8{
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x00-0x07
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x08-0x0F
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x10-0x17
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x18-0x1F
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x20-0x27
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x28-0x2F
	0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, // 0x30-0x37 0-7
	0x8, 0x9, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x38-0x3F 8-9

	0x0, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF, 0x0, // 0x40-0x47 A-F
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x48-0x4F
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x50-0x57
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x58-0x5F
	0x0, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF, 0x0, // 0x60-0x67 a-f
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x68-0x6F
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x70-0x77
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x78-0x7F

	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x80-0x87
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x88-0x8F
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x90-0x97
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0x98-0x9F
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0xA0-0xA7
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0xA8-0xAF
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0xB0-0xB7
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0xB8-0xBF

	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0xC0-0xC7
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0xC8-0xCF
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0xD0-0xD7
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0xD8-0xDF
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0xE0-0xE7
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0xE8-0xEF
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0xF0-0xF7
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 0xF8-0xFF
}
