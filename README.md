## HCL Sample

The code in this repository accompanies [this blog post](http://jen20.com/2015/09/07/using-hcl-part-1.html) about using HCL from Go.

Each part of the series corresponds to the code on the respective branch.

### How to use this example

```bash
go get github.com/jen20/hcl-sample
go get github.com/hashicorp/go-multierror
go get github.com/mitchellh/mapstructure
go get github.com/hashicorp/hcl
 
cd $GOPATH/src/github.com/jen20/hcl-sample/mapstructureusage/
go test
2015/12/20 22:13:32 mapstructureusage.Person{FirstName:"Frank", Surname:"Sinatra", City:"Hoboken", YearOfBirth:1915}
PASS
ok      github.com/jen20/hcl-sample/mapstructureusage   0.008s
```
