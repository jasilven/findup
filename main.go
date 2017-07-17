package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage:", os.Args[0], "dir1 dir2...")
		os.Exit(0)
	}

	results := make(map[string][]string)

	start := time.Now()
	for _, dir := range os.Args[1:] {
		err := filepath.Walk(dir, createWalker(&results))
		if err != nil {
			log.Fatalln(err)
		}
	}
	printResults(&results)

	fmt.Printf("Running time: %s\n", time.Since(start))
	fmt.Printf("Found %d files that have one or more duplicates\n", len(results))
}

func printResults(results *map[string][]string) {
	for sum, files := range *results {
		if len(files) > 1 {
			fmt.Println("#DUPLICATES:")
			for _, path := range files {
				fmt.Println(path)
			}
		} else {
			delete(*results, sum)
		}
	}
}

func createWalker(results *map[string][]string) filepath.WalkFunc {

	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() {
			dat := make([]byte, int(math.Min(1024, float64(info.Size()))))

			f, _ := os.Open(path)
			defer f.Close()

			_, e := io.ReadFull(f, dat)
			if e != nil {
				return e
			}
			dat = append(dat, []byte(strconv.FormatInt(info.Size(), 10))...)

			sum := fmt.Sprintf("%x", sha1.Sum(dat))
			(*results)[sum] = append((*results)[sum], path)
		}
		return nil
	}
}
