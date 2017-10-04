package main

import (
	"crypto/tls"
	"fmt"

	"github.com/mattn/go-xmpp"
)

func sendXMPP(message string) {
	if conf.JabberNotification == false {
		return
	}

	options := xmpp.Options{
		Host:     conf.JabberHost,
		User:     conf.JabberUser,
		Password: conf.JabberPassword,
		NoTLS:    true,
		StartTLS: true,
		TLSConfig: &tls.Config{
			ServerName:         conf.JabberJIDreciever,
			InsecureSkipVerify: conf.JabberInsecureSkipVerify,
		},
		Debug:         false,
		Session:       true,
		Status:        "xa",
		StatusMessage: "Hello! I'm using XMPP",
	}

	client, err := options.NewClient()
	if err != nil {
		fmt.Println(err)
	}

	client.Send(xmpp.Chat{
		Remote: conf.JabberJIDreciever,
		Type:   "chat",
		Text:   message,
	})
}
