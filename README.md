# fcm
[![Go](https://github.com/kwmt/fcm/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/kwmt/fcm/actions/workflows/go.yml) 
[![GoDoc](https://godoc.org/github.com/kwmt/fcm?status.svg)](http://godoc.org/github.com/kwmt/fcm) 
[![Go Report Card](https://goreportcard.com/badge/github.com/kwmt/fcm)](https://goreportcard.com/report/github.com/kwmt/fcm)

Usage
-----

```go
// send push message
msg := fcm.NewMessage("/topics/all").SetPriority(fcm.High)

c := fcm.NewClient("<YOUR SERVER KEY>")
resp, _ := c.Send(msg)
fmt.Println(resp)

// get regitration token info
respInfo, _ := c.GetRegistrationTokenInfo("<REGISTRATION TOKEN>")
fmt.Println(respInfo)
```

Install
-------

```
go get github.com/kwmt/fcm
```
