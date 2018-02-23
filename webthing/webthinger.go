package webthing

type WebThinger interface {
	toWebThing() (WebThing, error)
}
