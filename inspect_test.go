package inspectdata

import (
	"testing"
)

func TestBuild(t *testing.T) {
	if len(Version) <= 0 {
		t.Errorf("Version string length was empty, zero or less; got: %s", Version)
	}
	if len(Build) <= 0 {
		t.Errorf("Build string length was empty, zero or less; got: %s", Build)
	}
}

func TestTypeOf(t *testing.T) {
	testStr := "My string"
	v, _ := typeof(testStr)
	if v != "string" {
		t.Errorf("TypeOf should have detected string, but got: %s", v)
	}

	testBool := false
	v, _ = typeof(testBool)
	if v != "bool" {
		t.Errorf("TypeOf should have detected bool, but got: %s", v)
	}

	testInt := 32
	v, _ = typeof(testInt)
	if v != "int" {
		t.Errorf("TypeOf should have detected int, but got: %s", v)
	}

	var testInt8 int8 = -8
	v, _ = typeof(testInt8)
	if v != "int8" {
		t.Errorf("TypeOf should have detected int8, but got: %s", v)
	}

	var testInt16 int16 = 32767
	v, _ = typeof(testInt16)
	if v != "int16" {
		t.Errorf("TypeOf should have detected int16, but got: %s", v)
	}

	var testInt32 int32 = 2147483647
	v, _ = typeof(testInt32)
	if v != "int32" {
		t.Errorf("TypeOf should have detected int32, but got: %s", v)
	}

	var testInt64 int64 = -9223372036854775808
	v, _ = typeof(testInt64)
	if v != "int64" {
		t.Errorf("TypeOf should have detected int64, but got: %s", v)
	}

	var testUint uint
	testUint = 3
	v, _ = typeof(testUint)
	if v != "uint" {
		t.Errorf("TypeOf should have detected uint, but got: %s", v)
	}

	var testUint8 uint8
	testUint8 = 2
	v, _ = typeof(testUint8)
	if v != "uint8" {
		t.Errorf("TypeOf should have detected uint8, but got: %s", v)
	}

	var testUint16 uint16
	testUint16 = 32767
	v, _ = typeof(testUint16)
	if v != "uint16" {
		t.Errorf("TypeOf should have detected uint16, but got: %s", v)
	}

	var testUint32 uint32
	testUint32 = 0
	v, _ = typeof(testUint32)
	if v != "uint32" {
		t.Errorf("TypeOf should have detected uint32, but got: %s", v)
	}

	var testUint64 uint64
	testUint64 = 18446744073709551615
	v, _ = typeof(testUint64)
	if v != "uint64" {
		t.Errorf("TypeOf should have detected uint64, but got: %s", v)
	}

	var testFloat32 float32
	testFloat32 = 3.14
	v, _ = typeof(testFloat32)
	if v != "float32" {
		t.Errorf("TypeOf should have detected float32, but got: %s", v)
	}

	var testFloat64 float64
	testFloat64 = 3.141111111111111111111111111111
	v, _ = typeof(testFloat64)
	if v != "float64" {
		t.Errorf("TypeOf should have detected float64, but got: %s", v)
	}
}

