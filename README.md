# 🧰 aurcli – Um CLI leve para buscar e instalar pacotes do AUR com Go

**aurcli** é uma ferramenta de linha de comando desenvolvida em Go que permite **buscar pacotes no AUR (Arch User Repository)** e **instalá-los facilmente usando o `yay`**. Com uma interface simples e intuitiva, o objetivo é agilizar o processo de busca e instalação sem depender de interfaces gráficas ou comandos complexos.

---

### 🚀 Funcionalidades

* 🔍 Busca pacotes no AUR via API oficial
* 📦 Lista os resultados com nome e descrição
* 🧠 Permite selecionar o pacote desejado para instalar
* ⚙ Integração direta com o `yay` para instalação automatizada
* 💡 Escrita em Go, leve e rápida

---

### 📦 Requisitos

* Arch Linux ou derivado
* [yay](https://github.com/Jguer/yay) instalado
* `Go` (para compilar)
* (Opcional) `jq` caso queira usar um fallback via script

---

### ✨ Exemplos de uso

```bash
./aurcli
🔍 Digite o nome do pacote AUR para buscar: spotify
📦 Resultados encontrados:
[1] spotify - A proprietary music streaming service
[2] spotify-adblock - Adblock for Spotify client
...

Digite o número do pacote para instalar: 1
⚙ Instalando pacote: spotify...
```
---

### 📄 Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
