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

package stempel

import (
	"os"
	"strings"

	"github.com/mschoch/stempel/javadata"
)

// Trie is the external interface to work with the stempel trie
type Trie interface {
	GetLastOnPath([]rune) []rune
}

// Open attempts to open a file at the specified path, and use it to
// build a Trie
func Open(path string) (Trie, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	r := javadata.NewReader(f)
	method, err := r.ReadUTF()
	if err != nil {
		return nil, err
	}

	var rv Trie
	if strings.Contains(method, "M") {
		rv, err = newMultiTrie(r)
		if err != nil {
			return nil, err
		}
	} else {
		rv, err = newTrie(r)
		if err != nil {
			return nil, err
		}
	}
	return rv, nil
}
