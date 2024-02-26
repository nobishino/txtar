package txtar

import (
	"fmt"
	"io"
	"os"

	xtxtar "golang.org/x/tools/txtar"
)

func Write(paths []string, w io.Writer) error {
	files := make([]xtxtar.File, 0, len(paths))
	for _, p := range paths {
		if fi, err := os.Stat(p); err != nil {
			return err
		} else if fi.IsDir() {
			continue
		}
		d, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		a := xtxtar.File{
			Name: p,
			Data: d,
		}
		files = append(files, a)
	}
	arc := xtxtar.Archive{
		Files: files,
	}
	if _, err := fmt.Fprint(w, string(xtxtar.Format(&arc))); err != nil {
		return err
	}
	return nil
}
