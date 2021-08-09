package main

import (
    "encoding/json"
    "fmt"
)

type Matrix struct {
    row,col int
	data [][]int
}
func GetRows(row,col int) int{
    getrow := Matrix{row: row, col: col}
    return getrow.row
}

func GetColumns(row,col int) int{
    getcol := Matrix{row: row, col: col}
    return getcol.col
}
func AddingMatrices(row,col int) *Matrix {
    matrix1:= make([][]int, row, col)
    matrix2:= make([][]int, row, col)
     left := Matrix{row: row, col: col, data: matrix1}
    right := Matrix{row: row, col: col, data: matrix1}
     add :=  Matrix{row: row, col: col, data:matrix2}
    for i := 0; i < row; i++ {
        left.data[i] = make([]int, col)
        for j := 0; j < col; j++ {
         left.data[i][j] = i+j
        }
    }
    for i := 0; i < row; i++ {
        right.data[i] = make([]int, col)
        for j := 0; j < col; j++ {
           right.data[i][j] = i+j
        }
    }
    for i := 0; i < row; i++ {
        add.data[i] = make([]int, col)
        for j := 0; j < col; j++ {
            add.data[i][j] = left.data[i][j] + right.data[i][j]
        }
    }
    return &add
}
func SetMatrixElements(row,col int) *Matrix {
    matrix4 :=make([][]int, row, col)
    setdata := Matrix{row: row, col: col,data: matrix4}
    for i:=0 ; i<row ; i++ {
        setdata.data[i] = make([]int, col)
        for j:=0 ; j<col ; j++ {
            setdata.data[i][j] = 1
        }
    }
  return &setdata
}

func Printinjson(row,col int) {
    roww := row
    coll := col
   format:=SetMatrixElements(roww,coll)
   value,err := json.Marshal(format)
   if err!= nil {
       fmt.Println(err)
   }
   fmt.Println(value)

fmt.Println("Elements", format)
}


func main() {
 fmt.Println("Enter tne number of rows ")
 var rowsinput int
 fmt.Scanln(&rowsinput)
 fmt.Println("Enter tne number of columns ")
 var columninput int
 fmt.Scanln(&columninput)
 fmt.Println("Number of Rows", GetRows(rowsinput,columninput))
 fmt.Println("Number of columns", GetColumns(rowsinput,columninput))
 set := SetMatrixElements(rowsinput,columninput)
 fmt.Println(set.data)
 added := AddingMatrices(rowsinput,columninput)
 fmt.Println(added.data)
 Printinjson(rowsinput,columninput)
}