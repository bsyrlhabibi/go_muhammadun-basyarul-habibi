package main

// Mendefinisikan struct Kendaraan dengan atribut roda dan kecepatan.
type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

// Mendifinisikan struct Mobil sebagai Kendaraan.
type Mobil struct {
	Kendaraan
}

// Metode Berjalan() untuk meningkatkan KecepatanPerJam sebanyak 10.
func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}

// Metode TambahKecepatan() untuk menambah kecepatan saat berjalan.
func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
	m.KecepatanPerJam += kecepatanBaru
}

func main() {
	// Membuat dua objek Mobil, yaitu mobilCepat dan mobilLamban.
	mobilCepat := Mobil{}
	mobilLamban := Mobil{}

	// Memanggil metode Berjalan() pada mobilCepat tiga kali untuk meningkatkan kecepatannya.
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	// Memanggil metode Berjalan() pada mobilLamban satu kali.
	mobilLamban.Berjalan()

	// Output dari program ini tidak di-handle, sehingga hasil perubahan kecepatan tidak ditampilkan.
}
