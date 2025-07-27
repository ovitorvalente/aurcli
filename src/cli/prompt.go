package cli

import (
	"fmt"
	"strconv"

	"github.com/ovitorvalente/aurcli/src/api"
	"github.com/ovitorvalente/aurcli/src/installer"
)

func Execute() {
	var termo string
	fmt.Print("🔍 Digite o nome do pacote AUR para buscar: ")
	fmt.Scanln(&termo)

	pacotes, err := api.BuscarPacotes(termo)
	if err != nil {
		fmt.Println("❌ Erro ao buscar pacotes:", err)
		return
	}
	if len(pacotes) == 0 {
		fmt.Println("❌ Nenhum pacote encontrado.")
		return
	}

	fmt.Println("\n📦 Pacotes encontrados:")
	for i, p := range pacotes {
		fmt.Printf("[%d] %s - %s\n", i+1, p.Name, p.Description)
	}

	var escolhaStr string
	fmt.Print("\nDigite o número do pacote que deseja instalar: ")
	fmt.Scanln(&escolhaStr)

	escolha, err := strconv.Atoi(escolhaStr)
	if err != nil || escolha < 1 || escolha > len(pacotes) {
		fmt.Println("❌ Escolha inválida.")
		return
	}

	pacote := pacotes[escolha-1].Name
	installer.InstalarPacote(pacote)
}
