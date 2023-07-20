package cmd

import (
	"fmt"
	"log"

	"github.com/automated-pen-testing/api/internal/config/migration"
	"github.com/automated-pen-testing/api/internal/utils/crypto"
	"github.com/automated-pen-testing/api/pkg/enum"
	"github.com/automated-pen-testing/api/pkg/models/document"
	"github.com/automated-pen-testing/api/pkg/models/instruction"
	"github.com/automated-pen-testing/api/pkg/models/namespace"
	"github.com/automated-pen-testing/api/pkg/models/project"
	"github.com/automated-pen-testing/api/pkg/models/user"

	"gorm.io/gorm"
)

// Migrate is the command of migration
type Migrate struct {
	Cfg migration.Config
	Db  *gorm.DB
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

	query := "INSERT INTO users (`username`, `password`, `role`) VALUES (?,?,?)"

	if err := m.Db.Exec(query, m.Cfg.Root, crypto.GetMD5Hash(m.Cfg.Pass), enum.RoleAdmin).Error; err != nil {
		log.Println(fmt.Errorf("failed to insert root user error=%w", err))
	}
}
