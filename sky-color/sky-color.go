package main

import "fmt"

type Light struct {
	Wavelength int
	Intensity  int
}

func scatteredLight(light Light) Light {
	return Light{light.Wavelength, light.Intensity << 40 / (light.Wavelength * light.Wavelength * light.Wavelength * light.Wavelength)}
}

func main() {
	red := Light{700, 6}
	green := Light{546, 7}
	blue := Light{436, 8}

	fmt.Println("sky color!!")
	fmt.Println("red: ", scatteredLight(red))
	fmt.Println("green: ", scatteredLight(green))
	fmt.Println("blue: ", scatteredLight(blue))
}
