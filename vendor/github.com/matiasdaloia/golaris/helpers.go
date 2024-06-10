package golaris

import "os"

func (g *Golaris) CreateDirIfNotExists(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, 0750)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Golaris) CreateFileIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}

		defer func(f *os.File) {
			_ = f.Close()
		}(file)
	}

	return nil
}
