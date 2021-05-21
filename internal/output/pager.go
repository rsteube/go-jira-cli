package output

import "github.com/cli/cli/pkg/iostreams"

func Pager(f func(io *iostreams.IOStreams) error) error {
	return Output(func(io *iostreams.IOStreams, cs *iostreams.ColorScheme) error {
		io.SetPager("bat --style grid")
		if err := io.StartPager(); err != nil {
			return err
		}
		defer io.StopPager()
		return f(io)
	})
}

func Output(f func(io *iostreams.IOStreams, cs *iostreams.ColorScheme) error) error {
	io := iostreams.System()
	return f(io, io.ColorScheme())
}
