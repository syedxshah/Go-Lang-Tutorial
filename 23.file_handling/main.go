package main

import (
	//"bufio"
	"fmt"
	"os"
)

func main() {

	// f, err := os.Open("file.txt")

	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// } else {
	// 	fileinfo, err := f.Stat()

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println("File Name:", fileinfo.Name())
	// 	fmt.Println("File Size in bytes :", fileinfo.Size())
	// 	fmt.Println("File Mode:", fileinfo.Mode())
	// 	fmt.Println("File ModTime:", fileinfo.ModTime())
	// 	fmt.Println("Is Directory:", fileinfo.IsDir())
	// }

	// // file read

	// f, err := os.Open("file.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 12)

	// d, err := f.Read(buf)

	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	fmt.Print(string(buf[i]))
	// }

	// fmt.Print("\n")
	// fmt.Println("Bytes Read:", d)
	// fmt.Println("Data Read:", buf)

	// second way to read file

	// f, err := os.ReadFile("file.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(f))

	// // read folders

	// dir , err := os.Open(".")

	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fldinfo , err := dir.ReadDir(-1)

	// for _, v := range fldinfo {

	// 	fmt.Println("Name:", v.Name())
	// 	fmt.Println("Is Dir:", v.IsDir())
	// }

	// // create file
	// f, err := os.Create("file2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// d, err := f.WriteString("Hello World\n")

	// f.WriteString("Test Two \n")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Data Written:", d)

	// 	// using bufio to copy file data and paste to another file

	// 	s, err := os.Open("file2.txt")

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	defer s.Close()

	// 	d, err := os.Create("file3.txt")

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	defer d.Close()

	// 	reader := bufio.NewReader(s)

	// 	writer := bufio.NewWriter(d)

	// 	for {
	// 		b, err := reader.ReadByte()

	// 		if err != nil {
	// 			if err.Error() == "EOF" {
	// 				break
	// 			}
	// 			panic(err)
	// 		}
	// 		e := writer.WriteByte(b)
	// 		if e != nil {
	// 			panic(e)
	// 		}
	// 	}
	// 	writer.Flush()

	// 	fmt.Println("File Copied Successfully")

	// delete a file

	err := os.Remove("file3.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File Deleted Successfully")

}
