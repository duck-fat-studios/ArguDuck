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

// stringArgument represents a command-line argument with a string value.
type stringArgument struct {
	arguDuckArgument
	defaultValue string
}

// getName returns the full name of the string argument.
func (s *stringArgument) getName() string {
	return s.name
}

// getShort returns the abbreviated name of the string argument.
func (s *stringArgument) getShort() string {
	return s.short
}

// getGroup returns the group name of the string argument.
func (s *stringArgument) getGroup() string {
	return s.group
}

func (s *stringArgument) getHelp() string {
	return s.help
}

type intArgument struct {
	arguDuckArgument
	defaultValue int
}

// getName returns the full name of the int argument.
func (i *intArgument) getName() string {
	return i.name
}

// getShort returns the abbreviated name of the int argument.
func (i *intArgument) getShort() string {
	return i.short
}

// getGroup returns the group name of the int argument.
func (i *intArgument) getGroup() string {
	return i.group
}

// getHelp returns the help text of the int argument.
func (i *intArgument) getHelp() string {
	return i.help
}

type flagArgument struct {
	arguDuckArgument
}

// getName returns the full name of the flag argument.
func (f *flagArgument) getName() string {
	return f.name
}

// getShort returns the abbreviated name of the flag argument.
func (f *flagArgument) getShort() string {
	return f.short
}

// getGroup returns the group name of the flag argument.
func (f *flagArgument) getGroup() string {
	return f.group
}

// getHelp returns the help text of the flag argument.
func (f *flagArgument) getHelp() string {
	return f.help
}

type floatArgument struct {
	arguDuckArgument
	defaultValue float64
}

// getName returns the full name of the float32 argument.
func (f *floatArgument) getName() string {
	return f.name
}

// getShort returns the abbreviated name of the float32 argument.
func (f *floatArgument) getShort() string {
	return f.short
}

// getGroup returns the group name of the float32 argument.
func (f *floatArgument) getGroup() string {
	return f.group
}

// getHelp returns the help text of the float32 argument.
func (f *floatArgument) getHelp() string {
	return f.help
}
