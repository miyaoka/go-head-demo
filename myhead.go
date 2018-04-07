package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

func scanFile(filePath string, maxLines uint, showLineNum bool) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	s := bufio.NewScanner(file)

	if showLineNum {
		digits := strconv.FormatFloat(math.Floor(math.Log10(float64(maxLines)))+1, 'g', 4, 64)
		fmtStr := "%" + digits + "d  %s\n"
		for i := uint(0); s.Scan() && i < maxLines; i++ {
			fmt.Printf(fmtStr, i+1, s.Text())
		}
	} else {
		for i := uint(0); s.Scan() && i < maxLines; i++ {
			fmt.Println(s.Text())
		}
	}
}

func main() {
	maxLines := flag.Uint("n", 10, "lines")
	showLineNum := flag.Bool("l", false, "line number")
	flag.Parse()
	filePathes := flag.Args()
	fileCounts := len(filePathes)

	if fileCounts == 1 {
		scanFile(filePathes[0], *maxLines, *showLineNum)
	} else {
		lastIdx := fileCounts - 1
		for i, filePath := range filePathes {
			fmt.Println("==> ", filePath, " <==")
			scanFile(filePath, *maxLines, *showLineNum)
			if i != lastIdx {
				fmt.Println("") // Blank line as file seperator
			}
		}
	}
}
