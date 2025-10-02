package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// color codes
const (
	Reset     = "\033[0m"
	Bold      = "\033[1m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Magenta   = "\033[35m"
	Cyan      = "\033[36m"
	White     = "\033[37m"
	BrightRed = "\033[91m"
)

// for the custome name and user name printing
type Shell struct {
	history    []string
	currentDir string
}

func NewShell() *Shell {
	cwd, _ := os.Getwd() //cwd - current working directory
	return &Shell{
		history:    make([]string, 0),
		currentDir: cwd,
	}
}

func (s *Shell) printBanner() {
	banner := "_    .  ,   .           .\n" +
		"    *  / \\_ *  / \\_      _  *        *   /\\'__        *\n" +
		"      /    \\  /    \\,   ((        .    _/  /  \\  *'.\n" +
		" .   /\\/\\  /\\/ :' __ \\_  `          _^/  ^/    `--.\n" +
		"    /    \\/  \\  _/  \\-'\\      *    /.' ^_   \\_   .'\\  *\n" +
		"  /\\  .-   `. \\/     \\ /==~=-=~=-=-;.  _/ \\ -. `_/   \\\n" +
		" /  `-.__ ^   / .-'.--\\ =-=~_=-=~=^/  _ `--./ .-'  `-\n" +
		"/jgs     `.  / /       `.~-^=-=~=^=.-'      '-._ `._"

	fmt.Println(Cyan + banner + Reset)
	fmt.Println(Yellow + "Type 'help' for available commands, 'exit' to quit\n" + Reset)
}

func (s *Shell) getPrompt() string {
	username := os.Getenv("USER")
	if username == "" {
		username = os.Getenv("USERNAME")
	}

	if username == "" {
		username = "User"
	}

	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "Local host"
	}

	shortPath := s.currentDir
	home, _ := os.UserHomeDir()

	if after, ok := strings.CutPrefix(s.currentDir, home); ok {
		shortPath = "~" + after
	}

	currentTime := time.Now().Format("15:04:05")
	return fmt.Sprintf(
		"%s[%s]%s %s%s@%s%s:%s%s%s %s$ %s",
		Cyan, currentTime, Reset,
		Green+Bold, username, hostname, Reset,
		Blue+Bold, shortPath, Reset,
		Yellow, Reset,
	)
}

func (s *Shell) executeCommand(input string) {
	input = strings.TrimSpace(input)
	if input == "" {
		return
	}

	s.history = append(s.history, input)
	args := strings.Fields(input)
	command := args[0]

	switch command {
	case "exit", "quit":
		fmt.Println(Cyan + "Good, Bye!" + Reset)
		os.Exit(0)
	case "cd":
		s.changeDirectory(args)
	case "history":
		s.showHistroy()
	case "clear", "cls":
		s.clearScreen()
	case "help":
		s.ShowHelp()
	case "pwd":
		fmt.Println(Blue + s.currentDir + Reset)
	case "ls":
		s.ShowListFiles(args)
	default:
		s.runExternalCommand(args)
	}
}

func (s *Shell) runExternalCommand(args []string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", strings.Join(args, " "))
	} else {
		cmd = exec.Command(args[0], args[1:]...)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmd.Dir = s.currentDir
	if err := cmd.Run(); err != nil {
		fmt.Println(Red+"Error: ", err.Error()+Reset)
	}
}

func (s *Shell) ShowListFiles(args []string) {
	path := "."
	if len(args) >= 2 {
		path = args[1]
	}

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(Red + "Error: cannot list files for this path" + Reset)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(Blue + file.Name() + "/" + Reset)
		} else {
			fmt.Println(file.Name())
		}
	}
}

func (s *Shell) ShowHelp() {
	help := `
` + Cyan + Bold + `Available Commands:` + Reset + `
` + Green + `  cd [dir]` + Reset + `     - Change directory
` + Green + `  pwd` + Reset + `          - Print working directory
` + Green + `  ls/dir` + Reset + `       - List directory contents
` + Green + `  history` + Reset + `      - Show command history
` + Green + `  clear/cls` + Reset + `    - Clear the screen
` + Green + `  help` + Reset + `         - Show this help message
` + Green + `  exit/quit` + Reset + `    - Exit the shell

` + Yellow + `All other commands are passed to the system shell.` + Reset + `
`
	fmt.Println(help)
}

func (s *Shell) clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
	s.printBanner()
}

func (s *Shell) showHistroy() {
	if len(s.history) == 0 {
		fmt.Println(Yellow + "No Command Histroy yet" + Reset)
		return
	}

	fmt.Println(Cyan + "\n Command Histroy: " + Reset)
	for i, cmd := range s.history {
		fmt.Printf("%s%4d%s  %s\n", Magenta, i+1, Reset, cmd)
	}
	fmt.Println()
}

func (s *Shell) changeDirectory(args []string) {
	var path string
	if len(args) < 2 {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(Red + "Error: " + err.Error() + Reset)
			return
		}
		path = home
	} else {
		path = args[1]
	}

	//expand ~ to home directory
	if strings.HasPrefix(path, "~") {
		home, _ := os.UserHomeDir()
		path = filepath.Join(home, path[1:])
	}

	if err := os.Chdir(path); err != nil {
		fmt.Println(Red + "Error: " + err.Error() + Reset)
	} else {
		s.currentDir, _ = os.Getwd()
	}
}

func main() {
	shell := NewShell()
	shell.clearScreen()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(shell.getPrompt())
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(Red + "\nHm... Somethings wrong! " + err.Error() + Reset)
			break
		}
		shell.executeCommand(input)
	}
}
