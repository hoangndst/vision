package vision

import (
	"fmt"
	"github.com/hoangndst/vision/cmd"
	"math/rand"
	"os"
	"time"
)

// @title Vision Backend API
// @version v0.1.0
// @description This is the Vision Backend API.
func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	command := cmd.NewDefaultVisionCommand()

	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
