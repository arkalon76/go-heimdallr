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
	"fmt"
	"regexp"
	"strconv"
)

const (
	// VALID_FORMAT describes a normalized HKID. Example: R5533446
	VALID_FORMAT     = `^[A-NP-Z]{1,2}[0-9]{6}[0-9A]$`
	NORMALIZE        = `[\(\)]`
	CHARACTER_OFFSET = 55
)

// validateCheckNumber will calculate the checksum and validate it to the given code
//
// HK digit pattern
//
// XX999999C / X999999C
func validateCheckNumber(normalizedID string) bool {
	//B123456(6)
	//36 * 9 + char[2] * 8 + char[2] * 8 + char[3] * 7 + char[4] * 6 + char[5] * 5 + char[6] * 4 + char[7] * 3 + char[8] * 2
	sc := []rune(normalizedID)

	// Check if the second character is a letter or not. Below 10 then not a letter
	if int(sc[1])-CHARACTER_OFFSET < 10 {
		sum := 36 * 9
		for i, v := range sc {
			if i == (len(sc) - 1) {
				break
			}
			weight := 8 - i //TODO V needs to be converted to a Int if its and int and THe letter int convertion if its a letter to. Now it all thingks is a letter wich gives us minus
			fmt.Printf("Loop for %s. I=%d, weight=%d, sum=%d, append=%d", normalizedID, i, weight, sum, ((int(v) - CHARACTER_OFFSET) * weight))
			sum += ((int(v) - CHARACTER_OFFSET) * weight)
		}
		// sum := 36*9 + ((int(sc[0])-CHARACTER_OFFSET)*8) + ((int(sc[1])-CHARACTER_OFFSET)*7) + ((int(sc[2])-CHARACTER_OFFSET)*6) + (int(sc[3])-CHARACTER_OFFSET)*5 + (int(sc[4])-CHARACTER_OFFSET)*4 + (int(sc[5])-CHARACTER_OFFSET)*3 + (int(sc[6])-CHARACTER_OFFSET)*2
		checkDigit := 11 - (sum % 11)
		idcheck, err := strconv.Atoi(string(sc[7]))
		if err != nil {
			return false
		}
		if checkDigit != idcheck {
			fmt.Printf("INVALID_ID %s - Check digit is %d and rune 7 is %d\n", normalizedID, checkDigit, idcheck)
			return false
		} else {
			fmt.Printf("VALID_ID %s - Check digit is %d and rune 7 is %d\n", normalizedID, checkDigit, idcheck)
			return true
		}
	}
	return false
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
