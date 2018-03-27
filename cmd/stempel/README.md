# stempel command-line tool

A simple command-line tool which can read text from standard-input, apply the stemming rules using the specified table, and write to standard-output the stemmed text.

## Example

In this example, we specify:

- the stemming table `stemmer_200000.tbl`
- the input text from a free polish dictionary
- the input encoding iso-8859 (output is always utf-8)
- the output is redirected to a file

```
$ stempel -e iso-8859-2 -t ../../stemmer_200000.tbl < /Users/mschoch/Downloads/pl_PL.dic > pl_stemmed.txt
```
