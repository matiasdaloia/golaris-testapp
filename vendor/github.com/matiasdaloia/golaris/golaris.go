package golaris

import "fmt"

type Golaris struct {
	AppName string
	Debug   bool
	Version string
}

func (g *Golaris) New(rootPath string) error {
	inithPaths := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}
	err := g.Init(inithPaths)
	if err != nil {
		return err
	}

	err = g.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	return nil
}

func (g *Golaris) Init(p initPaths) error {
	for _, path := range p.folderNames {
		err := g.CreateDirIfNotExists(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Golaris) checkDotEnv(roothPath string) error {
	err := g.CreateFileIfNotExists(fmt.Sprintf("%s/.env", roothPath))
	if err != nil {
		return err
	}

	return nil
}
