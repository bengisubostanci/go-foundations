package main

import "fmt"

func main() {
	fmt.Println("--- Go Foundations: Control Structures & Loops ---")

	// --- 1. KONTROL YAPILARI (IF - ELSE) ---
	// Senaryo: Bir veri hattındaki (data pipeline) doluluk oranına göre sistem uyarısı üretmek.
	bufferLoad := 85 

	// Go'da if-else koşullarında parantez () kullanılmaz, süslü parantez {} zorunludur.
	if bufferLoad >= 90 {
		fmt.Println("CRITICAL: Buffer capacity is almost full! Scaling required.")
	} else if bufferLoad >= 75 {
		fmt.Println("WARNING: High buffer load detected. Monitoring performance.")
	} else {
		fmt.Println("SUCCESS: Buffer load is within safe parameters.")
	}

	// --- 2. DÖNGÜLER (ONLY 'FOR' LOOP EXISTS) ---
	// Go'da 'while' döngüsü yoktur, tüm döngüler 'for' ile kurulur.

	// Klasik Üçlü Döngü (i := başlangıç; koşul; artış)
	fmt.Println("\nExecuting fixed-step loop:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("Processing batch chunk %d\n", i)
	}

	// "While" Tarzı Döngü
	// Başlangıç ve artış adımlarını dışarıda bırakarak for döngüsünü while gibi kullanabiliriz.
	fmt.Println("\nExecuting conditional (while-like) loop:")
	counter := 1
	for counter <= 3 {
		fmt.Printf("Retry attempt: %d\n", counter)
		counter++ 
	}
}
