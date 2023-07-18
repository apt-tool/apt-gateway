package instruction

import "gorm.io/gorm"

// Interface manages the instruction methods
type Interface interface {
	Create(instruction *Instruction) error
	Delete(instructionID uint) error
	Get() ([]*Instruction, error)
}

func New(db *gorm.DB) Interface {
	return &core{
		db: db,
	}
}

type core struct {
	db *gorm.DB
}

func (c core) Create(instruction *Instruction) error {
	return c.db.Create(instruction).Error
}

func (c core) Delete(instructionID uint) error {
	return c.db.Delete(&Instruction{}, "id = ?", instructionID).Error
}

func (c core) Get() ([]*Instruction, error) {
	list := make([]*Instruction, 0)

	if err := c.db.Find(&list).Error; err != nil {
		return nil, ErrFailedToGet
	}

	return list, nil
}
