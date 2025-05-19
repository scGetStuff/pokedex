package commands

// I removed name; it is the same as the key
// it's not like there is a display name vs command name
type CliCommand struct {
	Description string
	Callback    func() error
}

// TOOD: not doing it
// I hate the idea of adding config param to callback function
// it is only used by the map command
// no reason to change signature for all becauses of 1
type Config struct {
	NextURL     string
	PreviousURL string
}
