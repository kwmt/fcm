package fcm

type downstreamHttpMessage struct {
	To               string        `json:"to,omitempty"`
	RegistrationIds  []string      `json:"registration_ids,omitempty"`
	Priority         string        `json:"priority,omitempty"`
	ContentAvailable bool          `json:"content_available,omitempty"`
	TimeToLive       int64         `json:"time_to_live,omitempty"`
	DryRun           bool          `json:"dry_run,omitempty"`
	notification     *notification `json:"notification,omitempty"`
	//	Data             Data         `json:"data,omitempty"`
}

type notification struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	Sound string `json:"sound,omitempty"`
	Badge int64  `json:"badge,omitempty"`
	Icon  string `json:"icon,omitempty"`
	Tag   string `json:"tag,omitempty"`
}

// TODO
// type Data struct {
// 	data map[string]interface{}
// }

func NewMessage(to string) *downstreamHttpMessage {
	m := new(downstreamHttpMessage)
	m.notification = new(notification)
	m.To = to
	return m
}

func (msg *downstreamHttpMessage) SetPriority(priority string) *downstreamHttpMessage {
	msg.Priority = priority
	return msg
}

func (msg *downstreamHttpMessage) SetTitle(title string) *downstreamHttpMessage {
	msg.notification.Title = title
	return msg
}
func (msg *downstreamHttpMessage) SetBody(body string) *downstreamHttpMessage {
	msg.notification.Body = body
	return msg
}
func (msg *downstreamHttpMessage) SetSound(sound string) *downstreamHttpMessage {
	msg.notification.Sound = sound
	return msg
}
func (msg *downstreamHttpMessage) SetBadge(badge int64) *downstreamHttpMessage {
	msg.notification.Badge = badge
	return msg
}
func (msg *downstreamHttpMessage) SetIcon(icon string) *downstreamHttpMessage {
	msg.notification.Icon = icon
	return msg
}
func (msg *downstreamHttpMessage) SetTag(tag string) *downstreamHttpMessage {
	msg.notification.Tag = tag
	return msg
}
