// Package arguduck provides a simple and flexible command-line argument parser.
package arguduck

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ArguDuck struct holds all the necessary information for command-line argument parsing.
// It includes registered arguments, their values, help text, and other parsing state data.
type ArguDuck struct {
	aboutMessage    string
	Args            map[string]interface{}
	arguments       []arguDuckArgumentInterface
	helpText        map[string][]string
	helpFlagFound   bool
	parsed          bool
	programArgs     []string
	shortToFullName map[string]string
}

// InitArguDuck initializes and returns a new instance of ArguDuck.
// This function should be used to create an ArguDuck object before adding arguments and parsing.
func InitArguDuck() *ArguDuck {
	return &ArguDuck{
		aboutMessage:    "",
		Args:            make(map[string]interface{}),
		helpText:        make(map[string][]string),
		helpFlagFound:   false,
		shortToFullName: make(map[string]string),
		parsed:          false,
		programArgs:     nil,
	}
}

// Parse analyzes the program's command line arguments based on the defined flags and parameters.
// It should be called after all arguments have been defined through the ArguDuck instance.
func (a *ArguDuck) Parse() {

	if a.parsed {
		return
	}

	a.programArgs = os.Args[1:]

	// If a help parameter doesnt exists we'll make it.
	_, ok := a.getFullNameFromShort("h")

	if ok {
		a.Flag("help", "h", "Displays this help text")
	}

	for _, arg := range a.programArgs {
		if arg == "--help" || arg == "-h" {
			a.helpFlagFound = true
			break
		}
	}

	if a.helpFlagFound {
		a.printHelpText()
		os.Exit(0)
	}

	for index, item := range a.programArgs {

		// Disregard anythign without a - or --
		if !strings.HasPrefix(item, "-") {
			continue
		}

		// Handle Full names
		if strings.HasPrefix(item, "--") {
			argName := strings.TrimPrefix(item, "--")
			a.setArgValue(argName, index+1)
			continue
		}

		// Handle Short Names
		if strings.HasPrefix(item, "-") && len(item) == 2 {
			shortName := strings.TrimPrefix(item, "-")
			argName, _ := a.getFullNameFromShort(shortName)
			a.setArgValue(argName, index+1)
			continue
		}

		// Handle Flags
		if strings.HasPrefix(item, "-") && len(item) > 2 {
			shortNames := strings.TrimPrefix(item, "-")

			for _, shortName := range shortNames {
				argName, _ := a.getFullNameFromShort(string(shortName))
				a.setArgValue(argName, index+1)
			}
			continue
		}
	}

	a.parsed = true
}

// Flag defines a new boolean flag argument.
// 'name' is the full name of the flag, 'short' is the abbreviated form, and 'help' describes the flag's purpose.
func (a *ArguDuck) Flag(name string, short string, help string, group ...string) (error, ArguDuckErrorString) {
	err, errorString := a.addArgument(name, short, help, false, a.determineGroup(group))

	if err != nil {
		return err, errorString
	}

	return nil, OK
}

// Float32 defines a new float32 argument.
// 'name' is the argument's name, 'short' is the abbreviated form, 'defaultValue' specifies the default value, and 'help' describes the argument.
func (a *ArguDuck) Float(name string, short string, defaultValue float64, help string, group ...string) (error, ArguDuckErrorString) {
	err, errorString := a.addArgument(name, short, help, defaultValue, a.determineGroup(group))

	if err != nil {
		return err, errorString
	}

	return nil, OK
}

// Int defines a new integer argument.
// Similar to Float32, it allows the specification of a name, short form, default value, and help description.

func (a *ArguDuck) Int(name string, short string, defaultValue int, help string, group ...string) (error, ArguDuckErrorString) {
	err, errorString := a.addArgument(name, short, help, defaultValue, a.determineGroup(group))
	
	if err != nil {
		return err, errorString
	}

	return nil, OK
}

// String defines a new string argument.
// This method is used for arguments that are expected to be strings, with similar parameters as in Int and Float32.
func (a *ArguDuck) String(name string, short string, defaultValue string, help string, group ...string) (error, ArguDuckErrorString) {
	err, errorString := a.addArgument(name, short, help, defaultValue, a.determineGroup(group))
	
	if err != nil {
		return err, errorString
	}

	return nil, OK
}

func (a *ArguDuck) determineGroup(group []string) string {
	if len(group) > 0 {
		return group[0]
	}
	return "Usage"
}

func (a *ArguDuck) GetAbout() string {
	return a.aboutMessage
}

