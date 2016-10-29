package fcm

type DownstreamHttpMessage struct {
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

func NewMessage(to string) *DownstreamHttpMessage {
	m := new(DownstreamHttpMessage)
	m.notification = new(notification)
	m.To = to
	return m
}

func (msg *DownstreamHttpMessage) SetPriority(priority string) *DownstreamHttpMessage {
	msg.Priority = priority
	return msg
}

func (msg *DownstreamHttpMessage) SetTitle(title string) *DownstreamHttpMessage {
	msg.notification.Title = title
	return msg
}
func (msg *DownstreamHttpMessage) SetBody(body string) *DownstreamHttpMessage {
	msg.notification.Body = body
	return msg
}
func (msg *DownstreamHttpMessage) SetSound(sound string) *DownstreamHttpMessage {
	msg.notification.Sound = sound
	return msg
}
func (msg *DownstreamHttpMessage) SetBadge(badge int64) *DownstreamHttpMessage {
	msg.notification.Badge = badge
	return msg
}
func (msg *DownstreamHttpMessage) SetIcon(icon string) *DownstreamHttpMessage {
	msg.notification.Icon = icon
	return msg
}
func (msg *DownstreamHttpMessage) SetTag(tag string) *DownstreamHttpMessage {
	msg.notification.Tag = tag
	return msg
}
