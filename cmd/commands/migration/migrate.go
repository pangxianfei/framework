package migration

import (
	"github.com/pangxianfei/framework/cmd"
	"github.com/pangxianfei/framework/database/migration"
)

func init() {
	cmd.Add(&Migrate{})
}

type Migrate struct {
}

func (mi *Migrate) Command() string {
	return "migrate"
}

func (mi *Migrate) Description() string {
	return "complete a task on the list"
}

func (mi *Migrate) Handler(arg *cmd.Arg) error {
	m := &migration.MigrationUtils{}
	m.Migrate()

	return nil
}
