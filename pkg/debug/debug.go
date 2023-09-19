package debug

import (
	"github.com/Code-Hex/dd/p"
	"github.com/darki73/goflaresync/pkg/log"
	"os"
)

// Dump dumps the arguments.
func Dump(args ...interface{}) {
	if _, err := p.P(args...); err != nil {
		log.Fatalf("failed to dump: %s", err.Error())
	}
}

// DieAndDump dumps the arguments and exits.
func DieAndDump(args ...interface{}) {
	Dump(args...)
	os.Exit(0)
}
