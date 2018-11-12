package alibuild

import (
	"fmt"
	"strings"

	"github.com/brinick/shell"
)

// New constructs an aliBuild type
func New(exe, packageName string, env []string) *AliBuild {
	return &AliBuild{
		executable:  exe,
		env:         env,
		packageName: packageName,
	}
}

// AliBuild is the type representing the ALICE build tool
type AliBuild struct {
	executable  string
	env         []string
	packageName string
}

// Which returns the path to the alibuild executable being used
func (ab *AliBuild) Which() string {
	return ab.executable
}

// Package returns the name of the package upon
// which aliBuild commands will act
func (ab *AliBuild) Package() string {
	return ab.packageName
}

// SetPackageName allows one to set the package to run on
func (ab *AliBuild) SetPackageName(p string) {
	ab.packageName = p
}

// DefaultEnv returns the default environement that is
// applied for each aliBuild command. These can be overridden
// or appended to using a shell.Env option when calling a
// particular command.
func (ab *AliBuild) DefaultEnv() []string {
	return ab.env
}

// Build maps to the alibuild build command
func (ab AliBuild) Build(args string, opts ...shell.Option) *shell.Result {
	cmd := fmt.Sprintf(
		"%s build %s %s",
		ab.executable,
		args,
		ab.packageName,
	)

	opts = append(opts, shell.Env(ab.env))
	result := shell.Run(cmd, opts...)
	return result
}

// Doctor maps to the alibuild doctor command
func (ab AliBuild) Doctor(args string, opts ...shell.Option) *shell.Result {
	cmd := fmt.Sprintf(
		"%s doctor %s %s",
		ab.executable,
		args,
		ab.packageName,
	)

	opts = append(opts, shell.Env(ab.env))
	result := shell.Run(cmd, opts...)
	return result
}

// Clean maps to the alibuild clean command
func (ab AliBuild) Clean(debug bool) *shell.Result {
	debugOption := "--debug"
	if !debug {
		debugOption = ""
	}

	cmd := fmt.Sprintf(
		"%s clean %s",
		ab.executable,
		debugOption,
	)

	return shell.Run(cmd)
}

func (ab AliBuild) Version() (string, error) {
	cmd := fmt.Sprintf("%s version", ab.executable)

	if res := shell.Run(cmd); res != nil {
		return "", res.Error
	} else {
		return res.Stdout.Text(true), nil
	}
}

// CommandHelp maps to the alibuild help for the given subcommand
func (ab AliBuild) CommandHelp(name string) *shell.Result {
	cmd := fmt.Sprintf(
		"%s %s --help",
		ab.executable,
		name,
	)

	return shell.Run(cmd)
}

// Help maps to the alibuild top level help command
func (ab AliBuild) Help() *shell.Result {
	cmd := fmt.Sprintf(
		"%s --help",
		ab.executable,
	)

	return shell.Run(cmd)
}

// HasFetchReposOption checks for the --fetch-repos option in
// the output of the aliBuild build --help command
func (ab *AliBuild) HasFetchReposOption() bool {
	res := ab.CommandHelp("build")
	if res.IsError() {

	}

	return strings.Contains(res.Stdout.Text(false), "fetch-repos")
}
