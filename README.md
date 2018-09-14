# InspectData 

[![pipeline status](https://gitlab.com/cjbarker/inspectdata/badges/master/pipeline.svg)](https://gitlab.com/cjbarker/inspectdata/pipelines)
[![coverage report](https://gitlab.com/cjbarker/inspectdata/badges/master/coverage.svg)](https://cjbarker.gitlab.io/inspectdata/test-coverage.html)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/cjbarker/inspectdata)](https://goreportcard.com/report/gitlab.com/cjbarker/inspectdata)
[![GitLab license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://gitlab.com/cjbarker//blob/master/LICENSE)

Go module provides conceptual (canonical) identification of data including Personally Identifiable Information (PII) and Payment Card Industry (PCI).

# Usage
Pull down the package

```bash
go get -u gitlab.com/cjbarker/inspectdata
```

Pass in argument of data to inspect and evaluate results.

```bash
// PCI  via PAN primary account number for credit card
input := "4444444444444448"

datum, err = Inspect(input)

if err != nil {
  // handle error
}

fmt.Printf("%+v\n", datum)
{Data:4444444444444448 DataType:string Canonical:PANVisa IsPII:false IsPCI:true}
```

# Supported Data Inspected
Various analysis via assertion, type checking, and regular expressions applies to inspect and determine any of the following data:

```bash
UIDv4                       // Universally Unique Identifier version 4
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
DateCCYYMMDD                // Date in Century Month Day (optionally with '-', '.', or '/'
PANAmex                     // Payment|Primary Card Number aka credit card number American Express
PANVisa                     // Payment|Primary Card Number aka credit card number Visa
PANMC                       // Payment|Primary Card Number aka credit card number Mastercard
PANDiscover                 // Payment|Primary Card Number aka credit card number Discover
PANDiners                   // Payment|Primary Card Number aka credit card number Diners Club
PANJCB                      // Payment|Primary Card Number aka credit card number JCB
```
