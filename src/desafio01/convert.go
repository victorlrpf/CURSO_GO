package main

import "fmt"

const pontoEbuli = 373.0

func main() {
	tempK := pontoEbuli
	tempC := tempK - 273.0
	fmt.Printf("Temperatura em K = %1.f\nTempratura de ebulição em °C é: %.1f\n", tempK, tempC)
}
