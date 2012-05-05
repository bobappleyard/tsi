package main

import (
	"os"
	"fmt"
	"flag"
	"github.com/bobappleyard/ts"
	_ "github.com/bobappleyard/ts/ext"
)

const version = "0.3"

var (
	ForcePrompt = flag.Bool("p", false, "force a prompt to appear")
	Compiling = flag.Bool("c", false, "compile rather than evaluate")
	Outfile = flag.String("o", "", "target to compile to")
	Strip = flag.Bool("s", false, "strip debugging info")
	ViewVersion = flag.Bool("version", false, "print BobScript version")
)

const msg = `usage: bsi [options] script args ...
Specifying no script opens a prompt. Takes multiple script files when -c is set.

Options:`

func usageMsg() {
	fmt.Fprintln(os.Stderr, msg)
	off := 0
	flag.VisitAll(func(f *flag.Flag) {
		if off < len(f.Name) {
			off = len(f.Name)
		}
	})
	flag.VisitAll(func(f *flag.Flag) {
		sep := make([]byte, off - len(f.Name))
		for i := range sep {
			sep[i] = 0x20
		}
		fmt.Fprintf(os.Stderr, "-%s %s-- %s\n", f.Name, sep, f.Usage)
	})
}

func fatal(err interface{}) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func newi(as []string) *ts.Interpreter {
	i := ts.New()
	args := i.Accessor("args")
	i.Import("system").Set(args, ts.Wrap(as))
	return i
}

func compile(as []string) {
	if len(as) == 0 {
		fatal("error: no files provided")
	}
	u := new(ts.Unit)
	for _, x := range as {
		f, err := os.Open(x)
		if err != nil {
			fatal(err)
		}
		u.Compile(f, x)
		f.Close()
	}
	targ := *Outfile
	if targ == "" {
		targ = as[0] + "c"
	}
	f, err := os.Create(targ)
	if err != nil {
		fatal(err)
	}
	u.Save(f)
	f.Close()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fatal(err)
		}
	}()/**/
	flag.Usage = usageMsg
	flag.Parse()
	as := flag.Args()
	
	if *Strip {
		ts.Strip = true
	}
	
	switch {
	case *ViewVersion:
		fmt.Fprintln(os.Stderr, version)
		os.Exit(2)
	case *Compiling:
		compile(as)
	case len(as) == 0:
		i := newi(as)
		i.Repl()
	default:
		i := newi(as)
		i.Load(as[0])
		if *ForcePrompt {
			i.Repl()
		}
	}
}