func TestInspectString(t *testing.T) {
	testStr := "My string"
	c, _ := inspectString(testStr)
	if c != Unknown {
		t.Errorf("inspectString should have detected canonical type unknown (plain text), but got: %v", c)
	}

	// validate UUID
	c, _ = inspectString("141a83c3-7f41-4403-9aa2-08a2208b7aa2")
	if c != UUIDv4 {
		t.Errorf("inspectString should have detected canonical type UUIDv4, but got: %v", c)
	}
	c, _ = inspectString("0c76e8e8-b81d-11e8-96f8-529269fb1459")
	if c != Unknown {
		t.Errorf("inspectString should have detected canonical type unknown, but got: %v", c)
	}

	// validate IPv4
	c, _ = inspectString("127.0.0.1")
	if c != IPv4 {
		t.Errorf("inspectString should have detected canonical type IPv4, but got: %v", c)
	}
	c, _ = inspectString("192.168.0.1")
	if c != IPv4 {
		t.Errorf("inspectString should have detected canonical type IPv4, but got: %v", c)
	}
	c, _ = inspectString("255.255.255.255")
	if c != IPv4 {
		t.Errorf("inspectString should have detected canonical type IPv4, but got: %v", c)
	}
	c, _ = inspectString("0.0.0.0")
	if c != IPv4 {
		t.Errorf("inspectString should have detected canonical type IPv4, but got: %v", c)
	}
	c, _ = inspectString("1.1.1.01")
	if c != IPv4 {
		t.Errorf("inspectString should have detected canonical type IPv4, but got: %v", c)
	}
	c, _ = inspectString("30.168.1.255.1")
	if c == IPv4 {
		t.Errorf("inspectString should not have detected canonical type IPv4")
	}
	c, _ = inspectString("1.255.1")
	if c == IPv4 {
		t.Errorf("inspectString should not have detected canonical type IPv4")
	}
	c, _ = inspectString("192.168.1.257")
	if c == IPv4 {
		t.Errorf("inspectString should not have detected canonical type IPv4")
	}
	c, _ = inspectString("-1.2.3.4")
	if c == IPv4 {
		t.Errorf("inspectString should not have detected canonical type IPv4")
	}
	c, _ = inspectString("3...3")
	if c == IPv4 {
		t.Errorf("inspectString should not have detected canonical type IPv4")
	}

	// validate IPv6
	c, _ = inspectString("1:2:3:4:5:6:7:8")
	if c != IPv6 {
		t.Errorf("inspectString should have detected canonical type IPv6, but got: %v", c)
	}
	c, _ = inspectString("2001:db8:3:4::192.0.2.33")
	if c != IPv6 {
		t.Errorf("inspectString should have detected canonical type IPv6, but got: %v", c)
	}
	c, _ = inspectString("::255.255.255.255")
	if c != IPv6 {
		t.Errorf("inspectString should have detected canonical type IPv6, but got: %v", c)
	}
	c, _ = inspectString("::ffff:255.255.255.255")
	if c != IPv6 {
		t.Errorf("inspectString should have detected canonical type IPv6, but got: %v", c)
	}

	// validate Email
	c, _ = inspectString("bob@mail.com")
	if c != Email {
		t.Errorf("inspectString should have detected canonical type email, but got: %v", c)
	}
	c, _ = inspectString("bob@mail")
	if c != Email {
		t.Errorf("inspectString should have detected canonical type email, but got: %v", c)
	}
	c, _ = inspectString("bob@mail-foo.com")
	if c != Email {
		t.Errorf("inspectString should have detected canonical type email, but got: %v", c)
	}
	c, _ = inspectString("bob@mail_foo.com")
	if c != Unknown {
		t.Errorf("inspectString should have detected canonical type unknown, but got: %v", c)
	}
	c, _ = inspectString("ç$€§/az@gmail.com")
	if c != Unknown {
		t.Errorf("inspectString should have detected canonical type unknown, but got: %v", c)
	}

	// validate country codes
	c, _ = inspectString("US")
	if c != CountryCode2 {
		t.Errorf("inspectString should have detected canonical type CountryCode2, but got: %v", c)
	}
	c, _ = inspectString("us")
	if c == CountryCode2 {
		t.Errorf("inspectString should have not detect canonical type CountryCode2")
	}
	c, _ = inspectString("USA")
	if c != CountryCode3 {
		t.Errorf("inspectString should have detected canonical type CountryCode3, but got: %v", c)
	}
	c, _ = inspectString("usa")
	if c == CountryCode3 {
		t.Errorf("inspectString should have not detect canonical type CountryCode3")
	}

	// validate Language Code
	c, _ = inspectString("en")
	if c != LanguageCode2 {
		t.Errorf("inspectString should have detected canonical type LanguageCode3, but got: %v", c)
	}
	c, _ = inspectString("eng")
	if c != LanguageCode3 {
		t.Errorf("inspectString should have detected canonical type LanguageCode3, but got: %v", c)
	}

	// validate postal code
	c, _ = inspectString("90210")
	if c != USPostalCode {
		t.Errorf("inspectString should have detected canonical type USPostalCode, but got: %v", c)
	}
	c, _ = inspectString("90210-1234")
	if c != USPostalCode {
		t.Errorf("inspectString should have detected canonical type USPostalCode, but got: %v", c)
	}
	c, _ = inspectString("902101234")
	if c == USPostalCode {
		t.Errorf("inspectString should not have detected canonical type USPostalCode")
	}

	// validate SSN
	c, _ = inspectString("867-53-0911")
	if c != SSN {
		t.Errorf("inspectString should have detected canonical type SSN, but got: %v", c)
	}
	c, _ = inspectString("867530911")
	if c != SSN {
		t.Errorf("inspectString should have detected canonical type SSN, but got: %v", c)
	}
	c, _ = inspectString("-867530911")
	if c == SSN {
		t.Errorf("inspectString should not have detected canonical type SSN")
	}

	// validate USD
	c, _ = inspectString("$1.02")
	if c != USD {
		t.Errorf("inspectString should have detected canonical type USD, but got: %v", c)
	}
	c, _ = inspectString("$ 1.02")
	if c != USD {
		t.Errorf("inspectString should have detected canonical type USD, but got: %v", c)
	}
	c, _ = inspectString("1.01")
	if c != USD {
		t.Errorf("inspectString should have detected canonical type USD, but got: %v", c)
	}
	c, _ = inspectString(".01")
	if c == USD {
		t.Errorf("inspectString should not have detected canonical type USD")
	}
	c, _ = inspectString("0.01")
	if c != USD {
		t.Errorf("inspectString should have detected canonical type USD, but got: %v", c)
	}
	c, _ = inspectString("+0.01")
	if c != USD {
		t.Errorf("inspectString should have detected canonical type USD, but got: %v", c)
	}
	c, _ = inspectString("-0.01")
	if c != USD {
		t.Errorf("inspectString should have detected canonical type USD, but got: %v", c)
	}
	c, _ = inspectString("$-0.01")
	if c != USD {
		t.Errorf("inspectString should have detected canonical type USD, but got: %v", c)
	}
	c, _ = inspectString("$ +0.01")
	if c != USD {
		t.Errorf("inspectString should have detected canonical type USD, but got: %v", c)
	}
	c, _ = inspectString("$ 1.0")
	if c == USD {
		t.Errorf("inspectString should not have detected canonical type USD")
	}

	// validate LatLong
	c, _ = inspectString("+90.0, -127.554334")
	if c != LatLong {
		t.Errorf("inspectString should have detected canonical type LatLong, but got: %v", c)
	}
	c, _ = inspectString("23,120")
	if c != LatLong {
		t.Errorf("inspectString should have detected canonical type LatLong, but got: %v", c)
	}
	c, _ = inspectString("-90, -122")
	if c != LatLong {
		t.Errorf("inspectString should have detected canonical type LatLong, but got: %v", c)
	}
	c, _ = inspectString("-90.000, -122.0000")
	if c != LatLong {
		t.Errorf("inspectString should have detected canonical type LatLong, but got: %v", c)
	}
	c, _ = inspectString("+90, +122")
	if c != LatLong {
		t.Errorf("inspectString should have detected canonical type LatLong, but got: %v", c)
	}
	c, _ = inspectString("47.1231231, 179.99999999")
	if c != LatLong {
		t.Errorf("inspectString should have detected canonical type LatLong, but got: %v", c)
	}
	c, _ = inspectString("-90., -122.")
	if c == LatLong {
		t.Errorf("inspectString should not have detected canonical type LatLong")
	}
	c, _ = inspectString("+90.1, -100.111")
	if c == LatLong {
		t.Errorf("inspectString should not have detected canonical type LatLong")
	}
	c, _ = inspectString("-91, 123.456")
	if c == LatLong {
		t.Errorf("inspectString should not have detected canonical type LatLong")
	}
	c, _ = inspectString("012, 122")
	if c == LatLong {
		t.Errorf("inspectString should not have detected canonical type LatLong")
	}

	// valid CCYYMMDD
	c, _ = inspectString("20181011")
	if c != DateCCYYMMDD {
		t.Errorf("inspectString should have detected canonical type DateCCYYMMDD, but got: %v", c)
	}
	c, _ = inspectString("2018-10-11")
	if c != DateCCYYMMDD {
		t.Errorf("inspectString should have detected canonical type DateCCYYMMDD, but got: %v", c)
	}
	c, _ = inspectString("2018/10/11")
	if c != DateCCYYMMDD {
		t.Errorf("inspectString should have detected canonical type DateCCYYMMDD, but got: %v", c)
	}
	c, _ = inspectString("2018.10.11")
	if c != DateCCYYMMDD {
		t.Errorf("inspectString should have detected canonical type DateCCYYMMDD, but got: %v", c)
	}
	c, _ = inspectString("20180401")
	if c != DateCCYYMMDD {
		t.Errorf("inspectString should have detected canonical type DateCCYYMMDD, but got: %v", c)
	}
	c, _ = inspectString("20180421")
	if c != DateCCYYMMDD {
		t.Errorf("inspectString should have detected canonical type DateCCYYMMDD, but got: %v", c)
	}
	c, _ = inspectString("19000101")
	if c != DateCCYYMMDD {
		t.Errorf("inspectString should have detected canonical type DateCCYYMMDD, but got: %v", c)
	}
	c, _ = inspectString("20991231")
	if c != DateCCYYMMDD {
		t.Errorf("inspectString should have detected canonical type DateCCYYMMDD, but got: %v", c)
	}
	c, _ = inspectString("20991331")
	if c == DateCCYYMMDD {
		t.Errorf("inspectString should not have detected canonical type DateCCYYMMDD")
	}
	c, _ = inspectString("20992920")
	if c == DateCCYYMMDD {
		t.Errorf("inspectString should not have detected canonical type DateCCYYMMDD")
	}
	c, _ = inspectString("21000101")
	if c == DateCCYYMMDD {
		t.Errorf("inspectString should not have detected canonical type DateCCYYMMDD")
	}

	// valid credit cards
	c, _ = inspectString("371449635398431")
	if c != PANAmex {
		t.Errorf("inspectString should have detected canonical type PANAmex, but got: %v", c)
	}
	c, _ = inspectString("36438936438936")
	if c != PANDiners {
		t.Errorf("inspectString should have detected canonical type PANDiners, but got: %v", c)
	}
	c, _ = inspectString("3566003566003566")
	if c != PANJCB {
		t.Errorf("inspectString should have detected canonical type PANJCB, but got: %v", c)
	}
	c, _ = inspectString("4444444444444448")
	if c != PANVisa {
		t.Errorf("inspectString should have detected canonical type PANVisa, but got: %v", c)
	}
	c, _ = inspectString("5500005555555559")
	if c != PANMC {
		t.Errorf("inspectString should have detected canonical type PANMC, but got: %v", c)
	}
	c, _ = inspectString("6011016011016011")
	if c != PANDiscover {
		t.Errorf("inspectString should have detected canonical type PANDiscover, but got: %v", c)
	}
}

