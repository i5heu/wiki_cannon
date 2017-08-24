package main

import "time"

func refreshCache() {
	if TMPCACHEWRITE == false && TMPCACHECACHEWRITE == false {
		TMPCACHEWRITE = true
		time.Sleep(500 * time.Millisecond)
		Geldlogfunc("geldlog-false")
		Geldlogfunc("geldlog-true")
		cache(false, "article-false")
		cache(true, "article-true")
		Namespacefunc("namespace-true")
		Namespacefunc("namespace-false")
		Eventlogfunc("event-false")
		Eventlogfunc("event-true")
		Projectfunc("project-false")
		Projectfunc("project-true")
		Lasteditfunc("lastedit-false")
		Lasteditfunc("lastedit-true")
		TMPCACHEWRITE = false
		time.Sleep(500 * time.Millisecond)

		TMPCACHECACHEWRITE = true
		time.Sleep(500 * time.Millisecond)

		for key, value := range TMPCACHE {
			TMPCACHECACHE[key] = value
		}

		TMPCACHECACHEWRITE = false

	}
}
