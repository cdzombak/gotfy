package gotfy

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageMarshalJSON(mainTest *testing.T) {
	testCases := []struct {
		name        string
		arg         Message
		expected    string
		expectedErr error
	}{
		{
			name:     "base case",
			expected: `{"topic":""}`,
		},
		{
			name:     "topic",
			arg:      Message{Topic: "topic"},
			expected: `{"topic":"topic"}`,
		},
		{
			name:     "Message",
			arg:      Message{Message: "Message"},
			expected: `{"topic":"","message":"Message"}`,
		},
		{
			name:     "Title",
			arg:      Message{Title: "Title"},
			expected: `{"topic":"","title":"Title"}`,
		},
		{
			name:     "Tags",
			arg:      Message{Tags: []string{"tag1", "tag2"}},
			expected: `{"topic":"","tags":["tag1","tag2"]}`,
		},
		{
			name:     "Priority negative",
			arg:      Message{Priority: -1},
			expected: `{"topic":""}`,
		},
		{
			name:     "Priority greater than 0",
			arg:      Message{Priority: 1},
			expected: `{"topic":"","priority":1}`,
		},
		{
			name: "Actions",
			arg: Message{Actions: []ActionButton{&ViewAction{
				Label: "action",
				Link:  &url.URL{Scheme: "http", Host: "host.com"},
				Clear: true,
			}}},
			expected: `{"topic":"","actions":[{"action":"view","label":"action","url":"http://host.com","clear":true}]}`,
		},
		{
			name:     "ClickURL",
			arg:      Message{ClickURL: &url.URL{Scheme: "h", Host: "t.com"}},
			expected: `{"topic":"","click":"h://t.com"}`,
		},
		{
			name:     "IconURL",
			arg:      Message{IconURL: &url.URL{Scheme: "h", Host: "t.com"}},
			expected: `{"topic":"","icon":"h://t.com"}`,
		},
		{
			name:     "Delay",
			arg:      Message{Delay: 1},
			expected: `{"topic":"","delay":"1ns"}`,
		},
		{
			name:     "Email",
			arg:      Message{Email: "Email"},
			expected: `{"topic":"","email":"Email"}`,
		},
		{
			name:     "Call",
			arg:      Message{Call: "Call"},
			expected: `{"topic":"","call":"Call"}`,
		},
		{
			name:     "AttachURLFilename",
			arg:      Message{AttachURLFilename: "AttachURLFilename"},
			expected: `{"topic":"","filename":"AttachURLFilename"}`,
		},
		{
			name:     "AttachURL",
			arg:      Message{AttachURL: &url.URL{Scheme: "h", Host: "t.com"}},
			expected: `{"topic":"","attachurl":"h://t.com"}`,
		},
		{
			name: "everything",
			arg: Message{
				Topic:    "Topic",
				Message:  "Message",
				Title:    "Title",
				Tags:     []string{"tag1", "tag2"},
				Priority: PriorityHigh,
				Actions: []ActionButton{&ViewAction{
					Label: "ajisdiopa",
					Link:  &url.URL{Scheme: "h", Host: "t.com"},
					Clear: true,
				}},
				ClickURL:          &url.URL{Scheme: "h", Host: "t.com"},
				IconURL:           &url.URL{Scheme: "h", Host: "t.com"},
				Delay:             100,
				Email:             "Email",
				Call:              "Call",
				AttachURLFilename: "AttachURLFilename",
				AttachURL:         &url.URL{Scheme: "h", Host: "t.com"},
			},
			expected: `{"topic":"Topic","message":"Message","title":"Title","tags":["tag1","tag2"],"priority":4,"actions":[{"action":"view","label":"ajisdiopa","url":"h://t.com","clear":true}],"click":"h://t.com","attachurl":"h://t.com","icon":"h://t.com","delay":"100ns","email":"Email","call":"Call","filename":"AttachURLFilename"}`,
		},
		{
			name: "test case failure 1/28/2024",
			arg: Message{
				Topic:             "9mm Luger brass 115 grain: $225.00/round, 1000 rounds",
				Title:             "",
				Tags:              []string{Nine},
				Priority:          0,
				Actions:           []ActionButton{},
				ClickURL:          &url.URL{},
				IconURL:           &url.URL{},
				Delay:             0,
				Email:             "",
				Call:              "",
				AttachURLFilename: "",
				AttachURL:         &url.URL{},
				Message:           "ZSR: 9mm - ZSR Buffalo Cartridge 115 Grain Full Metal Jacket - 1000 Rounds 8683262441013 - FREE SHIPPING\n5.0/5 stars, 204 ratings",
			},
			expected: `{"topic":"9mm Luger brass 115 grain: $225.00/round, 1000 rounds","message":"ZSR: 9mm - ZSR Buffalo Cartridge 115 Grain Full Metal Jacket - 1000 Rounds 8683262441013 - FREE SHIPPING\n5.0/5 stars, 204 ratings","tags":["nine"]}`,
		},
	}

	t := assert.New(mainTest)
	for _, tc := range testCases {
		actual, actualErr := json.Marshal(&tc.arg)

		if t.Nil(actualErr, fmt.Sprintf("should not return %s", actualErr)) {
			t.Equal([]byte(tc.expected), actual, tc.name)
		}
	}
}
