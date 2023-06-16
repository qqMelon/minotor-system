package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gookit/color"
	"github.com/shirou/gopsutil/cpu"
)

func main() {
	// Initialiser les couleurs
	red := color.FgRed.Render
	green := color.FgGreen.Render

	// Boucle infinie pour surveiller l'utilisation du CPU
	for {
		byCPUs, err := cpu.Percent(time.Second, true)
		if err != nil {
			log.Fatal("Error in gettings CPUs: ", err)
		}

		allCPUs, err := cpu.Percent(time.Second, false)
		if err != nil {
			log.Fatal("Error in getting CPUS data: ", err)
		}

		// Effacer l'affichage précédent en utilisant le caractère de retour à la ligne (\r)
		fmt.Print("\033[H\033[2J")

		// Afficher le tableau avec les couleurs
		fmt.Println("========== Utilisation du CPU ==========")
		for i, c := range byCPUs {
			fmt.Printf("CPU %v: %s\n", i, getColoredValue(c, red, green))
		}
		fmt.Println("========== Utilisation total des CPU ==========")
		fmt.Printf("CPU in use: %s\n", getColoredValue(allCPUs[0], red, green))

		//fmt.Printf("CPU 1: %s\n", getColoredValue(byCPUs[0], red, green))
		//fmt.Printf("CPU 2: %s\n", getColoredValue(byCPUs[1], red, green))
		// Ajoutez davantage de lignes pour afficher les autres cœurs du CPU

		// Pause de 1 seconde avant de vérifier à nouveau l'utilisation du CPU
		time.Sleep(time.Second)
	}
}

// Fonction utilitaire pour obtenir la valeur colorée en fonction du seuil
func getColoredValue(value float64, colorLow, colorHigh func(a ...interface{}) string) string {
	threshold := 80.0
	if value > threshold {
		return colorHigh(fmt.Sprintf("%.2f%%", value))
	}
	return colorLow(fmt.Sprintf("%.2f%%", value))
}
