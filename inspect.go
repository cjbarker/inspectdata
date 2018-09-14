package inspectdata

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Variables to identify the build
var (
	Version string
	Build   string
)

// Denotes the canonical type for a given piece of string data ex: IP address, email, or UUID
type CanonicalType int

// Inspected Data Types
const (
	Unknown       CanonicalType = iota
	UUIDv4                      // Universally Unique Identifier version 4
	IPv4                        // IP Address version 4
	IPv6                        // IP address version 6
	Email                       // Email address
	CountryCode2                // Country Code ISO ALPHA-2 Code
	CountryCode3                // Country Code ISO ALPHA-3 Code
	LanguageCode2               // Language Code ISO 639-1
	LanguageCode3               // Lanuage Code ISO 639-2/T
	USPostalCode                // USA postal code 5 digit or 5-4
	SSN                         // Social Security Number
	USD                         // USA Currency
	LatLong                     // Latitude, Longitude Geocoordinates
	DateCCYYMMDD                // Date in Century Month Daty (optionally with '-', '.', or '/'
	PANAmex                     // Payment|Primary Card Number aka credit card number American Express
	PANVisa                     // Payment|Primary Card Number aka credit card number Visa
	PANMC                       // Payment|Primary Card Number aka credit card number Mastercard
	PANDiscover                 // Payment|Primary Card Number aka credit card number Discover
	PANDiners                   // Payment|Primary Card Number aka credit card number Diner's Club
	PANJCB                      // Payment|Primary Card Number aka credit card number JCB
)

// Canonical structure representing a given piece of data aka datum.
// Handles inspecting numerous forms of data and applying conceptual/canonical determination.
// Example data includes, but is not limited to: IP address, UUID, SSN, Lat/Long, Credit Cards and more.
type Datum struct {
	Data      interface{}   // Actual atomic data value
	DataType  string        // Represents data type ex: string, int, bool, float32, etc.
	Canonical CanonicalType // Canonical inspected data type ex: UUIDv4, IPv4, SSN, etc.
	IsPII     bool          // Denotes if considered Personally Identifiable Information (ex: email addr)
	IsPCI     bool          // Denotes if considered Payment Card Industry data (ex: credit card no.)
}

