package main

import (
    "archive/zip"
    "fmt"
    "regexp"
    "io"
    "flag"
//    "os"
)

func dump (T string) {
    R,e:=zip.OpenReader(T)
    if e!=nil { panic(e) } 
    for _,f := range R.File { 
        if f.Name=="word/document.xml" { r,e:=f.Open(); if e!=nil { panic(e) }
            re,e:=regexp.Compile(`<w:t.*?>(.*?)<\/w:t>`); if e!=nil { panic(e) }
            x,e:=io.ReadAll(r); if e!=nil { panic(e) }
            for _,m := range re.FindAllStringSubmatch(string(x),-1) {
                fmt.Println(m[1])
            }
            r.Close(); break
        }
    }
    R.Close()
}

func help () {
    //help function here
}

func main () {
    Ob:=flag.Bool("b",false,"Run in batch mode")
    Oo:=flag.String("o","rdout","Output file")
    Os:=flag.Bool("s",false,"Enable 'smart mode', which tries to reconstruct the docx file's structure")
    Oh:=flag.Bool("h",false,help())
    flag.Parse()
    for _,f := range flag.Args() {
        dump(f)
    }
}