func (a *ArguDuck) SetAbout(message string) {
	a.aboutMessage = message
}

// addArgument is an internal method that adds a new argument to the ArguDuck instance.
// It is called by the public methods like String, Int, Float32, and Flag.
func (a *ArguDuck) addArgument(name string, short string, help string, defaultValue interface{}, group ...string) (error, ArguDuckErrorString) {
	var arg arguDuckArgumentInterface
	var groupName string

	if _, ok := a.Args[name]; ok {
		return fmt.Errorf("Argument %s already exists", name), ARG_IN_USE
	}

	if _, ok := a.shortToFullName[short]; ok {
		shrt, _ := a.getFullNameFromShort(short)
		return fmt.Errorf("Short %s already in use in %s", short, shrt), SHORT_IN_USE
	}

	if group != nil {
		groupName = group[0]
	}

	switch v := defaultValue.(type) {
	case bool:
		arg = &flagArgument{arguDuckArgument: arguDuckArgument{name: name, short: short, help: help, group: groupName}}
	case float64:
		arg = &floatArgument{arguDuckArgument: arguDuckArgument{name: name, short: short, help: help, group: groupName}, defaultValue: v}
	case int:
		arg = &intArgument{arguDuckArgument: arguDuckArgument{name: name, short: short, help: help, group: groupName}, defaultValue: v}
	case string:
		arg = &stringArgument{arguDuckArgument: arguDuckArgument{name: name, short: short, help: help, group: groupName}, defaultValue: v}
	default:
		return fmt.Errorf("Unknown type %s", v), UNKNOWN_TYPE
	}

	a.Args[name] = defaultValue
	a.arguments = append(a.arguments, arg)

	a.shortToFullName[arg.getShort()] = arg.getName()

	if groupName != "" {
		a.addHelpText(groupName, arg)
	}

	return nil, OK
}

// getFullNameFromShort is an internal method that retrieves the full argument name from its short form.
// It is used during the parsing process.
func (a *ArguDuck) getFullNameFromShort(shortName string) (string, bool) {
	fullName, ok := a.shortToFullName[shortName]
	return fullName, ok
}

// addHelpText adds help text for a specific argument to the ArguDuck instance.
// It organizes help text by group for easier display and management.
func (a *ArguDuck) addHelpText(groupName string, argument arguDuckArgumentInterface) {
	if groupName == "" {
		groupName = "Usage" // Default group name
	}
	helpText := a.setHelpText(argument)
	a.helpText[groupName] = append(a.helpText[groupName], helpText)
}

// printHelpText displays the help text for all arguments.
// It is typically called when the help flag is provided in the command line.
func (a *ArguDuck) printHelpText() {

	if a.GetAbout() != "" {
		fmt.Printf("%s\n\n", a.GetAbout())
	}

	// Print the basic usage first
	fmt.Println("Usage")
	for _, help := range a.helpText["Usage"] {
		fmt.Println(help)
	}
	fmt.Println("")

	for group, helptext := range a.helpText {
		if group != "Usage" {

			fmt.Println(group)
			for _, help := range helptext {
				fmt.Println(help)
			}
			fmt.Println("")
		}
	}
}

// setHelpText generates the help text string for a single argument.
// This method formats the help text for display based on the argument's properties.
func (a *ArguDuck) setHelpText(argInterface arguDuckArgumentInterface) string {
	var helpTextBuilder strings.Builder

	shortName := ""
	if argInterface.getShort() != "" {
		shortName = fmt.Sprintf("-%v", argInterface.getShort())
	}
	helpTextBuilder.WriteString(fmt.Sprintf("    %-25s %-4s %s", argInterface.getName(), shortName, argInterface.getHelp()))

	return helpTextBuilder.String()
}

// setArgValue sets the value of an argument based on its type and the passed command-line value.
// This method is called internally during parsing.
func (a *ArguDuck) setArgValue(argName string, argIndex int) error {
	argValue := a.Args[argName]

	switch argValue.(type) {
	case string:
		a.Args[argName] = a.programArgs[argIndex]
	case int:
		intValue, err := strconv.Atoi(a.programArgs[argIndex])
		if err != nil {
			return fmt.Errorf("invalid integer value for argument %v: %w", argName, err)
		}
		a.Args[argName] = intValue
	case bool:
		a.Args[argName] = true
	default:
		panic(fmt.Sprintf("Unsupported argument type for argument %v with value:%v\n", argName, argValue))
	}

	return nil
}
