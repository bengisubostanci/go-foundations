// Go'da klasik anlamda class (sınıf) yapısı yoktur. Bunun yerine verileri gruplamak için struct kullanırız. Pointers (işaretçiler) ise bir verinin kopyasını çıkarmak yerine, bellekteki adresini doğrudan referans göstermemizi sağlar. Bu, özellikle büyük veri yapılarında performansı devasa oranda artırır.

package main

import "fmt"

// 1. STRUCT TANIMLAMA
// Go'da bir veri modelini (örneğin bir veri hattı görevini veya kullanıcıyı) tanımlamak için struct kullanırız.
type DataTask struct {
	ID       int
	TaskName string
	Priority string
	IsDone   bool
}

// 2. POINTER MANTIĞINI ANLAMAK İÇİN FONKSİYONLAR
// Bu fonksiyon verinin KOPYASINI alır (Pass by Value). Asıl veriyi değiştiremez.
func completeTaskByValue(task DataTask) {
	task.IsDone = true
}

// Bu fonksiyon verinin BELLEK ADRESİNİ alır (Pass by Reference - Pointer ile).
// Başımdaki '*' işareti bu fonksiyonun bir pointer beklediğini söyler.
func completeTaskByPointer(task *DataTask) {
	task.IsDone = true // Go, pointer üzerinden nesneye erişirken otomatik olarak adresi çözer (de-referencing).
}

func main() {
	fmt.Println("--- Go Foundations: Structs and Pointers ---")

	// --- STRUCT INITIALIZATION (NESNE OLUŞTURMA) ---
	task1 := DataTask{
		ID:       101,
		TaskName: "Ingest S3 Logs to Database",
		Priority: "High",
		IsDone:   false,
	}

	fmt.Printf("Initial Task Status: %s (Completed: %t)\n", task1.TaskName, task1.IsDone)

	// --- VALUE VS POINTER DENEYİ ---

	// Deney 1: Değer ile fonksiyon çağırma (Kopyalama)
	completeTaskByValue(task1)
	fmt.Printf("After completeTaskByValue: Completed = %t (Değişmedi!)\n", task1.IsDone)

	// Deney 2: Pointer ile fonksiyon çağırma (Adres Gösterme)
	// Bir değişkenin başına '&' koyarsak, o değişkenin bellekteki adresini almış oluruz.
	completeTaskByPointer(&task1)
	fmt.Printf("After completeTaskByPointer: Completed = %t (Değişti!)\n", task1.IsDone)

	// --- 3. DOĞRUDAN POINTER İLE ÇALIŞMAK ---
	x := 42
	var p *int = &x // p, x'in bellekteki adresini tutan bir pointer'dır.

	fmt.Println("\nPointer Basics:")
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Memory Address of x (&x): %v\n", &x)
	fmt.Printf("Value stored in pointer p: %v\n", p)
	fmt.Printf("Value at pointer p (*p): %d\n", *p) // *p ile adresteki değere eriştik
}
