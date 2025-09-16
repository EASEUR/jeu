package logo

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else { // Linux, macOS
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func Logoontop() {
	Clear()
	fmt.Println("\n\n\n\n\n")
	fmt.Println("\033[31m" + `
▀████    ▐████▀     ███        ▄████████    ▄████████  ▄████████    ▄█   ▄█▄ 
  ███▌   ████▀  ▀█████████▄   ███    ███   ███    ███ ███    ███   ███ ▄███▀ 
   ███  ▐███       ▀███▀▀██   ███    ███   ███    ███ ███    █▀    ███▐██▀   
   ▀███▄███▀        ███   ▀  ▄███▄▄▄▄██▀   ███    ███ ███         ▄█████▀    
   ████▀██▄         ███     ▀▀███▀▀▀▀▀   ▀███████████ ███        ▀▀█████▄    
  ▐███  ▀███        ███     ▀███████████   ███    ███ ███    █▄    ███▐██▄   
 ▄███     ███▄      ███       ███    ███   ███    ███ ███    ███   ███ ▀███▄ 
████       ███▄    ▄████▀     ███    ███   ███    █▀  ████████▀    ███   ▀█▀ 
                              ███    ███                           ▀         
` + "\033[0m")
	fmt.Println("\n\n\n")
}
