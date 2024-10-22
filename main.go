package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Kelimeyi komut satırından al
	dosya1 := os.Args[1][9:]
	word := os.Args[2]
	words := os.Args[3]

	// \n kaçış dizisini gerçek satır sonu karakterine dönüştür
	word = strings.ReplaceAll(word, "\\n", "\n")

	dosya, _ := os.Create(dosya1)

	// Dosyayı oku
	var fileContent string
	if words == "standard" {
		file, err := os.ReadFile("standard.txt")
		if err != nil {
			fmt.Println("Dosya okunurken bir hata oluştu")
			panic(err)
		}
		fileContent = string(file)
	} else if words == "shadow" {
		file, err := os.ReadFile("shadow.txt")
		if err != nil {
			fmt.Println("Dosya okunurken bir hata oluştu")
			panic(err)
		}
		fileContent = string(file)
	} else if words == "thinkertoy" {
		file, err := os.ReadFile("thinkertoy.txt")
		if err != nil {
			fmt.Println("Dosya okunurken bir hata oluştu")
			panic(err)
		}
		fileContent = string(file)
	} else {
		fmt.Println("Geçersiz kelime grubu:", words)
		return
	}

	// Dosya içeriğini satırlara ayır
	lines := strings.Split(fileContent, "\n")

	// Her kelime için ASCII sanatını yazdır
	for i, line := range strings.Split(word, "\n") {
		if line == "" {
			if i != 0 {
				_, err := dosya.WriteString("\n")
				if err != nil {
					log.Fatal(err)
				}
			}
			continue
		}

		// Her bir karakteri işleyerek ASCII sanatını oluştur
		for h := 1; h < 9; h++ {
			for _, char := range line {
				printAsciiArtForCharacter(char, h, lines, dosya)
			}
			_, err := dosya.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	_, err := dosya.WriteString("\n")
	if err != nil {
		log.Fatal(err)
	}
}

func printAsciiArtForCharacter(char rune, lineIndex int, lines []string, dosya *os.File) {
	// ASCII karakterinin indeksini hesapla
	index := (int(char) - 32) * 9

	// ASCII sanatını yazdır
	if index >= 0 && index+8 <= len(lines) {
		_, err := dosya.WriteString(lines[index+lineIndex])
		if err != nil {
			log.Fatal(err)
		}
	}
}
