package fcm_test

import (
	"fmt"
	"github.com/kwmt/fcm"
)

func ExampleSend() {
	// send push message
	msg := fcm.NewMessage("/topics/all").SetPriority(fcm.High)

	c := fcm.NewClient("<YOUR SERVER KEY>")
	resp, _ := c.Send(msg)
	fmt.Println(resp)
}

func ExampleGetRegistrationTokenInfo() {
	// get regitration token info
	c := fcm.NewClient("<YOUR SERVER KEY>")
	resp, _ := c.GetRegistrationTokenInfo("<REGISTRATION TOKEN>")
	fmt.Println(resp)
}
