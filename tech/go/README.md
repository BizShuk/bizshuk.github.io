# Golang

### ref
- [Wiki](https://github.com/golang/go/wiki) , Working with Go ,Learning more about Go
- [IDE, Edit](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins) , [Vim-go](https://github.com/fatih/vim-go)
- [Function Types in Go](http://jordanorelli.com/post/42369331748/function-types-in-go-golang)
- [Go examples](https://gobyexample.com/)

### Go cmd
go get [options] <package> , download packaget to $GOPATH with hierarchical path
    options:
            -d , download only                                                               
            ~                                   

### Print Variable

- %v , value => only property value
- %+v , plus value => with key and property value
- %#c , shaped value => go format
- %s , string
- %d , digit

### Go implicit rules 

- function name start with uppercase => public
- function name start with lowercase => private



### code syntax

##### type
```
cases := []struct { 
    in, want string    
}{
    {"Hello, world", "dlrow , olleH"},
    {"Hello, world", "dlrow , olleH"},
    ...
}
```



```
for \_,c := range cases {
    // _ = index    
    // c = value
    
}
```

### Go build
Now, test that the package compiles with go build:
`go build github.com/user/stringutil`
Or, if you are working in the package's source directory, just:

`go build`
This won't produce an output file.

### Go install
places the package object inside the pkg directory of the workspace.
If it's a executable program , place it inside the bin directory


### Go test
1. file name end with \_test.go
2. func name will be `func Test<Tested_func_name>(t *testing.T)`
3. if call `t.Fail` or `t.Error` , each means failed