// Regular Expressions for Data Type Inspection
const reUUIDv4 = "^[0-9a-f]{8}-[0-9af]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
const reEmail = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
const reLatLong = `^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`
const reSSN = "^[0-9]{3}-?[0-9]{2}-?[0-9]{4}$"
const reUSPostal = "^[0-9]{5}(-[0-9]{4})?$"
const reUSD = `^\$?[ ]?[+-]?[0-9]{1,3}(?:,?[0-9]{3})*(?:\.[0-9]{2})$`
const reLangCode2 = "^[a-z]{2}$"
const reLangCode3 = "^[a-z]{3}$"
const reCountryCode2 = "^[A-Z]{2}$"
const reCountryCode3 = "^[A-Z]{3}$"
const reCCYYMMDD = `(19[0-9]{2}|20[0-9]{2})(-|/|.)?(0[1-9]|1[012])(-|/|.)?(0[1-9]|1[0-9]|2[0-9]|3[01])` // years 1900-2099
const rePANAmex = "^3[47][0-9]{13}$"
const rePANMC = "^5[1-5][0-9]{14}$"
const rePANVisa = "^4[0-9]{12}(?:[0-9]{3})?$"
const rePANDiners = "^3(?:0[0-5]|[68][0-9])[0-9]{11}$"
const rePANJCB = "^(?:2131|1800|35[0-9]{3})[0-9]{11}$"
const rePANDiscover = "^65[4-9][0-9]{13}|64[4-9][0-9]{13}|6011[0-9]{12}|(622(?:12[6-9]|1[3-9][0-9]|[2-8][0-9][0-9]|9[01][0-9]|92[0-5])[0-9]{10})$"
const reIPv4 = `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
const reIPv6 = `^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$`

// Inspects data determining its canonically representation and associated meta-data
// When error is nil it will always contain non-nil Datum
func Inspect(v interface{}) (datum Datum, err error) {
	datum = Datum{
		Data: v,
	}

	datum.DataType, err = typeof(v)
	if err != nil {
		return datum, err
	}

	str := v.(string)
	datum.Canonical, err = inspectString(str)
	if err != nil {
		return datum, err
	}

	switch datum.Canonical {
	case UUIDv4, IPv4, IPv6, Email, SSN:
		datum.IsPII = true
	case PANAmex, PANMC, PANVisa, PANDiscover, PANDiners, PANJCB:
		datum.IsPCI = true
	default:
		datum.IsPCI = false
		datum.IsPII = false
	}

	return datum, nil
}

// Determine data type via string formatting or assertion.
func typeof(v interface{}) (string, error) {
	strType := fmt.Sprintf("%T", v)

	if strType == "string" || strType == "bool" {
		return strType, nil
	}

	// type assertion for more granularity
	switch v.(type) {
	case int:
		return "int", nil
	case int8:
		return "int8", nil
	case int16:
		return "int16", nil
	case int32:
		return "int32", nil
	case int64:
		return "int64", nil
	case uint:
		return "uint", nil
	case uint8:
		return "uint8", nil
	case uint16:
		return "uint16", nil
	case uint32:
		return "uint32", nil
	case uint64:
		return "uint64", nil
	case float32:
		return "float32", nil
	case float64:
		return "float64", nil
	default:
		return "unknown", errors.New("Unable to determine data type for given string parameter - unknown")
	}
}

// Inspects the string to determine its CanonicalType based on series of regular expressions
func inspectString(v string) (CanonicalType, error) {
	// regular expressions to inspect
	var validUUID = regexp.MustCompile(reUUIDv4)
	var validIPv4 = regexp.MustCompile(reIPv4)
	var validIPv6 = regexp.MustCompile(reIPv6)
	var validEmail = regexp.MustCompile(reEmail)
	var validLatLong = regexp.MustCompile(reLatLong)
	var validCountryCode2 = regexp.MustCompile(reCountryCode2)
	var validCountryCode3 = regexp.MustCompile(reCountryCode3)
	var validLanguageCode2 = regexp.MustCompile(reLangCode2)
	var validLanguageCode3 = regexp.MustCompile(reLangCode3)
	var validUSPostalCode = regexp.MustCompile(reUSPostal)
	var validSSN = regexp.MustCompile(reSSN)
	var validUSD = regexp.MustCompile(reUSD)
	var validCCYYMMDD = regexp.MustCompile(reCCYYMMDD)
	var validPANAmex = regexp.MustCompile(rePANAmex)
	var validPANDiners = regexp.MustCompile(rePANDiners)
	var validPANJCB = regexp.MustCompile(rePANJCB)
	var validPANMC = regexp.MustCompile(rePANMC)
	var validPANVisa = regexp.MustCompile(rePANVisa)
	var validPANDiscover = regexp.MustCompile(rePANDiscover)

	if validUUID.MatchString(strings.ToLower(v)) {
		return UUIDv4, nil
	} else if validIPv4.MatchString(v) {
		return IPv4, nil
	} else if validIPv6.MatchString(v) {
		return IPv6, nil
	} else if validEmail.MatchString(v) {
		return Email, nil
	} else if validLatLong.MatchString(v) {
		return LatLong, nil
	} else if validCountryCode2.MatchString(v) {
		return CountryCode2, nil
	} else if validCountryCode3.MatchString(v) {
		return CountryCode3, nil
	} else if validLanguageCode2.MatchString(v) {
		return LanguageCode2, nil
	} else if validLanguageCode3.MatchString(v) {
		return LanguageCode3, nil
	} else if validUSPostalCode.MatchString(v) {
		return USPostalCode, nil
	} else if validSSN.MatchString(v) {
		return SSN, nil
	} else if validUSD.MatchString(v) {
		return USD, nil
	} else if validCCYYMMDD.MatchString(v) {
		return DateCCYYMMDD, nil
	} else if validPANAmex.MatchString(v) {
		return PANAmex, nil
	} else if validPANDiners.MatchString(v) {
		return PANDiners, nil
	} else if validPANMC.MatchString(v) {
		return PANMC, nil
	} else if validPANVisa.MatchString(v) {
		return PANVisa, nil
	} else if validPANJCB.MatchString(v) {
		return PANJCB, nil
	} else if validPANDiscover.MatchString(v) {
		return PANDiscover, nil
	} else {
		return Unknown, errors.New("Unable to determine canonical data - unknown")
	}
}
