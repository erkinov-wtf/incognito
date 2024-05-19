package messages

import (
	"math/rand"
	"time"
)

const (
	GotMessage = "Got your message!"
	Safe       = "🕵️‍♂️Your secret's safe with us. Your message has been dispatched into the digital void, anonymously."
	Mask       = "🎭 Behind the mask of anonymity, your message has been conveyed silently."
	Incognito  = "🌐 Your message is now traversing the net incognito. Mission accomplished!"
	Ghost      = "👻 Boo! Your message has ghosted its identity!"
	NotHeard   = "🙈 Seen but not heard. Your message is incognito!"
	Undercover = "🎩 Poof! Your message is now undercover."
	Flew       = "🌪️ Whoosh! Your message flew by anonymously!"
	Space      = "🚀 Off to space without a trace. Message sent!"
	Transmit   = "📡 Beep boop! Message transmitted to the unknown."
	NoIdCard   = "🧳 Packed and sent. No ID Card needed, I'm not security"
	SentRight  = "💫 Starlight, star bright, your message sent right!"
)

func GetRandomReply() string {
	replies := []string{
		GotMessage, Safe, Mask,
		Incognito, Ghost, NotHeard,
		Undercover, Flew, Space,
		Transmit, NoIdCard, SentRight,
	}

	// Seed the random number generator with the current time
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Generate a random index
	randomIndex := random.Intn(len(replies))

	// Return the replies at the random index
	return replies[randomIndex]
}
