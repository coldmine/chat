package main

// Chat has several channels.
type Chat struct {
	address  string
	channels []*Channel
}

// Channel is a place to talk each other by participants.
type Channel struct {
	participants []*User
	messages     []Message
}

// User should have id and also possibly have mail.
type User struct {
	id   string
	mail string
}

// Message is either a message string or an attachment.
// If attachment is not nil, then it's an attachment,
// or it's a message string.
type Message struct {
	sender     *User
	body       string
	attachment *Attachment
}

// Attachment is something like video, image, or file.
type Attachment struct {
	typ      FileType
	location string
}

type FileType int

const (
	UnknownFile FileType = iota
	GeneralFile
	ImageFile
	VideoFile
)

func main() {
	panic("unimplemented")
}
