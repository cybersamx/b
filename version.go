package b

import (
	"fmt"

	"github.com/cybersamx/c"
)

func Version() string {
	return fmt.Sprintf("a: v1.1.0\n%s", c.Version())
}
