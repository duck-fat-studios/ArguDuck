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
	helpText        map[string][]string
	programArgs     []string
	Args            map[string]interface{}
	arguments       []arguDuckArgumentInterface
	helpFlagFound   flag
	shortToFullName map[string]string
	parsed          flag
}

// InitArguDuck initializes and returns a new instance of ArguDuck.
// This function should be used to create an ArguDuck object before adding arguments and parsing.
func InitArguDuck() *ArguDuck {
	return &ArguDuck{
		programArgs:     nil,
		Args:            make(map[string]interface{}),
		helpText:        make(map[string][]string),
		shortToFullName: make(map[string]string),
		parsed:          false,
	}
}

// Parse analyzes the program's command line arguments based on the defined flags and parameters.
// It should be called after all arguments have been defined through the ArguDuck instance.
func (a *ArguDuck) Parse() {

	if a.parsed {
		return
	}

	a.programArgs = os.Args[1:]

	a.Flag("help", "h", "Displays this help text")

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

func (a *ArguDuck) Flag(name string, short string, help string, group ...string) {
	a.addArgument(name, short, help, false, group[0])
}

// Float32 defines a new float32 argument.
// 'name' is the argument's name, 'short' is the abbreviated form, 'defaultValue' specifies the default value, and 'help' describes the argument.
func (a *ArguDuck) Float32(name string, short string, defaultValue float32, help string, group ...string) {
	a.addArgument(name, short, help, defaultValue, group[0])
}

// Int defines a new integer argument.
// Similar to Float32, it allows the specification of a name, short form, default value, and help description.

func (a *ArguDuck) Int(name string, short string, defaultValue int, help string, group ...string) {
	a.addArgument(name, short, help, defaultValue, group[0])
}

// String defines a new string argument.
// This method is used for arguments that are expected to be strings, with similar parameters as in Int and Float32.
func (a *ArguDuck) String(name string, short string, defaultValue string, help string, group ...string) {
	a.addArgument(name, short, help, defaultValue, group[0])
}

// addArgument is an internal method that adds a new argument to the ArguDuck instance.
// It is called by the public methods like String, Int, Float32, and Flag.
func (a *ArguDuck) addArgument(name string, short string, help string, defaultValue interface{}, group string) {
	var arg arguDuckArgumentInterface

	switch v := defaultValue.(type) {
	case flag: 
		arg = &FlagArgument{arguDuckArgument: arguDuckArgument{name: name, short: short, help: help, group: group}}
	case float32:
		arg = &Float32Argument{arguDuckArgument: arguDuckArgument{name: name, short: short, help: help, group: group}, defaultValue: v}
	case int:
		arg = &IntArgument{arguDuckArgument: arguDuckArgument{name: name, short: short, help: help, group: group}, defaultValue: v}
	case string:
		arg = &StringArgument{arguDuckArgument: arguDuckArgument{name: name, short: short, help: help, group: group}, defaultValue: v}
	default:
		fmt.Errorf("Unsupported argument type for argument %v with value %v\n", name, v)
	}

	a.Args[name] = defaultValue
	a.arguments = append(a.arguments, arg)
	a.shortToFullName[arg.getShort()] = arg.getName()

	if group != "" {
		a.addHelpText(group, arg)
	}
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

	// Print the basic ussage first
	fmt.Println("Ussage")
	for _, help := range a.helpText["Ussage"] {
		fmt.Println(help)
	}
	fmt.Println("")

	for group, helptext := range a.helpText {
		if group != "Ussage" {

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

	fmt.Printf("ArgName: %v\n", argName)

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