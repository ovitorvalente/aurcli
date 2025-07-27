package gui

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/ovitorvalente/aurcli/src/api"
	"github.com/ovitorvalente/aurcli/src/installer"
	"github.com/ovitorvalente/aurcli/src/models"
)

func Run() {
	a := app.NewWithID("aurcli")

	w := a.NewWindow("AUR CLI - Instalador")
	w.Resize(fyne.NewSize(400, 800))

	// Entrada de busca
	entrada := widget.NewEntry()
	entrada.SetPlaceHolder("🔍 Digite o nome do pacote AUR...")

	// Lista de resultados
	var pacotes []models.Package
	resultados := widget.NewList(
		func() int {
			return len(pacotes)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Pacote - Descrição")
		},
		func(i int, obj fyne.CanvasObject) {
			p := pacotes[i]
			obj.(*widget.Label).SetText(fmt.Sprintf("%s - %s", p.Name, p.Description))
		},
	)

	resultados.OnSelected = func(i int) {
		if i >= 0 && i < len(pacotes) {
			p := pacotes[i]
			dialog.ShowConfirm("Instalar pacote", fmt.Sprintf("Deseja instalar %s?", p.Name), func(ok bool) {
				if ok {
					installer.InstalarPacote(p.Name)
					dialog.ShowInformation("Instalação", "Instalação iniciada no terminal.", w)
				}
			}, w)
		}
	}

	// Botão de busca
	buscarBtn := widget.NewButton("Buscar", func() {
		termo := strings.TrimSpace(entrada.Text)
		if termo == "" {
			dialog.ShowInformation("Campo vazio", "Digite algo para buscar.", w)
			return
		}

		resultado, err := api.BuscarPacotes(termo)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		pacotes = resultado
		if len(pacotes) == 0 {
			dialog.ShowInformation("Nada encontrado", "Nenhum pacote encontrado.", w)
		}
		resultados.Refresh()
	})

	// Cabeçalho
	cabecalho := container.NewVBox(
		widget.NewLabelWithStyle("Instalador de Pacotes AUR", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		entrada,
		buscarBtn,
		widget.NewSeparator(),
	)

	// Layout final usando Border para expandir a lista
	layoutFinal := container.NewBorder(
		cabecalho,
		nil, nil, nil,
		resultados,
	)

	w.SetContent(container.NewPadded(layoutFinal))
	w.CenterOnScreen()
	w.ShowAndRun()
}
