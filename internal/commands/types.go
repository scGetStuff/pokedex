package commands

// I removed name; it is the same as the key
// it's not like there is a display name vs command name
type CliCommand struct {
	Description string
	Callback    func(cmdArgs []string) error
}

// not doing it
// I hate the idea of adding a parameter for one command to all the callback functions
// I did add an args array to handle parameters to commands
// type Config struct {
// 	NextURL     string
// 	PreviousURL string
// }
