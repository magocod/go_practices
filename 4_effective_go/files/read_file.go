package main

import (
	"fmt"
	"log"
	"os"
    "io"
    "reflect"
    "strings"
)

// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}

func main() {

	fileContent, err := Contents("model.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
    // fmt.Println(fileContent)
    // fmt.Println(reflect.TypeOf(fileContent))

    ContentSlice := strings.Split(fileContent, "\n")
    // fmt.Println(ContentSlice)
    // fmt.Println(reflect.TypeOf(ContentSlice))

    header := strings.Split(ContentSlice[0], ";")
    fmt.Println(header)
    for i, v := range header {
        fmt.Println("header: ", i, v)
    }
    fmt.Println(reflect.TypeOf(header))

    for index, value := range ContentSlice[1:] {
        fmt.Println("value: ", index, value)
        values := strings.Split(value, ";")

        if len(values) == len(header) {
            row := make(map[int]string)
            for subi, subvalue := range values {
                fmt.Println("values: ", subi, subvalue)
                row[subi] = subvalue
            }
            // fmt.Println(row)
            for key, mapvalue := range row {
                fmt.Println("map: ", key, mapvalue)
            }
        }

    }

}