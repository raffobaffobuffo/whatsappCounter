package main

import (
    "bufio"
    "bytes"
    "fmt"
    "io"
    "os"
    "strings"
)

func findName(line string) (err error) {
	if !strings.Contains(line, "]") {
		return
	}
	start := strings.IndexByte(line, ']')
	if !strings.Contains(line[start:], ":") {
		return
	}
	end := strings.IndexByte(line[start:], ':')
	if end < 0 && end+start < start {
		return
	}
	name := line[start+2:end+start]
	if _, ok := db[name]; ok {
		db[name] += 1
	} else {
		db[name] = 1
	}
	return
	}

func readByLine(fn string) (err error) {

    file, err := os.Open(fn)
    if err != nil {
        return err
    }
    defer file.Close()

    // Start reading from the file with a reader.
    reader := bufio.NewReader(file)
    for {
        var buffer bytes.Buffer

        var l []byte
        var isPrefix bool
        for {
            l, isPrefix, err = reader.ReadLine()
            buffer.Write(l)
            // If we've reached the end of the line, stop reading.
            if !isPrefix {
                break
            }
            // If we're at the EOF, break.
            if err != nil {
                if err != io.EOF {
                    return err
                }
                break
            }
        }
        line := buffer.String()

        // Process the line here.
	newLine := limitLength(line, 40)
	findName(newLine)
	//fmt.Printf("%s\n", newLine)

        if err == io.EOF {
            break
        }
    }
    if err != io.EOF {
        fmt.Printf(" > Failed with error: %v\n", err)
        return err
    }
    return
}

func limitLength(s string, length int) string {
    if len(s) < length {
        return s
    }
    return s[:length]
}

var db = make(map[string]int)

func main() {
	readByLine(os.Args[1])
	fmt.Println(db)
	return
}
