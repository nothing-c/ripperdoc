package main

import (
    "archive/zip"
    "fmt"
)

func main () {
    T:="Impreza Wiring.docx"
    R,e:=zip.OpenReader(T)
    if e!=nil { panic(e) } 
    for _,f := range R.File { fmt.Println(f.Name) }
    defer R.Close()
}
