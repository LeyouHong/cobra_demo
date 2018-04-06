package version

import "fmt"

var (
	Version = "0.0.0"
)

func init() {
	fmt.Print("The current version is : ")
}
