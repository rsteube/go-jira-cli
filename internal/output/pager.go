package output

import "github.com/cli/cli/pkg/iostreams"

func Pager(f func(io *iostreams.IOStreams) error) error {
	io := iostreams.System()
	io.SetPager("bat --style header,grid")
	err := io.StartPager()
	if err != nil {
		return err
	}
	defer io.StopPager()
	return f(io)
}
