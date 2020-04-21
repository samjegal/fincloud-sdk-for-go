package exports

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

type Package struct {
	f     *token.FileSet
	p     *ast.Package
	files map[string][]byte
}

type LoadPackageErrorInfo interface {
	PkgCount() int
}

type errorInfo struct {
	m string
	c int
}

func (ei errorInfo) Error() string {
	return ei.m
}

func (ei errorInfo) PkgCount() int {
	return ei.c
}

func LoadPackage(dir string) (pkg Package, err error) {
	pkg.files = map[string][]byte{}
	pkg.f = token.NewFileSet()
	packages, err := parser.ParseDir(pkg.f, dir, nil, 0)
	if err != nil {
		return
	}
	if len(packages) < 1 {
		err = errorInfo{
			m: fmt.Sprintf("didn't find any package in '%s'", dir),
			c: len(packages),
		}
		return
	}
	if len(packages) > 1 {
		err = errorInfo{
			m: fmt.Sprintf("found more than one package in '%s'", dir),
			c: len(packages),
		}
		return
	}
	for pn := range packages {
		p := packages[pn]
		if exp := ast.PackageExports(p); !exp {
			err = fmt.Errorf("package '%s' doesn't contain any exports", pn)
			return
		}
		pkg.p = p
		return
	}
	panic("failed to return package")
}

func (pkg Package) GetExports() (c Content) {
	c = NewContent()
	ast.Inspect(pkg.p, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if t, ok := x.Type.(*ast.StructType); ok {
				c.addStruct(pkg, x.Name.Name, t)
			} else if t, ok := x.Type.(*ast.InterfaceType); ok {
				c.addInterface(pkg, x.Name.Name, t)
			}
		case *ast.FuncDecl:
			c.addFunc(pkg, x)
			return false
		case *ast.GenDecl:
			if x.Tok == token.CONST {
				c.addConst(pkg, x)
			}
		}
		return true
	})
	return
}

func (pkg Package) Name() string {
	return pkg.p.Name
}

func Get(pkgDir string) (Content, error) {
	pkg, err := LoadPackage(pkgDir)
	if err != nil {
		return Content{}, err
	}
	return pkg.GetExports(), nil
}

func (pkg Package) getText(start token.Pos, end token.Pos) string {
	p := pkg.f.Position(start)
	if _, ok := pkg.files[p.Filename]; !ok {
		b, err := ioutil.ReadFile(p.Filename)
		if err != nil {
			panic(err)
		}
		pkg.files[p.Filename] = b
	}
	return string(pkg.files[p.Filename][p.Offset : p.Offset+int(end-start)])
}

func (pkg Package) translateFieldList(fl []*ast.Field, cb func(*string, string)) {
	for _, f := range fl {
		var name *string
		if f.Names != nil {
			n := pkg.getText(f.Names[0].Pos(), f.Names[0].End())
			name = &n
		}
		t := pkg.getText(f.Type.Pos(), f.Type.End())
		cb(name, t)
	}
}

func (pkg Package) buildFunc(ft *ast.FuncType) (f Func) {
	appendString := func(s, a string) string {
		if s != "" {
			s += ","
		}
		s += a
		return s
	}
	if ft.Params.List != nil {
		p := ""
		pkg.translateFieldList(ft.Params.List, func(n *string, t string) {
			p = appendString(p, t)
		})
		f.Params = &p
	}
	if ft.Results != nil {
		r := ""
		pkg.translateFieldList(ft.Results.List, func(n *string, t string) {
			r = appendString(r, t)
		})
		f.Returns = &r
	}
	return
}
