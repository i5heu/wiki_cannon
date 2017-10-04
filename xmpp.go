package main

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/mattn/go-xmpp"
)

func reciveXMPP() {
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

	for {
		chat, err := client.Recv()
		checkErr(err)

		switch v := chat.(type) {
		case xmpp.Chat:
			bar := strings.Split(v.Remote, "/")
			fmt.Println(v.Remote, v.Text)
			if bar[0] == conf.JabberJIDreciever {
				processXmppText(v.Text)
			}
		case xmpp.Presence:
			fmt.Println(v.From, v.Show)
		}
	}
}

func processXmppText(Text string) { //The Place for the XMPP API Logic
	foo := "ECHO:\n" + Text //Simple echo for dev.
	sendXMPP(foo)
}

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
