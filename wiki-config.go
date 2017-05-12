package main

import (
	"crypto/sha256"
	"encoding/hex"
)

var guestmode bool = true
var dblogin string = "USER:PASSWORD@/wiki_cannon"
var personalpwdTMP string = "wiki"
var guestpwd string = "cannon"

/*############# END OF CONFIG ################*/

var foo = sha256.Sum256([]byte(personalpwdTMP))
var personalpwd = hex.EncodeToString(foo[:])
