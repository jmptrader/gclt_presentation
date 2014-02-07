package main
import (
    "github.com/romanoff/gofh"
    "fmt"
    "os"
)
func main() {
	fh := gofh.Init()
	fh.HandleCommand("init", initProject)
	fh.HandleCommand("deploy", deployProject)
	fh.SetDefaultHandler(showUsage)
	fh.Parse(os.Args[1:])
}
func showUsage() {
	fmt.Println("Please, use 'project init' or 'project deploy' command")
}
func initProject(options map[string]string) {
	fmt.Println("Your init project code goes here")
}
func deployProject(options map[string]string) {
	fmt.Println("Your deploy project code goes here")
}
