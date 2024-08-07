package main

import (
    "archive/zip"
    "fmt"
    "regexp"
    "io"
    "flag"
    "os"
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
    A := []string{"o","s","h"}
    fmt.Println("Ripperdoc: a tool to quickly grab all the text out of a .docx file")
    fmt.Println("Usage: rdoc [-sh] [-o outfile]")
    for _,x := range A {
        fmt.Println("-" + x + ": " + flag.Lookup(x).Usage)
    }
    os.Exit(0)
}

func main () {
    Oo:=flag.String("o","rdout","Write output to a file instead of stdout")
    Os:=flag.Bool("s",false,"Enable 'smart mode', which tries to reconstruct the docx file's structure")
    Oh:=flag.Bool("h",false,"Display this message")
    flag.Parse()
    if *Oh == true { help() }
    if *Oo != "rdout" { fmt.Println("output file: " + *Oo) }
    if *Os != false { fmt.Println("smart mode on") }
    for _,f := range flag.Args() {
        dump(f)
    }
}
