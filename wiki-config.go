package main

import (
	"crypto/sha256"
	"encoding/hex"
)

var guestmode bool = true
var dblogin string = "USER:PASSWORD@/wiki_cannon"
var personalpwdTMP string = "wiki"
var guestpwd string = "cannon"

var templatefolder string = "/home/her/CODE/wiki_cannon/template"

/*############# END OF CONFIG ################*/
var wcversion string = "WC-Beta-1.2.23"

var foo = sha256.Sum256([]byte(personalpwdTMP))
var personalpwd = hex.EncodeToString(foo[:])
