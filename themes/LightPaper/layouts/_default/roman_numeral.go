package main

import (
    "strconv"
    "strings"
)

// ToRomanNumeral converts an integer to a Roman numeral.
func ToRomanNumeral(year int) string {
    if year <= 0 || year > 3999 {
        return ""
    }

    romanNumerals := map[int]string{
        1000: "M",
        900:  "CM",
        500:  "D",
        400:  "CD",
        100:  "C",
        90:   "XC",
        50:   "L",
        40:   "XL",
        10:   "X",
        9:    "IX",
        5:    "V",
        4:    "IV",
        1:    "I",
    }

    var result strings.Builder

    for value, numeral := range romanNumerals {
        for year >= value {
            result.WriteString(numeral)
            year -= value
        }
    }

    return result.String()
}
