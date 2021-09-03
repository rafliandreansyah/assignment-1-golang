package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	noPeserta           int
	nama                string
	alamat              string
	pekerjaan           string
	alasanMemilihGolang string
}

func main() {

	argument := os.Args
	printData(argument)

}

func printData(arguments []string) {

	startIndex, _ := strconv.Atoi(arguments[1])

	for _, value := range dummyData() {
		if value.noPeserta == startIndex {
			fmt.Println("=====Data Siswa====")
			fmt.Println("Nama:", value.nama)
			fmt.Println("Alamat:", value.alamat)
			fmt.Println("Pekerjaan:", value.pekerjaan)
			fmt.Println("Alasan memilih golang:", value.alasanMemilihGolang)
			fmt.Println("===================")
			return
		}
	}

	fmt.Println("Siswa tidak ditemukan")

}

func dummyData() []student {
	students := []student{
		{1, "Firman Aulia Insani", "Jakarta", "Mahasiswa", "Karena bahasa yang populer dan diminati banyak perusahaan"},
		{2, "Andri Nur Hidayatulloh", "Bandung", "Software Engineer", "Karena bahasa yang populer dan diminati banyak perusahaan"},
		{3, "I Komang Widnyana", "Surabaya", "UX Design", "Bahasa yang sangat wajib dipelajari karena sangat populer"},
		{4, "Erico", "Jakarta", "Android Developer", "Banyak perusahaan yang mencari tenaga ahli pada bahasa golang"},
		{5, "Jose Yolanda Purba", "Jakarta", "Dev Ops", "Karena bahasa yang populer dan diminati banyak perusahaan"},
		{6, "Andri Kuwito", "Yogyakarta", "Node js Developer", "Karena bahasa yang populer dan diminati banyak perusahaan"},
		{7, "Sandy Budiman", "Bandung", "Ios Developer", "Karena bahasa yang populer dan diminati banyak perusahaan"},
		{8, "Rafli Andreansyah", "Blitar", "Mahasiswa", "Karena golang memiliki performance yang sangat bagus sehingga saya mau gunakan untuk project kedepan"},
		{9, "Taufiq Hidayah", "Malang", "Mahasiswa", "Bahasa pemrograman yang populer dan mau saya gunakan untuk project kedepan"},
		{10, "Melvita Sari", "Jember", "Project Manager", "Ingin menambah wawasan pada golang"},
	}

	return students
}
