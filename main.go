package main

import (
	"io"
	"strings"
)

func main(){

	mergeReaders(strings.NewReader("abcd"), strings.NewReader("efg"))
}

func mergeReaders(r1, r2 io.Reader) (io.Reader, error){

	read1 := make([]byte, 1)
	read2 := make([]byte, 1)
	result :=""
	for {
		data1, err1 := r1.Read(read1)
		data2, err2 := r2.Read(read2)

		if err1 != nil {
			if err1 == io.EOF {
				return strings.NewReader(result), nil
			}
		} else if err2 != nil{
			if err2 == io.EOF {
				return strings.NewReader(result), nil
			}
		}

		if data2>0 && data1 >0 {
			result += string(read1[:data1]) + string(read2[:data2])
		}
	}
}

//func getLenOfReader(stream io.Reader) int{
//	buf := new(bytes.Buffer)
//	buf.ReadFrom(stream)
//	fmt.Println(buf.String())
//	return buf.Len()
//}
