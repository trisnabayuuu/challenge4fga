package main

import (
	"fmt"
	"os"
	"strconv"
)

type teman struct {
	nama      string
	alamat    string
	pekerjaan string
	num       int
	alasan    string
}

func GetData(num int) teman {
	switch num {
	case 1:
		return teman{
			"abdi",
			"Jl.cianjur No.1",
			"IT",
			1,
			"Ingin belajar tentang Golang",
		}
	case 2:
		return teman{
			"sukri",
			"Jl.mangga No.2",
			"pengangguran",
			2,
			"ingin dapat pekerjaan",
		}
	case 3:
		return teman{
			"somad",
			"Jl.buku No.3",
			"software engineer",
			3,
			"ingin belajar teknologi baru",
		}
	case 4:
		return teman{
			"mamat",
			"Jl.papua No.4",
			"Mahasiswa",
			4,
			"lulus tepat waktu",
		}
	default:
		return teman{}
	}
}
func main() {
	arguments := os.Args[1]
	num, err := strconv.Atoi(arguments)
	if err != nil {
		fmt.Println("data harus integer")
		return
	}

	dataTeman := GetData(num)
	if dataTeman.nama == "" {
		fmt.Println("data tidak")
		return
	}

	fmt.Println("Data teman:")
	fmt.Println("nama:", dataTeman.nama)
	fmt.Println("alamat:", dataTeman.alamat)
	fmt.Println("pekerjaan:", dataTeman.pekerjaan)
	fmt.Println("alasan", dataTeman.alasan)
}
