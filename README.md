# stempel

[![Build Status](https://travis-ci.org/blevesearch/stempel.svg?branch=master)](https://travis-ci.org/blevesearch/stempel)
[![Coverage Status](https://coveralls.io/repos/github/blevesearch/stempel/badge.svg?branch=master)](https://coveralls.io/github/blevesearch/stempel?branch=master)
[![GoDoc](https://godoc.org/github.com/blevesearch/stempel?status.svg)](https://godoc.org/github.com/blevesearch/stempel)
[![Go Report Card](https://goreportcard.com/badge/github.com/blevesearch/stempel)](https://goreportcard.com/report/github.com/blevesearch/stempel)

A Go implementation of the [Stempel](http://www.getopt.org/stempel/) stemmer, an algorithmic stemmer frequently used with (but not limited to) the Polish language.

## License

This Go implementation is distributed under the Apache 2 license.

This product is based on software developed by the [Egothor Project](http://egothor.sf.net/).

This project now also includes the original stemming tables pretrained by Andrzej Białecki, and [as noted by Lucene](https://github.com/apache/lucene-solr/blob/f41eabdc5fa091079b83cdc7813cdcfb05dfbf46/lucene/analysis/stempel/src/java/overview.html#L46-L47) the tables and other additions are covered by Apache License 2.0.