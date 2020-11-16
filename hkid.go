/*
Copyright 2020 Kenth Fagerlund

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// heimdallr package description
package heimdallr

import (
	"regexp"
)

const (
	// VALID_FORMAT describes a normalized HKID. Example: R5533446
	VALID_FORMAT       = `^[A-NP-Z]{1,2}[0-9]{6}[0-9A]$`
	NORMALIZE          = `[\(\)]`
	CHARACTER_OFFSET   = 55
	PREFIX_SPACE_VALUE = 324
)

// validateCheckNumber will calculate the checksum and validate it to the given code
//
// HK digit pattern
//
// XX999999C / X999999C
func validateCheckNumber(normalizedID string) bool {
	sc := []rune(normalizedID)
	var sum int
	// If we only have 8 character then add
	if len(sc) == 8 {
		sum = PREFIX_SPACE_VALUE
	}
	for i, v := range sc {
		if i == (len(sc) - 1) { //We exit early so we don't add check digit into our calculation
			break
		}
		weight := len(sc) - i
		sum += (runeToInt(v) * weight)
		//mt.Printf("Loop for %s. R=%#U,%d I=%d, weight=%d, sum=%d, append=%d  ==||==  ", normalizedID, v, runeToInt(v), i, weight, sum, (runeToInt(v) * weight))

	}
	// sum := 36*9 + ((int(sc[0])-CHARACTER_OFFSET)*8) + ((int(sc[1])-CHARACTER_OFFSET)*7) + ((int(sc[2])-CHARACTER_OFFSET)*6) + (int(sc[3])-CHARACTER_OFFSET)*5 + (int(sc[4])-CHARACTER_OFFSET)*4 + (int(sc[5])-CHARACTER_OFFSET)*3 + (int(sc[6])-CHARACTER_OFFSET)*2
	checkDigit := 11 - (sum % 11)
	idcheck := runeToInt(sc[len(sc)-1])

	if checkDigit != idcheck {
		return false
	} else {
		return true
	}
}

func runeToInt(r rune) int {
	var intvalue int
	if int(r) > 64 {
		intvalue = int(r) - 55
	} else {
		intvalue = int(r) - 48
	}
	return intvalue
}

// ValidateHKID will check if the given Hong Kong ID number is correctly formatted and
func ValidateHKID(id string) (ok bool, err error) {
	normalizer := regexp.MustCompile(NORMALIZE)
	normalizedStr := normalizer.ReplaceAllString(id, "")
	match, err := regexp.MatchString(VALID_FORMAT, normalizedStr)
	if err != nil {
		return false, err
	}
	if match {
		return validateCheckNumber(normalizedStr), nil
	}
	return match, nil
}
