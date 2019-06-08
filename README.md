# Noerr
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/strosel/noerr)

A small library to check to check and log errors

### Usage
Turns your code from this:
```Go
v1, err := f1()
if err != nil {
    log.Prinln(err)
}

v2, err := f2()
if err != nil {
    log.Fatal(err)
}

v3, err := f3()
if err != nil {
    panic(err)
}

```
To this:
```Go
v1, err := f1()
noerr.Log(err) //This also returns a bool if you want to nest ifs

v2, err := f2()
noerr.Fatal(err)

v3, err := f3()
noerr.Panic(err)

```