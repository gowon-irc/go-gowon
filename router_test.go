package gowon

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMessageRouter(t *testing.T) {
	mr := NewMessageRouter()

	assert.NotNil(t, mr.Commands)
	assert.NotNil(t, mr.Regexes)
}

func TestMessageAddCommand(t *testing.T) {
	mr := NewMessageRouter()
	mr.AddCommand("test", func(Message) (string, error) { return "testing", nil })

	assert.Contains(t, mr.Commands, "test")
}

func TestMessageAddRegex(t *testing.T) {
	mr := NewMessageRouter()
	mr.AddRegex("test", func(Message) (string, error) { return "testing", nil })

	assert.Contains(t, mr.Regexes, "test")
}

func createTestMessage(command, msg string) Message {
	return Message{
		Command: command,
		Msg:     msg,
	}
}

func createTestMessageRouter(commands, regexes int) *MessageRouter {
	mr := NewMessageRouter()

	for i := 1; i <= commands; i++ {
		c := fmt.Sprintf("test%s", fmt.Sprint(i))

		mr.AddCommand(c, func(Message) (string, error) { return c, nil })
	}

	for i := 1; i <= regexes; i++ {
		r := fmt.Sprintf("regex %s", fmt.Sprint(i))

		mr.AddRegex(r, func(Message) (string, error) { return r, nil })
	}

	return mr
}

func TestMessageRoute(t *testing.T) {
	cases := []struct {
		name     string
		m        Message
		mr       *MessageRouter
		expected string
	}{
		{
			name:     "single command",
			m:        createTestMessage("test1", ""),
			mr:       createTestMessageRouter(1, 0),
			expected: "test1",
		},
		{
			name:     "single regex",
			m:        createTestMessage("", "this contains regex 1 :)"),
			mr:       createTestMessageRouter(0, 1),
			expected: "regex 1",
		},
		{
			name:     "command and regex pick command",
			m:        createTestMessage("test1", "regex 1"),
			mr:       createTestMessageRouter(1, 1),
			expected: "test1",
		},
		{
			name:     "command and regex pick regex",
			m:        createTestMessage("test2", "regex 1"),
			mr:       createTestMessageRouter(1, 1),
			expected: "regex 1",
		},
		{
			name:     "second command",
			m:        createTestMessage("test2", ""),
			mr:       createTestMessageRouter(2, 0),
			expected: "test2",
		},
		{
			name:     "second regex",
			m:        createTestMessage("", "regex 2"),
			mr:       createTestMessageRouter(0, 2),
			expected: "regex 2",
		},
		{
			name:     "empty router",
			m:        createTestMessage("test1", ""),
			mr:       createTestMessageRouter(0, 0),
			expected: "",
		},
		{
			name:     "no match",
			m:        createTestMessage("test3", ""),
			mr:       createTestMessageRouter(2, 2),
			expected: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out, _ := tc.mr.Route(tc.m)

			assert.Equal(t, tc.expected, out)
		})
	}
}
