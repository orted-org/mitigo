package main

import (
	"fmt"
	"time"
)

// source of truth for date conversion
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

// nepali digit
var NepaliDigit map[int]string = map[int]string{
	0: "०",
	1: "१",
	2: "२",
	3: "३",
	4: "४",
	5: "५",
	6: "६",
	7: "७",
	8: "८",
	9: "९",
}

// name of all the BS months
var NepaliMonthName map[int][2]string = map[int][2]string{
	1:  {"Baisakh", "वैशाख"},
	2:  {"Jeshta", "जेठ"},
	3:  {"Ashadh", "असार"},
	4:  {"Shrawan", "श्रावण"},
	5:  {"Bhadau", "भदौ"},
	6:  {"Ashwin", "असोज"},
	7:  {"Kartik", "कार्तिक"},
	8:  {"Mangsir", "मंसिर"},
	9:  {"Poush", "पौष"},
	10: {"Magh", "माघ"},
	11: {"Falgun", "फागुन"},
	12: {"Chaitra", "चैत"},
}

type MitigoDate struct {
	Year  int
	Month int
	Day   int
}

// first nepali date
var firstNepaliDate = MitigoDate{
	Year:  2000,
	Month: 0,
	Day:   1,
}

// first english date
var firstEnglishDate time.Time = time.Date(1943, time.April, 14, 0, 0, 0, 0, time.UTC)

// for number of cumulative total days year wise
var cumulativeDaysInNepaliYear = make(map[int]int)

// init function to calculate the cumulative days in year
func init() {
	days := 0
	for i := 0; i < len(NepaliDateMap); i++ {
		for _, inv := range NepaliDateMap[firstNepaliDate.Year+i] {
			days += inv
		}
		cumulativeDaysInNepaliYear[firstNepaliDate.Year+i] = days
	}
}

// function to convert AD to BS
func ADtoBS(ad time.Time) (MitigoDate, error) {
	ad = time.Date(ad.Year(), ad.Month(), ad.Day(), 0, 0, 0, 0, time.UTC)

	if ad.Before(firstEnglishDate) {
		return MitigoDate{}, fmt.Errorf("date out of supported range")
	}

	// difference between first english and requested english date
	daysDiff := int(ad.Sub(firstEnglishDate).Hours()/24) + 1
	for i := (daysDiff / 365) - 1; i <= (daysDiff/365)+1; i++ {
		if cumulativeDaysInNepaliYear[firstNepaliDate.Year+i] > daysDiff {
			//nows days left to consider
			daysDiff -= cumulativeDaysInNepaliYear[firstNepaliDate.Year+i-1]
			for j := 0; j < 12; j++ {
				daysDiff -= NepaliDateMap[firstNepaliDate.Year+i][j]
				if daysDiff <= 0 {
					return MitigoDate{firstNepaliDate.Year + i, j + 1, NepaliDateMap[firstNepaliDate.Year+i][j] + daysDiff}, nil
				}
			}
			break
		}
	}
	return MitigoDate{}, fmt.Errorf("date out of supported range")
}

// function to convert BS to AD
func BStoAD(bs MitigoDate) (time.Time, error) {
	if diffDays := DaysBetweenBSDates(firstNepaliDate, bs); diffDays < 0 {
		return firstEnglishDate, fmt.Errorf("date out of supported range")
	} else {
		tempEnglish := firstEnglishDate
		return tempEnglish.Add(time.Duration(time.Hour * 24 * time.Duration(diffDays))), nil
	}
}

// function to calculate number of days between two BS dates
func DaysBetweenBSDates(start, end MitigoDate) int {
	diffStart := start.Day
	diffEnd := end.Day
	diffStart += (cumulativeDaysInNepaliYear[start.Year-1])
	diffEnd += (cumulativeDaysInNepaliYear[end.Year-1])
	for i := 0; i < start.Month-1; i++ {
		diffStart += NepaliDateMap[start.Year][i]
	}
	for i := 0; i < end.Month-1; i++ {
		diffEnd += NepaliDateMap[end.Year][i]
	}

	return diffEnd - diffStart
}

// returns BS date in Nepali typing in format YYYYdelimMMdelimDay
func BSDateInNepali(bs MitigoDate, delim string) string {
	str := ""
	ToNepaliDigitString(bs.Year, &str)
	str += delim
	ToNepaliDigitString(bs.Month, &str)
	str += delim
	ToNepaliDigitString(bs.Day, &str)
	return str
}

// returns BS month name in Nepali or English typing
func BSMonthName(bs MitigoDate, inNepali bool) string {
	if inNepali {
		return NepaliMonthName[bs.Month][1]
	} else {
		return NepaliMonthName[bs.Month][0]
	}
}

// converts English digit to Nepali digit
func ToNepaliDigitString(digit int, str *string) {
	if digit == 0 {
		return
	}
	ToNepaliDigitString(digit/10, str)
	*str += NepaliDigit[digit%10]
}

// returns back date after nDays of days. nDays can be negative or positive according to need
func BSDaysAfter(bs MitigoDate, nDays int) (MitigoDate, error) {
	inAD, err := BStoAD(bs)
	if err != nil {
		return bs, err
	}
	inAD = inAD.Add(time.Duration(time.Hour * 24 * time.Duration(nDays)))
	return ADtoBS(inAD)
}
