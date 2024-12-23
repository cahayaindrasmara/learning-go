package main

import (
	"a21hc3NpZ25tZW50/helper"
	"fmt"
	"strings"
)

var Students = []string{
	"A1234_Aditira_TI",
	"B2131_Dito_TK",
	"A3455_Afis_MI",
}

var StudentStudyPrograms = map[string]string{
	"TI": "Teknik Informatika",
	"TK": "Teknik Komputer",
	"SI": "Sistem Informasi",
	"MI": "Manajemen Informasi",
}

type studentModifier func(string, *string)

func Login(id string, name string) string {
	result := ""

	if len(id) == 0 || len(name) == 0 {
		return "ID or Name is undefined!"
	}

	for _, student := range Students {
		// fmt.Println(student)
		splitStudent := strings.Split(student, "_")
		// fmt.Println(splitStudent)

		studentID := splitStudent[0]
		studentName := splitStudent[1]

		if id == studentID && name == studentName {
			result = "Login berhasil: " + studentName
			break
		} else {
			result = "Login gagal: data mahasiswa tidak ditemukan"
		}
	}

	return result // TODO: replace this
}

func Register(id string, name string, major string) string {
	result := fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)

	if len(id) == 0 || len(name) == 0 || len(major) == 0 {
		return "ID, Name or Major is undefined!"
	}

	for _, student := range Students {
		splitData := strings.Split(student, "_")

		// fmt.Println(splitData)
		studentID := splitData[0]

		if id == studentID {
			// result = "Registrasi gagal: id sudah digunakan"
			// break
			return "Registrasi gagal: id sudah digunakan"
		}
		// else {
		// 	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
		// }
	}
	// newStudents := fmt.Sprintf(", %s, %s, %s", id, name, major)
	newStudents := fmt.Sprintf("%s_%s_%s", id, name, major)
	Students = append(Students, newStudents)
	fmt.Println(Students)

	return result // TODO: replace this
}

func GetStudyProgram(code string) string {
	result := " "

	if program, exists := StudentStudyPrograms[code]; exists {
		result = program
	} else {
		result = "Kode program studi tidak ditemukan"
	}
	return result // TODO: replace this
}

func ModifyStudent(programStudi, nama string, fn studentModifier) string {
	// result := ""

	for i, student := range Students {
		splitData := strings.Split(student, "_")

		studentName := splitData[1]

		if studentName == nama {
			fn(programStudi, &Students[i])
			return "Program studi mahasiswa berhasil diubah."
		}
		// else {
		// 	return "Mahasiswa tidak ditemukan."
		// }

	}
	return "Mahasiswa tidak ditemukan." // TODO: replace this
}

func UpdateStudyProgram(programStudi string, students *string) {
	// splitData := strings.Split(*students, "_")

	// splitData[2] = programStudi

	// *students = strings.Join(splitData, "_")

	*students = strings.Join(strings.Split(*students, "_")[0:2], "_") + "_" + programStudi

	// TODO: answer here
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		for i, student := range Students {
			fmt.Println(i+1, student)
		}

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Change student study program")
		fmt.Println("5. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			helper.ClearScreen()
			var nama, programStudi string
			fmt.Print("Masukkan nama mahasiswa: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan program studi baru: ")
			fmt.Scanln(&programStudi)

			fmt.Println(ModifyStudent(programStudi, nama, UpdateStudyProgram))
			helper.Delay(5)
		case "5":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}

}

// func main() {
// 	fmt.Println(Login(Students, StudentStudyPrograms))
// }
