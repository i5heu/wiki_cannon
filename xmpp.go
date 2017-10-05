package main

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/mattn/go-xmpp"
)

func newXmppClient() (client *xmpp.Client) {
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

	return
}

func reciveXMPP() {

	client := newXmppClient()

	for {
		chat, err := client.Recv()
		fmt.Println(">>>>> ", chat)

		if err != nil {

			if err.Error() == "EOF" {
				time.Sleep(1 * time.Second)
				client = newXmppClient()
			}

		}

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
	TextSplitt := strings.Split(Text, "\n")

	switch TextSplitt[0] {
	case "#ping", "#Ping":
		sendXMPP("Pong - Im Online!")
	case "#echo", "#Echo":
		foo := "ECHO:\n" + Text //Simple echo for dev.
		sendXMPP(foo)
	case "hello", "Hello", "Hi", "hi":
		sendXMPP("Hello im a Wiki-Cannon Jabber Bot \n for more informations visit \n > https://github.com/i5heu/wiki_cannon")
	case "h", "help", "Help", "H":
		sendXMPP("--- HELP ---\nType folowing comands for default Template:\n\ngeldlog")
	case "geldlog", "geld", "Geld", "Geldlog":
		sendXMPP("#geldlog\nName=\"\nVal=\nTags=\"\nCat=\"\nText=\"")
	case "#geldlog", "#Geldlog":
		GeldLogXMPP(Text)
	default:
		sendXMPP("No Method found")
	}

}

func GeldLogXMPP(Text string) {
	type Data struct {
		Name string
		Tag  string
		Cat  string
		Text string
		Val  int
	}

	var data Data

	if _, err := toml.Decode(Text, &data); err != nil {
		// handle error
		sendXMPP(err.Error())
		return
	}

	jsondata := API2STRUCT{
		APPWRITE: "Geldlog",
		Title1:   data.Name,
		Title2:   data.Cat,
		Text1:    data.Text,
		Tags1:    data.Tag,
		Num1:     data.Val,
	}

	ItemWriter(jsondata)
	foo := "--- Geldlog Saved ---"
	sendXMPP(foo)
}

func sendXMPP(message string) {
	if conf.JabberNotification == false {
		return
	}

	client := newXmppClient()

	client.Send(xmpp.Chat{
		Remote: conf.JabberJIDreciever,
		Type:   "chat",
		Text:   message,
	})
}
