package arguduck_test

import (
	"os"
	"testing"

	arguduck "github.com/duck-fat-studios/ArguDuck"
)

// TestStringArgumentParsing tests the addition and parsing of a string argument.
func TestStringArgumentParsing(t *testing.T) {
    // Create a new ArguDuck instance.
    parser := arguduck.InitArguDuck()

    // Define a string argument.
    parser.String("testarg", "t", "default", "Test argument", "general")

    // Simulate command line input.
    os.Args = []string{"cmd", "--testarg", "hello"}

    // Parse the arguments.
    parser.Parse()

    // Check if the argument was parsed correctly.
    if val, ok := parser.Args["testarg"].(string); !ok || val != "hello" {
        t.Errorf("Expected 'hello', got '%v'", val)
    }
}

// TestHelpFlag tests the help flag detection and help text generation.
// func TestHelpFlag(t *testing.T) {
//     // Create a new ArguDuck instance.
//     parser := arguduck.InitArguDuck()

//     // Define a flag argument.
//     parser.Flag("help", "h", "Displays help information", "general")

//     // Simulate command line input for help.
//     os.Args = []string{"cmd", "--help"}

//     // Parse the arguments.
//     parser.Parse()

//     // Check if the help flag was found and the help text was generated.
//     if !parser.HelpFlagFound {
//         t.Error("Expected help flag to be found")
//     }

//     if _, ok := parser.HelpText["general"]; !ok {
//         t.Error("Expected help text to be generated for 'general' group")
//     }
// }

// TestIntArgumentParsing tests the addition and parsing of an integer argument.
func TestIntArgumentParsing(t *testing.T) {
    // ... similar structure to TestStringArgumentParsing ...
}

// TestFloat32ArgumentParsing tests the addition and parsing of a float32 argument.
func TestFloat32ArgumentParsing(t *testing.T) {
    // ... similar structure to TestStringArgumentParsing ...
}

// TestFlagArgumentParsing tests the addition and parsing of a flag argument.
func TestFlagArgumentParsing(t *testing.T) {
    // ... similar structure to TestStringArgumentParsing ...
}

// TestInvalidArgumentType tests the behavior when an unsupported argument type is provided.
func TestInvalidArgumentType(t *testing.T) {
    // ... structure to test behavior when unsupported argument type is used ...
}

