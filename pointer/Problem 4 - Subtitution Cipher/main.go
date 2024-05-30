package main

import "fmt"

type student struct {
	name       string
	nameEncode string
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var namaEncode = ""
	// your code here
	for i := 0; i < len(s.name); i++ {
		if s.name[i] > 96 && s.name[i]+3 > 122 {
			namaEncode += string(s.name[i] - 23)
		} else if s.name[i] < 91 && s.name[i]+3 > 90 {
			namaEncode += string(s.name[i] - 23)
		} else {
			namaEncode += string(s.name[i] + 3)
		}
	}
	return namaEncode
}

func (s *student) Decode() string {
	var namaDecode = ""
	// your code here
	for i := 0; i < len(s.name); i++ {
		if s.name[i] > 96 && s.name[i]-3 < 97 {
			namaDecode += string(s.name[i] + 23)
		} else if s.name[i] < 91 && s.name[i]-3 < 41 {
			namaDecode += string(s.name[i] + 23)
		} else {
			namaDecode += string(s.name[i] - 3)
		}
	}
	return namaDecode
}
func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of Student's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}
