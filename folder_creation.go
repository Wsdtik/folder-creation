package folder_creation

import (
	"github.com/wsdtik/folder-creation/creator"
	"github.com/wsdtik/folder-creation/looker"
)

func Finder(dir string) (*looker.DirLook, error) {
	cd := &creator.DirCreator{}
	return cd.Create("dir")
}
