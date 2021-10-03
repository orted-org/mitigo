package main

import (
	"fmt"
	"time"
)

var NepaliDateMap map[int][12]int = map[int][12]int{
	2000: {30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2001: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2002: {31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2003: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2004: {30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2005: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2006: {31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2007: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2008: {31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31},
	2009: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2010: {31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2011: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2012: {31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2013: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2014: {31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2015: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2016: {31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2017: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2018: {31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2019: {31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2020: {31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2021: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2022: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	2023: {31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2024: {31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2025: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2026: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2027: {30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2028: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2029: {31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30},
	2030: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2031: {30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2032: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2033: {31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2034: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2035: {30, 32, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31},
	2036: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2037: {31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2038: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2039: {31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2040: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2041: {31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2042: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2043: {31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2044: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2045: {31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2046: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2047: {31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2048: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2049: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	2050: {31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2051: {31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2052: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2053: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	2054: {31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2055: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2056: {31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30},
	2057: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2058: {30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2059: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2060: {31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2061: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2062: {30, 32, 31, 32, 31, 31, 29, 30, 29, 30, 29, 31},
	2063: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2064: {31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2065: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2066: {31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31},
	2067: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2068: {31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2069: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2070: {31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2071: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2072: {31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2073: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2074: {31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2075: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2076: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	2077: {31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2078: {31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2079: {31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2080: {31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	2081: {31, 31, 32, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2082: {30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2083: {31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30},
	2084: {31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30},
	2085: {31, 32, 31, 32, 30, 31, 30, 30, 29, 30, 30, 30},
	2086: {30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2087: {31, 31, 32, 31, 31, 31, 30, 30, 29, 30, 30, 30},
	2088: {30, 31, 32, 32, 30, 31, 30, 30, 29, 30, 30, 30},
	2089: {30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2090: {30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
}

var NepaliDigit map[int]rune = map[int]rune{
	0: '०',
	1: '१',
	2: '२',
	3: '३',
	4: '४',
	5: '५',
	6: '६',
	7: '७',
	8: '८',
	9: '९',
}
var NepaliDays map[int][2]string = map[int][2]string{}

type GomitiDate struct {
	Year  int
	Month int
	Day   int
}

var firstNepaliDate = GomitiDate{
	Year:  2000,
	Month: 0,
	Day:   1,
}

// first english date
var firstEnglishDate time.Time = time.Date(1943, time.April, 14, 0, 0, 0, 0, time.UTC)

// number of total days in a year
var totalDaysInNepaliYear = make(map[int]int)

func init() {
	days := 0
	for i := 0; i < len(NepaliDateMap); i++ {
		for _, inv := range NepaliDateMap[firstNepaliDate.Year+i] {
			days += inv
		}
		totalDaysInNepaliYear[firstNepaliDate.Year+i] = days
	}
}
func main() {
	fmt.Println(NepaliDigit)
	date := time.Date(1985, time.November, 23, 0, 0, 0, 0, time.UTC)
	out, _ := ADtoBS(date)
	rev := BStoAD(out)
	fmt.Println(date)
	fmt.Println(rev)
}

func ADtoBS(ad time.Time) (GomitiDate, error) {
	ad = time.Date(ad.Year(), ad.Month(), ad.Day(), 0, 0, 0, 0, time.UTC)

	// difference between first english and requested english date
	daysDiff := int(ad.Sub(firstEnglishDate).Hours() / 24)
	for i := (daysDiff / 365) - 1; i <= (daysDiff/365)+1; i++ {
		if totalDaysInNepaliYear[firstNepaliDate.Year+i] > daysDiff {

			//nows days left to consider
			daysDiff -= totalDaysInNepaliYear[firstNepaliDate.Year+i-1]
			for j := 0; j < 12; j++ {
				daysDiff -= NepaliDateMap[firstNepaliDate.Year+i][j]
				if daysDiff <= 0 {
					return GomitiDate{firstNepaliDate.Year + i, j + 1, NepaliDateMap[firstNepaliDate.Year+i][j] + daysDiff + 1}, nil
				}
			}
			break
		}
	}
	return GomitiDate{}, fmt.Errorf("date out of supported range")
}
func BStoAD(bs GomitiDate) time.Time {

	tempEnglish := firstEnglishDate
	return tempEnglish.Add(time.Duration(time.Hour * 24 * time.Duration(DaysBetweenBSDates(firstNepaliDate, bs))))
}
func DaysBetweenBSDates(start, end GomitiDate) int {
	diffStart := start.Day
	diffEnd := end.Day
	diffStart += (totalDaysInNepaliYear[start.Year-1])
	diffEnd += (totalDaysInNepaliYear[end.Year-1])
	for i := 0; i < start.Month-1; i++ {
		diffStart += NepaliDateMap[start.Year][i]
	}
	for i := 0; i < end.Month-1; i++ {
		diffEnd += NepaliDateMap[end.Year][i]
	}

	return diffEnd - diffStart
}
