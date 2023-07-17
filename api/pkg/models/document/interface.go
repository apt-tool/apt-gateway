package document

import "gorm.io/gorm"

type Interface interface {
	Create(document *Document) error
}

func GetInterface(db *gorm.DB) Interface {
	return &core{
		db: db,
	}
}

type core struct {
	db *gorm.DB
}

func (c *core) Create(document *Document) error {
	if err := c.db.Create(document).Error; err != nil {
		return err
	}

	for _, item := range document.Instruction {
		tmp := &DocumentInstructions{
			DocumentID:    document.ID,
			InstructionID: item.ID,
		}

		if err := c.db.Create(tmp).Error; err != nil {
			return err
		}
	}

	return nil
}
