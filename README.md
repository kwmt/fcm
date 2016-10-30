# fcm
[![CircleCI](https://circleci.com/gh/kwmt/fcm.svg?style=svg&circle-token=1a97631ce162453dd004f748bc276acec2d3c2c9)](https://circleci.com/gh/kwmt/fcm) [![GoDoc](https://godoc.org/github.com/kwmt/fcm?status.svg)](http://godoc.org/github.com/kwmt/fcm) 


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
