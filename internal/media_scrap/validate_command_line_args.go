package mediascrap

import (
	"flag"
	"fmt"
)

func validateArgs(args []string, fs *flag.FlagSet) error {
	// argumentSet accumulates the command line arguments received from the command line. If the user didn't sent it will not be here
	argumentsSet := make(map[string]bool)

	fs.Visit(func(f *flag.Flag) { argumentsSet[f.Name] = true })
	for _, arg := range args {
		if !argumentsSet[arg] {
			return fmt.Errorf("missing required -%s argument", arg)
		}
	}

	return nil
}
