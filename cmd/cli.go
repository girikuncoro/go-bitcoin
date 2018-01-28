package main

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

type Cli struct {
	Cmd  *cobra.Command // current command
	Args []string       // arguments for current command

	stdIn    io.Reader
	rootCmd  *cobra.Command // the command hierarchy
	priceCmd *priceCmd

	printer *Printer
}

const (
	// CliProgram cli program name
	CliProgram = "coin"
)

// Load environment configuration and register sub commands
func NewCli() *Cli {
	cli := &Cli{
		stdIn: os.Stdin,
	}
	cli.rootCmd = &cobra.Command{
		Use:   CliProgram,
		Short: "Geka Coin CLI",
		RunE:  cli.usageRunner(),
	}
	cli.printer = NewPrinter(os.Stderr)
	cli.SetOutput(os.Stdout, os.Stderr)

	registerPriceCmds(cli)
	return cli
}

func (cli *Cli) SetOutput(stdout, stderr io.Writer) *Cli {
	if stdout != nil {
		cli.rootCmd.SetOutput(stdout)
	}
	if stderr != nil {
		cli.rootCmd.SetOutput(stderr)
	}
	return cli
}

// print usage of current command
func (cli *Cli) usage() error {
	if cli.Cmd != nil {
		return cli.Cmd.Usage()
	}
	return cli.rootCmd.Usage()
}

// runner execute current command with args
func (cli *Cli) runner(runner func(*Cli) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cli.rootCmd.SilenceUsage = true
		cli.Cmd = cmd
		cli.Args = args
		return runner(cli)
	}
}

// usageRunner calls usage on current command
func (cli *Cli) usageRunner() func(*cobra.Command, []string) error {
	return cli.runner(func(cli *Cli) error {
		return cli.usage()
	})
}

func (cli *Cli) Run() {
	if err := cli.rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