func TestInspect(t *testing.T) {
	input := "My string"
	datum, err := Inspect(input)
	if err == nil {
		t.Errorf("Inspect should not be able to inspect plain text string without unknown error")
	}
	if datum.Canonical != Unknown {
		t.Errorf("Inspect canonical data type should be unknown, but got: %v", datum.Canonical)
	}

	// test PCI
	input = "4444444444444448"
	datum, err = Inspect(input)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !datum.IsPCI {
		t.Errorf("VISA credit card number data should be denoted as PCI")
	}
	if datum.IsPII {
		t.Errorf("VISA credit card number data should not be denoted as PII")
	}

	//fmt.Printf("%+v\n", datum)

	// test PII
	input = "bob@mail.com"
	datum, err = Inspect(input)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !datum.IsPII {
		t.Errorf("Email data should be denoted as PII")
	}
	if datum.IsPCI {
		t.Errorf("Email data should not be denoted as PCI")
	}

	// test non PII, PCI
	input = "20180914"
	datum, err = Inspect(input)
	if err != nil {
		t.Errorf(err.Error())
	}
	if datum.IsPII {
		t.Errorf("CCYYMMDD data should not be denoted as PII")
	}
	if datum.IsPCI {
		t.Errorf("CCYYMMDD data should not be denoted as PCI")
	}
}
