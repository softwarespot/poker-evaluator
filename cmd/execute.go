package cmd

import "flag"

func Execute(args []string) error {
	var showHelp bool
	flagBoolVarP(&showHelp, "help", "h", false, "Display the help text and exit")

	var showVersion bool
	flagBoolVarP(&showVersion, "version", "v", false, "Display the version of the application and exit")

	var asJSON bool
	flagBoolVarP(&asJSON, "json", "j", false, "Output the result as JSON")
	flag.Parse()

	if showHelp {
		cmdHelp()
		return nil
	}
	if showVersion {
		return cmdVersion(asJSON)
	}
	return cmdPlay(args, asJSON)
}
