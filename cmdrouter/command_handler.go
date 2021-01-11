package cmdrouter

// CommandHandler is an interface for command execution
type CommandHandler interface {
	Handle(ctx *Context)
}
