package gotfy

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Message represents a notification to be sent to the specified topic.
// See: https://docs.ntfy.sh/publish/
type Message struct {
	Topic string `json:"topic"`           // Target topic name.
	Email string `json:"email,omitempty"` // Address for e-mail notifications. See: https://docs.ntfy.sh/publish/#e-mail-notifications
	Call  string `json:"call,omitempty"`  // Phone number for voice call. See: https://docs.ntfy.sh/publish/#phone-calls

	Message  string         `json:"message,omitempty"`  // Message body.
	Title    string         `json:"title,omitempty"`    // Message title. See: https://docs.ntfy.sh/publish/#message-title
	Tags     []string       `json:"tags,omitempty"`     // List of tags that may or not map to emojis. See: https://docs.ntfy.sh/publish/#tags-emojis
	Priority Priority       `json:"priority,omitempty"` // Message priority with 1=min, 3=default and 5=max. See: https://docs.ntfy.sh/publish/#message-priority
	Actions  []ActionButton `json:"actions,omitempty"`  // Custom user action buttons for notifications. See: https://docs.ntfy.sh/publish/#action-buttons
	ClickURL *url.URL       `json:"click,omitempty"`    // Website to open when notification is clicked. See: https://docs.ntfy.sh/publish/#click-action
	IconURL  *url.URL       `json:"icon,omitempty"`     // URL to use as notification icon. See: https://docs.ntfy.sh/publish/#icons

	Delay time.Duration `json:"delay,omitempty"` // Duration by which to delay delivery. See: https://docs.ntfy.sh/publish/#scheduled-delivery

	AttachURL         *url.URL `json:"attachurl,omitempty"` // URL of an attachment. See: https://docs.ntfy.sh/publish/#attach-file-from-a-url
	AttachURLFilename string   `json:"filename,omitempty"`  // User-facing file name for the attachment pointed to by AttachURL.
}

func (m *Message) MarshalJSON() ([]byte, error) {
	buf, err := json.Marshal(m.Topic)
	if err != nil {
		return nil, err
	}
	buf = []byte(fmt.Sprintf(`{"topic":%s`, buf))

	if x := m.Message; x != "" {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"message":%s`, mm)...)
	}

	if x := m.Title; x != "" {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"title":%s`, mm)...)
	}

	if x := m.Tags; len(x) > 0 {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"tags":%s`, mm)...)
	}

	if x := m.Priority; x > 0 {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"priority":%s`, mm)...)
	}

	if x := m.Actions; len(x) > 0 {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"actions":%s`, mm)...)
	}

	type urls struct {
		name string
		url  *url.URL
	}

	for _, v := range []urls{
		{"click", m.ClickURL},
		{"attachurl", m.AttachURL},
		{"icon", m.IconURL},
	} {
		mm, err := urlString(v.url)
		if err != nil {
			return nil, err
		}

		if mm == nil {
			continue
		}

		buf = append(buf, fmt.Sprintf(`,"%s":%s`, v.name, mm)...)
	}

	if x := m.Delay; x > 0 {
		mm, err := json.Marshal(x.String())
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"delay":%s`, mm)...)
	}

	if x := m.Email; x != `` {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"email":%s`, mm)...)
	}

	if x := m.Call; x != `` {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"call":%s`, mm)...)
	}

	if x := m.AttachURLFilename; x != `` {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"filename":%s`, mm)...)
	}

	return append(buf, '}'), nil
}

func urlString(u *url.URL) ([]byte, error) {
	if u == nil {
		return nil, nil
	}

	s := u.String()
	if s == "" {
		return nil, nil
	}

	return json.Marshal(s)
}
