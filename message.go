package fcm

type MessageData struct {
	Message *DownstreamHttpMessage `json:"message,omitempty"`
}

// see https://firebase.google.com/docs/cloud-messaging/http-server-ref?hl=en#downstream-http-messages-json
type DownstreamHttpMessage struct {
	Token            string        `json:"token,omitempty"`
	Topic            string        `json:"topic,omitempty"`
	RegistrationIds  []string      `json:"registration_ids,omitempty"`
	Priority         string        `json:"priority,omitempty"`
	ContentAvailable bool          `json:"content_available,omitempty"`
	TimeToLive       int64         `json:"time_to_live,omitempty"`
	DryRun           bool          `json:"dry_run,omitempty"`
	Notification     *notification `json:"notification,omitempty"`
	data
}

// see https://firebase.google.com/docs/cloud-messaging/http-server-ref?hl=en#notification-payload-support
type notification struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	Sound string `json:"sound,omitempty"`
	Badge int64  `json:"badge,omitempty"`
	Icon  string `json:"icon,omitempty"`
	Tag   string `json:"tag,omitempty"`
}

type data struct {
	Data map[string]interface{} `json:"data,omitempty"`
}

// create *MessageData instance
func NewMessage(topic string) *MessageData {
	message := new(MessageData)
	message.Message = new(DownstreamHttpMessage)
	message.Message.Notification = new(notification)
	message.Message.Topic = topic
	return message
}

func NewMessageWithToken(token string) *MessageData {
	message := new(MessageData)
	message.Message = new(DownstreamHttpMessage)
	message.Message.Notification = new(notification)
	message.Message.Token = token
	return message
}

func (msg *MessageData) SetPriority(priority string) *MessageData {
	msg.Message.Priority = priority
	return msg
}

func (msg *MessageData) SetTitle(title string) *MessageData {
	msg.Message.Notification.Title = title
	return msg
}
func (msg *MessageData) SetBody(body string) *MessageData {
	msg.Message.Notification.Body = body
	return msg
}
func (msg *MessageData) SetSound(sound string) *MessageData {
	msg.Message.Notification.Sound = sound
	return msg
}
func (msg *MessageData) SetBadge(badge int64) *MessageData {
	msg.Message.Notification.Badge = badge
	return msg
}
func (msg *MessageData) SetIcon(icon string) *MessageData {
	msg.Message.Notification.Icon = icon
	return msg
}
func (msg *MessageData) SetTag(tag string) *MessageData {
	msg.Message.Notification.Tag = tag
	return msg
}
func (msg *MessageData) SetData(data map[string]interface{}) *MessageData {
	msg.Message.Data = data
	return msg
}
