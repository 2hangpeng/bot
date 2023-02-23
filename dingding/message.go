package dingding

type Message interface {
	ToBytes() ([]byte, error)
}

type Type string

const (
	// text text
	text Type = "text"
	// markdown markdown
	markdown Type = "markdown"
	// link link
	link Type = "link"
	// actionCard actionCard
	actionCard Type = "actionCard"
	// feedCard feedCard
	feedCard Type = "feedCard"
)
