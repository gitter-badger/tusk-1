package ui

const (
	commandActionString = "Running"
	skippedString       = "Skipping"

	outputPrefix = "=> "
)

// PrintCommand prints the command to be executed.
func PrintCommand(command string) {
	if Verbosity <= VerbosityLevelQuiet {
		return
	}

	printf(
		LoggerStderr,
		"[%s] %s\n",
		blue(commandActionString),
		bold(command),
	)
}

// PrintSkipped prints the command skipped and the reason.
func PrintSkipped(command string, reason string) {
	if Verbosity < VerbosityLevelVerbose {
		return
	}

	printf(
		LoggerStderr,
		"[%s] %s\n%s%s\n",
		yellow(skippedString),
		bold(command),
		cyan(outputPrefix),
		reason,
	)
}

// PrintCommandError prints an error from a running command.
func PrintCommandError(err error) {
	if Verbosity <= VerbosityLevelQuiet {
		return
	}

	printf(
		LoggerStderr,
		"%s%s\n",
		red(outputPrefix),
		err.Error(),
	)
}
