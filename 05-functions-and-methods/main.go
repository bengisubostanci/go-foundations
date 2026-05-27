//Go'da "fonksiyon" ve "metot" kavramları birbirinden farklıdır.

//Fonksiyonlar: Bağımsız çalışır, girdi alır ve çıktı üretir. Go fonksiyonlarının en güzel yanı birden fazla değer dönebilmeleridir (örneğin hem sonucu hem de bir hata mesajını aynı anda dönebilirler).

//Metotlar: Doğrudan bir struct yapısına bağlanırlar (diğer dillerdeki sınıf içi metotlar gibi).

package main

import (
	"errors"
	"fmt"
)

// 1. STANDART STRUCT TANIMLAMA
type CloudServer struct {
	HostName string
	IP       string
	IsActive bool
}

// 2. ÇOKLU DEĞER DÖNEN FONKSİYON (MULTIPLE RETURN VALUES)
// Go'da bir fonksiyon hata yönetimi için genellikle (sonuç, hata) şeklinde iki değer döner.
func divideMetrics(totalRequests int, totalSeconds int) (int, error) {
	if totalSeconds == 0 {
		// Hata durumunu errors.New ile fırlatıyoruz
		return 0, errors.New("division by zero: seconds cannot be zero")
	}
	return totalRequests / totalSeconds, nil // Hata yoksa 'nil' (null) dönüyoruz
}

// 3. STRUCT METODU (METHOD WITH VALUE RECEIVER)
// Bu metot sadece CloudServer nesnesi üzerinden çağrılabilir. Veriyi sadece okur (kopyasını alır).
func (s CloudServer) DisplayDetails() {
	fmt.Printf("Server: %s | IP: %s | Active: %t\n", s.HostName, s.IP, s.IsActive)
}

// 4. STRUCT METODU (METHOD WITH POINTER RECEIVER)
// Başına '*' koyarak nesnenin adresini alıyoruz. Bu sayede metot, nesnenin durumunu kalıcı olarak değiştirebilir.
func (s *CloudServer) ToggleStatus() {
	s.IsActive = !s.IsActive // Sunucu durumunu tersine çevirir (true ise false, false ise true)
}

func main() {
	fmt.Println("--- Go Foundations: Functions and Methods ---")

	// --- FONKSİYON KULLANIMI VE HATA KONTROLÜ ---
	// divideMetrics fonksiyonunu çağırıyoruz ve dönen iki değeri (reqPerSec, err) yakalıyoruz.
	reqPerSec, err := divideMetrics(1200, 4)
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Printf("Traffic Density: %d requests/second\n", reqPerSec)
	}

	// Hata senaryosu denemesi
	_, err = divideMetrics(100, 0) // İlk değeri kullanmayacağımız için '_' (blank identifier) ile geçtik
	if err != nil {
		fmt.Println("Expected Error Caught:", err)
	}

	// --- METOT KULLANIMI ---
	fmt.Println("\nExecuting Struct Methods:")
	
	server01 := CloudServer{
		HostName: "prod-api-01",
		IP:       "10.0.1.50",
		IsActive: false,
	}

	// Value Receiver metodunu çağırma
	server01.DisplayDetails()

	// Pointer Receiver metodunu çağırma (İçerideki IsActive değerini değiştirecek)
	server01.ToggleStatus()
	
	fmt.Println("After altering status via Pointer Method:")
	server01.DisplayDetails()
}
