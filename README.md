# fcm
[![CircleCI](https://circleci.com/gh/kwmt/fcm.svg?style=svg&circle-token=1a97631ce162453dd004f748bc276acec2d3c2c9)](https://circleci.com/gh/kwmt/fcm)


Usage
-----

```go
msg := fcm.NewMessage("/topics/all").SetPriority(fcm.High)

resp, _ := fcm.NewClient("<YOUR SERVER KEY>").Send(msg)
fmt.Println(resp.MessageId)
```

Install
-------

```
go get github.com/kwmt/fcm
```
