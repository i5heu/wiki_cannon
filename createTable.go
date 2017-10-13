package main

import "fmt"

func CreateTable() {
	fmt.Println("Create DB Tables")
	_, err = db.Exec("CREATE TABLE `BUarticle` ( `BUID` integer NOT NULL, `BUtimestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `id` integer NOT NULL, `timec` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `timelastedit` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `needlogin` integer DEFAULT '1', `namespace` varchar(128) NOT NULL DEFAULT 'main', `title` varchar(128) NOT NULL DEFAULT 'NO TITLE', `text` longtext, `tags` text, `viewcounter` integer DEFAULT NULL, `editcounter` integer DEFAULT NULL );")

	checkErr(err)

	_, err = db.Exec("CREATE TABLE `BUitems` ( `BUItemID` integer NOT NULL, `ItemID` integer DEFAULT NULL, `timecreate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `timelastedit` timestamp DEFAULT NULL, `needlogin` integer DEFAULT '1', `APP` varchar(20) DEFAULT NULL, `viewcounter` integer DEFAULT NULL, `editcounter` integer DEFAULT NULL, `title1` varchar(128) DEFAULT NULL, `title2` varchar(128) NOT NULL, `text1` text, `text2` text, `tags1` varchar(256) DEFAULT NULL, `num1` integer DEFAULT NULL, `num2` integer DEFAULT NULL, `num3` integer DEFAULT NULL );")

	checkErr(err)

	_, err = db.Exec("CREATE TABLE `article` ( `id` integer NOT NULL, `timec` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `timelastedit` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `needlogin` integer DEFAULT '1', `namespace` varchar(128) NOT NULL DEFAULT 'main', `title` varchar(128) NOT NULL DEFAULT 'NO TITLE', `text` longtext, `tags` text, `viewcounter` integer DEFAULT NULL, `editcounter` integer DEFAULT NULL );")

	checkErr(err)

	_, err = db.Exec("CREATE TABLE `eventlog` ( `id` integer NOT NULL, `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `name` varchar(256) DEFAULT NULL, `changeAPP` varchar(128) DEFAULT NULL, `changeID` integer DEFAULT NULL );")

	checkErr(err)

	_, err = db.Exec("CREATE TABLE `items` ( `ItemID` integer NOT NULL, `timecreate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `timelastedit` timestamp DEFAULT NULL, `needlogin` integer DEFAULT '1', `APP` varchar(20) DEFAULT NULL, `viewcounter` integer DEFAULT NULL, `editcounter` integer DEFAULT NULL, `title1` varchar(128) DEFAULT NULL, `title2` varchar(128) NOT NULL, `text1` text, `text2` text, `tags1` varchar(256) DEFAULT NULL, `num1` integer DEFAULT NULL, `num2` integer DEFAULT NULL, `num3` integer DEFAULT NULL );")

	checkErr(err)

	fmt.Println("DB Tables Createt")
}
