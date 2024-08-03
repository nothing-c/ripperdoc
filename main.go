package main

import (
    "archive/zip"
    "fmt"
)

func main () {
    T:="Impreza Wiring.docx"
    R,e:=zip.OpenReader(T)
    if e!=nil { panic(e) } 
    
    R.Close()
}
