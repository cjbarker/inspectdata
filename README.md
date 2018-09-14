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
