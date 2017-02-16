package fcm

type DownstreamHttpMessage struct {
	To               string        `json:"to,omitempty"`
	RegistrationIds  []string      `json:"registration_ids,omitempty"`
	Priority         string        `json:"priority,omitempty"`
	ContentAvailable bool          `json:"content_available,omitempty"`
	TimeToLive       int64         `json:"time_to_live,omitempty"`
	DryRun           bool          `json:"dry_run,omitempty"`
	Notification     *notification `json:"notification,omitempty"`
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
	m.Notification = new(notification)
	m.To = to
	return m
}

func (msg *DownstreamHttpMessage) SetPriority(priority string) *DownstreamHttpMessage {
	msg.Priority = priority
	return msg
}

func (msg *DownstreamHttpMessage) SetTitle(title string) *DownstreamHttpMessage {
	msg.Notification.Title = title
	return msg
}
func (msg *DownstreamHttpMessage) SetBody(body string) *DownstreamHttpMessage {
	msg.Notification.Body = body
	return msg
}
func (msg *DownstreamHttpMessage) SetSound(sound string) *DownstreamHttpMessage {
	msg.Notification.Sound = sound
	return msg
}
func (msg *DownstreamHttpMessage) SetBadge(badge int64) *DownstreamHttpMessage {
	msg.Notification.Badge = badge
	return msg
}
func (msg *DownstreamHttpMessage) SetIcon(icon string) *DownstreamHttpMessage {
	msg.Notification.Icon = icon
	return msg
}
func (msg *DownstreamHttpMessage) SetTag(tag string) *DownstreamHttpMessage {
	msg.Notification.Tag = tag
	return msg
}
