package gowon

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

type Message struct {
	Module    string            `json:"module"`
	Nick      string            `json:"nick,omitempty"`
	Code      string            `json:"code"`
	Raw       string            `json:"raw"`
	Host      string            `json:"host"`
	Source    string            `json:"source"`
	User      string            `json:"user"`
	Arguments []string          `json:"arguments"`
	Tags      map[string]string `json:"tags"`
	Msg       string            `json:"msg,omitempty"`
	Dest      string            `json:"dest,omitempty"`
	Command   string            `json:"command,omitempty"`
	Args      string            `json:"args,omitempty"`
}

const ErrorMessageParseMsg = "message couldn't be parsed as message json"

const ErrorMessageNoModuleMsg = "message body does not contain a module source"

const ErrorMessageNoBodyMsg = "message body does not contain any message content"

const ErrorMessageNoDestinationMsg = "message body does not contain a destination"

func GetCommand(msg string) string {
	if strings.HasPrefix(msg, ".") {
		return strings.TrimPrefix(strings.Fields(msg)[0], ".")
	}

	return ""
}

func GetArgs(msg string) string {
	if !strings.HasPrefix(msg, ".") {
		return msg
	}

	return strings.TrimSpace(strings.TrimPrefix(msg, strings.Fields(msg)[0]))
}

func CreateMessageStruct(body []byte) (m Message, err error) {
	err = json.Unmarshal(body, &m)
	if err != nil {
		return m, errors.Wrap(err, ErrorMessageParseMsg)
	}

	if m.Module == "" {
		return m, errors.New(ErrorMessageNoModuleMsg)
	}

	if m.Msg == "" {
		return m, errors.New(ErrorMessageNoBodyMsg)
	}

	if m.Dest == "" {
		return m, errors.New(ErrorMessageNoDestinationMsg)
	}

	if m.Command == "" {
		m.Command = GetCommand(m.Msg)
	}

	if m.Args == "" {
		m.Args = GetArgs(m.Msg)
	}

	return m, nil
}
