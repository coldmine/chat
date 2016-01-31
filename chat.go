package main

// Chat has several channels.
type Chat struct {
	Address  string
	Channels []*Channel
}

// Channel is a place to talk each other by participants.
type Channel struct {
	Participants []*User
	Messages     []Message
}

// User should have id and also possibly have mail.
type User struct {
	Id   string
	Mail string
}

// Message is either a message string or an attachment.
// If attachment is not nil, then it's an attachment,
// or it's a message string.
type Message struct {
	Sender     *User
	Body       string
	Attachment *Attachment
}

// Attachment is something like video, image, or file.
type Attachment struct {
	Typ      FileType
	Location string
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
