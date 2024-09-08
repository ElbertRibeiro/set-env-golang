package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	folderPath := "C:\\caminho\\para\\sua\\pasta"

	batFileName := "set_env_var.bat"

	batContent := fmt.Sprintf(`@echo off
setx PASTA_VAR "%s"
echo Variável PASTA_VAR definida para: %s
`, folderPath, folderPath)

	err := os.WriteFile(batFileName, []byte(batContent), 0644)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo .bat:", err)
		return
	}

	cmd := exec.Command("cmd", "/C", batFileName)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Erro ao executar o arquivo .bat:", err)
		return
	}

	fmt.Println("Script .bat executado com sucesso.")
}
