package general

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "v0.1.0"

func CurrentVersion(c *cobra.Command, args []string){
	fmt.Printf("Version: %v\n", Version)
}