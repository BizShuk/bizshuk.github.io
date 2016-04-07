# Golang

### ref
- [Wiki](https://github.com/golang/go/wiki) , Working with Go ,Learning more about Go
- [IDE, Edit](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins) , [Vim-go](https://github.com/fatih/vim-go)


cases := []struct {
    
    in, want string    
}{
    {"Hello, world", "dlrow , olleH"},
    {"Hello, world", "dlrow , olleH"},
    ...
}





for \_,c := range cases {
    // _ = index    
    // c = value
    
}



### Go test
1. file name end with \_test.go
2. func name will be `func Test<Tested_func_name>(t *testing.T)`
3. if call `t.Fail` or `t.Error` , each means failed
