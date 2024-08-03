package main

import (
    "archive/zip"
    "fmt"
    //"encoding/xml"
    "regexp"
    "io"
)

func main () {
    T:="Impreza Wiring.docx"
    R,e:=zip.OpenReader(T)
    if e!=nil { panic(e) } 
    for _,f := range R.File { 
        if f.Name=="word/document.xml" { r,e:=f.Open(); if e!=nil { panic(e) }
            //x:=xml.NewDecoder(r)
            re,e:=regexp.Compile(`<w:t>(.*?)<\/w:t>`); if e!=nil { panic(e) }
            x,e:=io.ReadAll(r); if e!=nil { panic(e) }
            //fmt.Printf("%q \n", re.FindAllStringSubmatch(string(x),-1))
            for _,m := range re.FindAllStringSubmatch(string(x),-1) {
                fmt.Println(m[1])
            }
            r.Close(); break
        }
    }
    R.Close()
}
