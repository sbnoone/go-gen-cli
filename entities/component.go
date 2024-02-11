package entities

import (
	"fmt"
	"gen/lib"
	"path/filepath"

	"github.com/iancoleman/strcase"
)

type Component struct {
	FileName string
	Config   bool
	Types    bool
}

func (c *Component) Create() {
	componentName := strcase.ToCamel(lib.FileNameWithoutExtension(c.FileName))

	content := []byte(fmt.Sprintf(
		`interface Props {}

export const %s = (p: Props) => {
  return null
}
`, componentName))
	lib.CreateFile(c.FileName, ".tsx", content)

	if c.Config {
		fileName := fmt.Sprintf("%s.config%s", lib.FileNameWithoutExtension(c.FileName), filepath.Ext(c.FileName))
		lib.CreateFile(fileName, ".ts", []byte{})
	}

	if c.Types {
		fileName := fmt.Sprintf("%s.types%s", lib.FileNameWithoutExtension(c.FileName), filepath.Ext(c.FileName))
		lib.CreateFile(fileName, ".ts", []byte{})
	}
}
