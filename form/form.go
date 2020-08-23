package form

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/xxnmxx/cli"
)

// validFunctions
type validFunc func(string) bool

func validFloat(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return false
	}
	return true
}

func validString(input string) bool {
	return true
}

type Form struct {
	s          *bufio.Scanner
	b          *strings.Builder
	out        io.Writer
	prompt     string
	lists      map[string][]string
	recs       map[string][]string
	validFuncs map[string]validFunc
}

func (f *Form) regValidFunc(typ string, vf validFunc) {
	f.validFuncs[typ] = vf
}

func NewForm(in io.Reader, out io.Writer, p string) *Form {
	f := Form{
		s:      bufio.NewScanner(in),
		b:      new(strings.Builder),
		out:    out,
		prompt: p,
		lists:  make(map[string][]string),
		recs:   make(map[string][]string),
	}
	f.validFuncs = make(map[string]validFunc)
	f.regValidFunc("float", validFloat)
	f.regValidFunc("string", validString)
	return &f
}

func (f *Form) CreateList(name string, list ...string) error {
	if _, ok := f.lists[name]; !ok {
		f.lists[name] = []string{} // Initialize
		f.recs[name] = []string{}  // ditto!
		f.lists[name] = append(f.lists[name], list...)
		return nil
	} else {
		return fmt.Errorf("error: list already exists.\n")
	}
}

func (f *Form) BuildList(name string) string {
	var b strings.Builder
	for i, v := range f.lists[name] {
		if i != len(f.lists[name])-1 {
			b.WriteString(v + "\t")
		} else {
			b.WriteString(v + "\n")
		}
	}
	return b.String()
}

func (f *Form) buildList(name string) {
	for i, v := range f.lists[name] {
		if i != len(f.lists[name])-1 {
			f.b.WriteString(v + "\t")
		} else {
			f.b.WriteString(v + "\n")
		}
	}
}

func (f *Form) buildRec(name string) {
	for _, rec := range f.recs[name] {
		f.b.WriteString(rec + "\t")
	}
}

// Input implement float parser only now.
func (f *Form) Input(name string, typ string) {
	for i := 0; i < len(f.lists[name]); i++ {
		f.b.Reset()
		f.buildList(name)
		f.buildRec(name)
		cli.ClearScreen()
		f.b.WriteString(f.prompt)
		f.Tabler()
		f.s.Scan()
		// Insert parser here!
		// Float only now.
		if !f.validFuncs[typ](f.s.Text()) {
			//log.Fatalf("parse error: want %v",typ)
			fmt.Printf("parse error: want %v", typ)
			f.s.Scan()
			i--
			continue
		}
		f.recs[name] = append(f.recs[name], f.s.Text())
	}
	cli.ClearScreen()
	f.b.Reset()
	f.buildList(name)
	f.buildRec(name)
	f.Tabler()
}

func (f *Form) Tabler() {
	w := tabwriter.NewWriter(f.out, 0, 4, 1, 0, 0)
	fmt.Fprint(w, f.b.String())
	w.Flush()
}
