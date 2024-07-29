package creator

import (
	
	"os"
	"github.com/wsdtik/folder-creation/existing"
	"github.com/wsdtik/folder-creation/looker"
)

type DirCreator struct {
}

func (dc *DirCreator) Create(dir string) (*looker.DirLook, error) {

	if existing.ExistedDir(dir) {
		foundDir, err := existing.CorrectDir(dir)
		if err != nil {
			return nil, err
		}
		err = os.Chdir(foundDir)
		if err != nil {
			return nil, err
		}
		return &looker.DirLook{DirPath: foundDir}, nil
	}


	


	return &looker.DirLook{}, nil
}
