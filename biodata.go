package main

import (
	"fmt"
	"os"
	"strconv"
)

type namaSiswa struct {
	absensi   int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	Program := os.Args
	idSiswa, _ := strconv.Atoi(Program[1])
	getDataKelas(idSiswa)
}

var dataKelas []namaSiswa

func getDataKelas(noabsen int) {
	dataKelas = []namaSiswa{
		{absensi: 1, nama: "M Yoga Irvandi", alamat: "Sumatra Selatan,Martapura", pekerjaan: "Mahasiswa", alasan: "Ingin mempelajari hal baru,sebagai acuan untuk judul Skripsi,menambah pengalaman membuat project\n"},
		{absensi: 2, nama: "I Gede Diva Dwijayana", alamat: "Karangasen, Bali", pekerjaan: "Mahasiswa", alasan: "karena penasaran dengan bahasa pemrograman Golang, ingin mempelajarinya, dan membuat sebuah aplikasi dari Golang\n"},
		{absensi: 3, nama: "Daffa Haryadi", alamat: "Depok", pekerjaan: "Mahasiswa", alasan: "Ingin melanjutkan pembelajaran golang saya dan menambah pengalaman membuat project\n"},
		{absensi: 4, nama: "Denada Ramschie", alamat: "Manado Sulawesi Utara", pekerjaan: "Mahasiswa", alasan: "Untuk mempelajari bidang ilmu baru\n"},
	}
	for _, siswa := range dataKelas {

		if siswa.absensi == noabsen {
			fmt.Println("Nama 			    : ", siswa.nama)
			fmt.Println("Alamat 			    : ", siswa.alamat)
			fmt.Println("Pekerjaan  	 	    : ", siswa.pekerjaan)
			fmt.Println("Alasan memilih kelas Golang : ", siswa.alasan)
		}
	}
}
