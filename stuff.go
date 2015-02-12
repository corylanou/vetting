package bar

import (
	"fmt"

	"github.com/corylanou/vetting/foo"
)

type Stuff struct {
	Rows []*foo.Row
	Err  error
}

func DoIt() {
	row := &foo.Row{Name: "blah"}
	stuff := &Stuff{
		Rows: foo.Rows{row},
	}
	fmt.Printf("%+v", stuff)

}
