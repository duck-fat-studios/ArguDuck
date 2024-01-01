package arguduck

// arguDuckArgumentInterface is an interface that all argument types must implement.
// It provides methods to retrieve argument properties.
type arguDuckArgumentInterface interface {
	getName() string
	getShort() string
	getGroup() string
	getHelp() string
}

// arguDuckArgument is a base struct for command-line arguments, providing common properties.
type arguDuckArgument struct {
	name  string
	short string
	help  string
	group string
}

// StringArgument represents a command-line argument with a string value.
type StringArgument struct {
	arguDuckArgument
	defaultValue string
}

// getName returns the full name of the string argument.
func (s *StringArgument) getName() string {
	return s.name
}

// getShort returns the abbreviated name of the string argument.
func (s *StringArgument) getShort() string {
	return s.short
}

// getGroup returns the group name of the string argument.
func (s *StringArgument) getGroup() string {
	return s.group
}

func (s *StringArgument) getHelp() string {
	return s.help
}

type IntArgument struct {
	arguDuckArgument
	defaultValue int
}

// getName returns the full name of the int argument.
func (i *IntArgument) getName() string {
	return i.name
}

// getShort returns the abbreviated name of the int argument.
func (i *IntArgument) getShort() string {
	return i.short
}

// getGroup returns the group name of the int argument.
func (i *IntArgument) getGroup() string {
	return i.group
}

// getHelp returns the help text of the int argument.
func (i *IntArgument) getHelp() string {
	return i.help
}

type FlagArgument struct {
	arguDuckArgument
}

// getName returns the full name of the flag argument.
func (f *FlagArgument) getName() string {
	return f.name
}

// getShort returns the abbreviated name of the flag argument.
func (f *FlagArgument) getShort() string {
	return f.short
}

// getGroup returns the group name of the flag argument.
func (f *FlagArgument) getGroup() string {
	return f.group
}

// getHelp returns the help text of the flag argument.
func (f *FlagArgument) getHelp() string {
	return f.help
}

type Float32Argument struct {
	arguDuckArgument
	defaultValue float32
}

// getName returns the full name of the float32 argument.
func (f *Float32Argument) getName() string {
	return f.name
}

// getShort returns the abbreviated name of the float32 argument.
func (f *Float32Argument) getShort() string {
	return f.short
}

// getGroup returns the group name of the float32 argument.
func (f *Float32Argument) getGroup() string {
	return f.group
}

// getHelp returns the help text of the float32 argument.
func (f *Float32Argument) getHelp() string {
	return f.help
}
