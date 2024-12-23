package main

import (
	"a21hc3NpZ25tZW50/helper"
	"fmt"
	"strings"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {

	if len(id) == 0 || len(name) == 0 {
		return "ID or Name is undefined!"
	}

	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	studentRecords := strings.Split(Students, ",")
	// fmt.Println(split1)
	for _, record := range studentRecords {
		parts := strings.Split(strings.TrimSpace(record), "_")
		// fmt.Println(parts)

		studentID, studentName, programStudy := parts[0], parts[1], parts[2]

		if studentID == id && studentName == name {
			// return "Login berhasil: " + name

			programs := strings.Split(StudentStudyPrograms, ",")

			for _, program := range programs {
				programParts := strings.Split(strings.TrimSpace(program), "_")

				programCode := programParts[0]

				if programStudy == programCode {
					return fmt.Sprintf("Login berhasil: %s (%s)", studentName, programCode)
				}
			}
		}

	}

	return "Login gagal: data mahasiswa tidak ditemukan" // TODO: replace this
}

func Register(id string, name string, major string) string {

	if len(id) == 0 || len(name) == 0 || len(major) == 0 {
		return "ID, Name or Major is undefined!"
	}

	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	StudentsRecords := strings.Split(Students, "'")
	for _, record := range StudentsRecords {
		parts := strings.Split(record, "_")

		studentID := parts[0]

		if studentID == id {
			return "Registrasi gagal: id sudah digunakan"
		}
	}

	Students += fmt.Sprintf(", %s %s %s", id, name, major)
	fmt.Println(Students)
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major) // TODO: replace this
}

func GetStudyProgram(code string) string {
	var result string
	if len(code) == 0 {
		return "Code is undefined!"
	}

	programStudy := strings.Split(StudentStudyPrograms, ",")
	fmt.Println(programStudy)

	for _, program := range programStudy {
		programParts := strings.Split(strings.TrimSpace(program), "_")

		programCode, programName := programParts[0], programParts[1]
		fmt.Println(programCode)

		if programCode == code {
			result = programName
		}
	}

	return result // TODO: replace this
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

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
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}

	// fmt.Println(GetStudyProgram(StudentStudyPrograms))
	// fmt.Println(Login(Students, StudentStudyPrograms))

}
