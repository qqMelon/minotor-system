package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gookit/color"
	"github.com/shirou/gopsutil/cpu"
)

func main() {
	// Init colors
	red := color.FgRed.Render
	green := color.FgGreen.Render

	// Infinit loop for monitoring
	for {
		byCPUs, err := cpu.Percent(time.Second, true)
		if err != nil {
			log.Fatal("Error in gettings CPUs: ", err)
		}

		allCPUs, err := cpu.Percent(time.Second, false)
		if err != nil {
			log.Fatal("Error in getting CPUS data: ", err)
		}

		// Remove previous line with (\r)
		fmt.Print("\033[H\033[2J")

		fmt.Println("========== Utilisation du CPU ==========")
		for i, c := range byCPUs {
			fmt.Printf("CPU %v: %s\n", i, getColoredValue(c, red, green))
		}
		fmt.Println("========== Utilisation total des CPU ==========")
		fmt.Printf("CPU in use: %s\n", getColoredValue(allCPUs[0], red, green))

		// Waiting 1 second before recast
		time.Sleep(time.Second)
	}
}

// Utility function for change color based on used ressources
func getColoredValue(value float64, colorLow, colorHigh func(a ...interface{}) string) string {
	threshold := 80.0
	if value > threshold {
		return colorHigh(fmt.Sprintf("%.2f%%", value))
	}
	return colorLow(fmt.Sprintf("%.2f%%", value))
}
