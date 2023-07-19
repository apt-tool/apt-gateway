package cmd

import "github.com/spf13/cobra"

type Runner struct{}

func (r Runner) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "runner",
		Short: "start apt runner",
		Run: func(_ *cobra.Command, _ []string) {
			r.main()
		},
	}
}

func (r Runner) main() {

}
