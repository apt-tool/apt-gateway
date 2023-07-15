package cmd

import "github.com/spf13/cobra"

type Migrate struct{}

func (m Migrate) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "execute database migrations",
		Run: func(_ *cobra.Command, _ []string) {
			m.main()
		},
	}
}

func (m Migrate) main() {

}
