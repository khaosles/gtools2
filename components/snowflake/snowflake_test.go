package snowflake

import (
	"fmt"
	"testing"

	"github.com/langwan/langgo/core"
)

func TestRun(t *testing.T) {
	core.EnvName = core.Development
	snowflake, _ := New(2)
	fmt.Printf("Int64  ID: %d\n", snowflake.machineID)
	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", snowflake.Gen())
	fmt.Printf("Int64  ID: %d\n", snowflake.Gen())
	fmt.Printf("Int64  ID: %d\n", snowflake.Gen())
	fmt.Printf("Int64  ID: %d\n", snowflake.Gen())
	fmt.Printf("Int64  ID: %d\n", snowflake.Gen())
}
