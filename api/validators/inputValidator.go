package validators

import (
	"regexp"
	"strings"
)

func ValidateName(name string) bool {

	if name != "" && regexp.MustCompile("^[a-zA-Z ]{2,30}$").MatchString(name) {
		return true
	}
	return false
}

func ValidateContactDetails(ContactType string, ContactValue string) (bool, string) {
	if strings.EqualFold(ContactType, "phone-number") || strings.EqualFold(ContactType, "email") {
		if strings.EqualFold(ContactType, "phone-number") {
			if regexp.MustCompile("^[0-9]{10}$").MatchString(ContactValue) {
				return true, "Operation Successful"
			}
			return false, "Invalid phone number"
		}

		if regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString(ContactValue) {
			return true, "Operation Successful"
		}
		return false, "Invalid email Id"
	}
	return false, "Invalid type of contact. Please enter valid phone number or Email"
}
