package commands

const (
	StartMessage = `Welcome to the Incognito Bot, inspired by @puanonymous.
Since many students enjoyed PU Anonymous, this bot was created to continue the legacy. It still lacks some features, but it gets the job done - sending your messages immediately to the channel while keeping your identity hidden.
Use /help command to get all available commands, use /issues commands to get list of latest issues.
Improvements and updates are on the horizon, and your feedback is the key for that! 🚀`

	HelpMessage = `Here are some commands you can use:
/start - Start interacting with the bot
/feedback - Give feedback to the developer
/issues - List of noticeable known issues of the bot 
/help - Get help information`

	FeedbackMessage = `Please send your feedback now. Your next message will be sent as feedback.`

	KnownIssues = `Currently are some problems:
1. Any Text Makeups, such as spoilers, underlines, bold and italic, does not supported by bot currently.
2. Replies to messages are not supported yet.

If you found any other issues, please use /feedback command to send you issues/suggestions.
Thank you {^-^}`

	UnknownCommandMessage = `I didn't recognize that command. Try /help for a list of commands.`
)
