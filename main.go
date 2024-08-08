package main

import (
    "archive/zip"
    "fmt"
    "regexp"
    "io"
    "flag"
    "os"
    "sync"
)

func dump (T string, C chan string) {
    var S string
    R,e:=zip.OpenReader(T)
    if e!=nil { panic(e) } 
    for _,f := range R.File { 
        if f.Name=="word/document.xml" { r,e:=f.Open(); if e!=nil { panic(e) }
            re,e:=regexp.Compile(`<w:t.*?>(.*?)<\/w:t>`); if e!=nil { panic(e) }
            x,e:=io.ReadAll(r); if e!=nil { panic(e) }
            for _,m := range re.FindAllStringSubmatch(string(x),-1) {
                S += m[1] + "\n"
                //fmt.Println(m[1])
            }
            r.Close(); break
        }
    }
    R.Close()
    C <- S
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
    var OF *os.File
    Oo:=flag.String("o","","Write output to a file instead of stdout")
    Os:=flag.Bool("s",false,"Enable 'smart mode', which tries to reconstruct the docx file's structure")
    Oh:=flag.Bool("h",false,"Display this message")
    flag.Parse()
    C:=make(chan string,flag.NArg()) // Enough to buffer all the output 
    var W sync.WaitGroup
    var e error // Have to do this to make the compiler happy -__-
    if *Oo != "" { OF,e = os.OpenFile(*Oo, os.O_RDWR|os.O_CREATE, 0644); if e!=nil { panic(e) } } else { OF = os.Stdout }
    if *Oh == true { help() }
    if *Os != false { fmt.Println("smart mode on") }
    for _,f := range flag.Args() {
        W.Add(1)
        go func(s string) {
            defer W.Done()
            dump(s,C)
        }(f)
    }
    W.Wait()
    for len(C) > 0 {
        _,e := io.WriteString(OF,<-C + "\n~\n"); if e!=nil { panic(e) }
    }
    os.Exit(0)
}
