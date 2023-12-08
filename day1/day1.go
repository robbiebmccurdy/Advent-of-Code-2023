package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "unicode"
)

func main() {
    // Open the file (replace 'calibration.txt' with your file path)
    file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var sum int
    scanner := bufio.NewScanner(file)

    // Process each line of the file
    for scanner.Scan() {
        line := scanner.Text()
        firstDigit, lastDigit := -1, -1

        // Find the first and last digit
        for _, r := range line {
            if unicode.IsDigit(r) {
                if firstDigit == -1 {
                    firstDigit = int(r - '0')
                }
                lastDigit = int(r - '0')
            }
        }

        // If both digits are found, add them to the sum
        if firstDigit != -1 && lastDigit != -1 {
            value, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
            sum += value
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    // Output the total sum
    fmt.Println("Total sum of calibration values:", sum)
}

