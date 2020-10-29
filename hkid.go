package heimdallr

import (
	"regexp"
)

const (
	// VALID_FORMAT describes a normalized HKID. Example: R5533446
	VALID_FORMAT = `^[A-NP-Z]{1,2}[0-9]{6}[0-9A]$`
	NORMALIZE    = `[\(\)]`
)

// ValidateHKID will check if the given Hong Kong ID number is correctly formatted and
func ValidateHKID(id string) (ok bool, err error) {
	normalizer := regexp.MustCompile(NORMALIZE)
	normalizedStr := normalizer.ReplaceAllString(id, "")
	match, err := regexp.MatchString(VALID_FORMAT, normalizedStr)
	if err != nil {
		return false, err
	}
	return match, nil
}
