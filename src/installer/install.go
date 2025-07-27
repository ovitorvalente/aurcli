package installer

import (
	"fmt"
	"os"
	"os/exec"
)

func InstalarPacote(pacote string) {
	fmt.Printf("🔧 Instalando o pacote: %s\n", pacote)
	cmd := exec.Command("yay", "-S", pacote)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Erro ao instalar o pacote: %v\n", err)
		return
	}
}