package migration

import (
	"github.com/pangxianfei/framework/cmd"
	"github.com/pangxianfei/framework/database/migration"
)

func init() {
	cmd.Add(&MigrationRollback{})
}

type MigrationRollback struct {
}

func (mr *MigrationRollback) Command() string {
	return "migration:rollback"
}

func (mr *MigrationRollback) Description() string {
	return "complete a task on the list"
}

func (mr *MigrationRollback) Handler(arg *cmd.Arg) error {
	m := &migration.MigrationUtils{}
	m.Rollback()
	return nil
}
