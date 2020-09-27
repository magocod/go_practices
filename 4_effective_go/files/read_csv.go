package main

import (
    //"bufio"
    "encoding/csv"
    "fmt"
    // "io"
    "log"
    "os"
)

func main() {
    // Open the file
    csvfile, err := os.Open("model.csv")
    if err != nil {
        log.Fatalln("Couldn't open the csv file", err)
    }
    defer csvfile.Close()

    // Parse the file
    r := csv.NewReader(csvfile)
    r.Comma = ';'
    // r.FieldsPerRecord = 8 
    r.FieldsPerRecord = -8

    // Read each record from csv
    records, err := r.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    // var result []map[string]string
    
    header := records[0]
    fmt.Println("header", 0, header)

    // fmt.Print(records)
    for index, record := range records[1:] {
        fmt.Println(index, record)
        if len(record) == len(header) {
            row := make(map[string]string)
            for subi, subvalue := range record {
                fmt.Println("record: ", subi, subvalue, "prop: ", header[subi])
                // row[subi] = subvalue
                row[header[subi]] = subvalue
            }
            // fmt.Println(row)
            for key, mapvalue := range row {
                fmt.Println("map: ", key, mapvalue)
            }
        }
    }
        
}