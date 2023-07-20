package cmd

import (
	"fmt"
	"log"

	"github.com/automated-pen-testing/api/pkg/models/document"
	"github.com/automated-pen-testing/api/pkg/models/instruction"
	"github.com/automated-pen-testing/api/pkg/models/namespace"
	"github.com/automated-pen-testing/api/pkg/models/project"
	"github.com/automated-pen-testing/api/pkg/models/user"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type Migrate struct {
	Db *gorm.DB
}

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
	models := []interface{}{
		&document.Document{},
		&instruction.Instruction{},
		&namespace.Namespace{},
		&project.Project{},
		&user.User{},
	}

	for _, item := range models {
		if err := m.Db.AutoMigrate(item); err != nil {
			log.Println(fmt.Errorf("failed to migrate model error=%w", err))
		}
	}
}
