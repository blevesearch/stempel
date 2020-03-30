//  Copyright (c) 2018 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/blevesearch/stempel"
	"golang.org/x/text/encoding/charmap"
)

var input = flag.String("i", "", "input file")
var encoding = flag.String("e", "", "encoding of input")
var output = flag.String("o", "", "output file")
var trie = flag.String("t", "", "sample.tbl")

func main() {

	flag.Parse()

	trie, err := stempel.Open(*trie)
	if err != nil {
		log.Fatal(err)
	}

	var reader io.Reader = os.Stdin
	if *input != "" {
		var file *os.File
		file, err = os.Open(*input)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			cerr := file.Close()
			if cerr != nil {
				log.Fatal(cerr)
			}
		}()
		reader = file
	}

	if *encoding != "" {
		reader, err = findEncoding(*encoding, reader)
		if err != nil {
			log.Fatal(err)
		}
	}

	var writer = os.Stdout
	if *output != "" {
		var err error
		writer, err = os.Create(*output)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			cerr := writer.Close()
			if cerr != nil {
				log.Fatal(cerr)
			}
		}()
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		before := scanner.Text()
		hasSlash := strings.Index(before, "/")
		if hasSlash > 0 {
			before = before[0:hasSlash]
		}
		buff := []rune(before)
		diff := trie.GetLastOnPath(buff)
		buff = stempel.Diff(buff, diff)
		fmt.Fprintf(writer, string(buff))
		fmt.Fprintln(writer)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findEncoding(encoding string, r io.Reader) (io.Reader, error) {
	// walk through all charmaps looking for a match
	for _, enc := range charmap.All {
		if cm, ok := enc.(*charmap.Charmap); ok {
			if strings.Map(mapForCompare, cm.String()) == strings.Map(mapForCompare, encoding) {
				return cm.NewDecoder().Reader(r), nil
			}
		}
	}
	return nil, fmt.Errorf("no charmap found for encoding %s", encoding)
}

func mapForCompare(r rune) rune {
	// remove space and punctuation
	if unicode.IsSpace(r) {
		return -1
	} else if unicode.IsPunct(r) {
		return -1
	}
	// otherwise return lowercase
	return unicode.ToLower(r)
}
