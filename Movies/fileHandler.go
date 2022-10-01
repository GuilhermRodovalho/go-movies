package Movies

import (
	"encoding/json"
	"io/fs"
	"os"
)

type fileHandler struct{}

func NewFileHandler() Handler {
	return fileHandler{}
}

func (f fileHandler) Read() (Movies, error) {
	var mvs Movies
	data, err := os.ReadFile(MoviesFileName)
	if err != nil {
		// The file probably is empty
		return nil, err
	}

	json.Unmarshal(data, &mvs)

	return mvs, nil
}

func (f fileHandler) Write(mv *Movie) error {
	mvs, err := f.Read()

	mvs = append(mvs, *mv)
	file, err := os.OpenFile(MoviesFileName, os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	err = file.Chmod(fs.FileMode(0777))
	if err != nil {
		return err
	}

	jsn, err := json.Marshal(mvs)
	if err != nil {
		return err
	}
	_, err = file.Write(jsn)
	if err != nil {
		return err
	}

	return nil
}
