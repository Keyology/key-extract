package main
import (
	"bufio"
	"fmt"
	"key-extract/extract"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$Keyextract ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)

	case "start":
		fmt.Println("Welcome to key extract please enter the command extract followed by the link")
	case "extract":
		fmt.Println("***THIS IS ARGS:***", arrCommandStr[1:])
		fmt.Println("extracting site....")
		link := arrCommandStr[1:]
		result := strings.Join(link, "")
		fmt.Println("data extracted")
		extract.GetLink(result)
	case "help":
		fmt.Println("")


	}

	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
