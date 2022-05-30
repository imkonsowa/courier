package utils

import (
	"strings"
)

type Country string

func (c Country) String() string {
	return string(c)
}

const (
	Cameroon   Country = `Cameroon`
	Ethiopia   Country = `Ethiopia`
	Morocco    Country = `Morocco`
	Mozambique Country = `Mozambique`
	Uganda     Country = `Uganda`
)

var countryRegex = map[Country]string{
	Cameroon:   `\(237\)\ ?[2368]\d{7,8}$`,
	Ethiopia:   `\(251\)\ ?[1-59]\d{8}$`,
	Morocco:    `\(212\)\ ?[5-9]\d{8}$`,
	Mozambique: `\(258\)\ ?[28]\d{7,8}$`,
	Uganda:     `\(256\)\ ?\d{9}$`,
}

var countryPrefix = map[Country]string{
	Cameroon:   `237`,
	Ethiopia:   `251`,
	Morocco:    `212`,
	Mozambique: `258`,
	Uganda:     `256`,
}

var AllCountries = []Country{
	Cameroon,
	Ethiopia,
	Morocco,
	Mozambique,
	Uganda,
}

func IsValidCountry(country string) bool {
	for _, c := range AllCountries {
		if Country(country) == c {
			return true
		}
	}

	return false
}

func CountryFromPhone(phone string) Country {
	phone = strings.TrimSpace(phone)

	// TODO: check the provided regexes correctness, many numbers are not matching!
	// HINT: Go regexes are slow anyway ^_____^

	/*for country, regex := range countryRegex {
		match, err := regexp.MatchString(regex, phone)
		if err == nil && match {
			return country
		}
	}*/

	for country, prefix := range countryPrefix {
		if strings.HasPrefix(phone, prefix) {
			return country
		}
	}

	return "NA"
}

func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
