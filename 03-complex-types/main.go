package main

import "fmt"

func main() {
	fmt.Println("--- Go Foundations: Complex Data Types (Slices & Maps) ---")

	// --- 1. ARRAYS (DİZİLER) ---
	// Go'da dizilerin boyutu sabittir, sonradan değiştirilemez. Sektörde nadiren doğrudan kullanılırlar.
	var fixedArray [3]string
	fixedArray[0] = "Data-1"
	fixedArray[1] = "Data-2"
	fixedArray[2] = "Data-3"

	// --- 2. SLICES (DİLİMLER) ---
	// Boyutu dinamiktir. Go'da listelerle çalışırken standart olarak 'Slice' kullanırız.
	// Tanımlarken köşeli parantezin içine boyut YAZILMAZ.
	dataSourcePipeline := []string{"Kafka", "Spark", "Airflow"}

	// Slices yapısına yeni eleman eklemek için 'append' fonksiyonu kullanılır.
	dataSourcePipeline = append(dataSourcePipeline, "AWS-S3")

	fmt.Println("\nIterating over a Slice using 'range':")
	// 'range' anahtar kelimesi bize hem indeksi (i) hem de o indekteki değeri (value) döner.
	for index, value := range dataSourcePipeline {
		fmt.Printf("Index: %d, Technology: %s\n", index, value)
	}

	// --- 3. MAPS (ANAHTAR - DEĞER ÇİFTLERİ) ---
	// Diğer dillerdeki Dictionary veya HashMap yapısıdır.
	// make(map[KeyType]ValueType) şeklinde initialize edilir.
	systemMetrics := make(map[string]int)

	// Eleman ekleme
	systemMetrics["cpu_usage"] = 45
	systemMetrics["memory_usage"] = 72
	systemMetrics["active_jobs"] = 8

	// Haritadan veri okuma
	fmt.Println("\nReading from a Map:")
	fmt.Printf("Current Memory Usage: %d%%\n", systemMetrics["memory_usage"])

	// Go'ya özgü Map kontrolü: Bir anahtar map içinde var mı yok mu kontrol etmek.
	// 'ok' değişkeni boolean bir değer döner (true/false).
	val, ok := systemMetrics["disk_space"]
	if ok {
		fmt.Printf("Disk Space: %d GB\n", val)
	} else {
		fmt.Println("Metric 'disk_space' not found in the system metrics.")
	}
}
