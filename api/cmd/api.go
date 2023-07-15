package cmd

import "github.com/spf13/cobra"

type API struct{}

func (a API) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "build apt api",
		Run: func(_ *cobra.Command, _ []string) {
			a.main()
		},
	}
}

func (a API) main() {

}
