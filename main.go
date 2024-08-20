package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/zeebo/blake3"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
)

var (
	text       string
	file       string
	algorithms []string // Changed to slice of strings
	upperCase  bool
	outputJSON bool // New flag for JSON output
)

func init() {
	pflag.StringVarP(&text, "text", "t", "", "Input text")
	pflag.StringVarP(&file, "file", "f", "", "Input file")
	pflag.BoolVarP(&upperCase, "upper-case", "u", false, "Upper case output")
	pflag.StringSliceVarP(&algorithms, "algorithm", "a", []string{"sha256"}, "Hash algorithms: crc32, crc64, md5, sha1, sha224, sha256, sha384, sha512, blake2b, blake2s, blake3")
	pflag.BoolVarP(&outputJSON, "json", "j", false, "Output in JSON format")
	pflag.Parse()
}

func getHash(b io.Reader, algorithm string) string {
	var hash hash.Hash

	switch algorithm {
	case "crc32":
		hash = crc32.NewIEEE()
	case "crc64":
		hash = crc64.New(crc64.MakeTable(crc64.ECMA))
	case "md5":
		hash = md5.New()
	case "sha1":
		hash = sha1.New()
	case "sha224":
		hash = sha256.New224()
	case "sha256":
		hash = sha256.New()
	case "sha384":
		hash = sha512.New384()
	case "sha512":
		hash = sha512.New()
	case "blake2b":
		hash, _ = blake2b.New256(nil)
	case "blake2s":
		hash, _ = blake2s.New256(nil)
	case "blake3":
		hash = blake3.New()
	default:
		log.Fatalln("Invalid algorithm")
	}

	if _, err := io.Copy(hash, b); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func main() {
	var results = make(map[string]string)

	if file == "" && text == "" {
		log.Fatalln("Can't use empty string")
	}

	for _, algorithm := range algorithms {
		var reader io.Reader
		var err error

		if file != "" {
			var f *os.File
			f, err = os.Open(file)
			if err != nil {
				log.Fatalln(err)
			}
			defer f.Close()
			// Seek to the beginning of the file for each algorithm
			_, err = f.Seek(0, io.SeekStart)
			if err != nil {
				log.Fatalln("Error seeking file:", err)
			}
			reader = f
		} else {
			// Reinitialize the reader for each algorithm
			reader = strings.NewReader(text)
		}

		hashResult := getHash(reader, algorithm)
		if upperCase {
			hashResult = strings.ToUpper(hashResult)
		}
		results[algorithm] = hashResult
	}

	if outputJSON {
		jsonOutput, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			log.Fatalln("Error generating JSON output:", err)
		}
		fmt.Println(string(jsonOutput))
	} else {
		for algorithm, hash := range results {
			fmt.Printf("%s: %s\n", algorithm, hash)
		}
	}
}
