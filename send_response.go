package gotfy

import (
	"encoding/json"
	"time"
)

// SendResponse is the response from the Ntfy server after successfully sending a message.
type SendResponse struct {
	ID      string   `json:"id"`      // :"bUhbhgmmbeW0"
	Time    UnixTime `json:"time"`    // :1685150791
	Expires UnixTime `json:"expires"` // :1685193991
	Event   string   `json:"event"`   // :"message"
	Topic   string   `json:"topic"`   // :"Server"
	Message string   `json:"message"` // :"triggered"
}

// UnixTime allows unmarshalling a Unix timestamp from JSON into a time.Time.
// See: https://ikso.us/posts/unmarshal-timestamp-as-time/
type UnixTime struct {
	time.Time
}

// UnmarshalJSON unmarshals a Unix timestamp from JSON into a time.Time.
func (u *UnixTime) UnmarshalJSON(b []byte) error {
	var ts int64
	if err := json.Unmarshal(b, &ts); err != nil {
		return err
	}
	if ts <= 0 {
		u.Time = time.Time{}
	} else {
		u.Time = time.Unix(ts, 0)
	}
	return nil
}
