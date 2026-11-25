package notifier

import (
	"errors"
	"fmt"
)

type Notifier interface {
	Notify(to, msg string) error
}

var ErrMessageTooLong = errors.New("message cannot be longer than 160 characters")

var _ Notifier = (*EmailNotifier)(nil)

type EmailNotifier struct {
	from string
}

func (n *EmailNotifier) Notify(to, msg string) error {
	runes := []rune(msg)

	if len(runes) > 160 {
		return ErrMessageTooLong
	}

	fmt.Printf("Sending email from %s to %s with content: %s/n", n.from, to, msg)

	return nil
}

type SMSNotifier struct {
	from string
}

func (n *SMSNotifier) Notify(to, msg string) error {
	runes := []rune(msg)

	if len(runes) > 160 {
		return ErrMessageTooLong
	}

	fmt.Printf("Sending sms message from %s to %s with content: %s/n", n.from, to, msg)

	return nil
}

func Broadcast(ns []Notifier, to, msg string) error {
	var res error
	for i := range ns {
		err := ns[i].Notify(to, msg)
		if err != nil {
			errors.Join(res, err)
		}
	}

	return res
}
