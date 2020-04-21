package exports

import (
	"go/ast"
	"strings"
)

type Content struct {
	Consts     map[string]Const     `json:"consts,omitempty"`
	Funcs      map[string]Func      `json:"funcs,omitempty"`
	Interfaces map[string]Interface `json:"interfaces,omitempty"`
	Structs    map[string]Struct    `json:"structs,omitempty"`
}

type Const struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Func struct {
	Params  *string `json:"params,omitempty"`
	Returns *string `json:"returns,omitempty"`
}

type Interface struct {
	Methods map[string]Func
}

type Struct struct {
	AnonymousFields []string          `json:"anon,omitempty"`
	Fields          map[string]string `json:"fields,omitempty"`
}

func NewContent() Content {
	return Content{
		Consts:     make(map[string]Const),
		Funcs:      make(map[string]Func),
		Interfaces: make(map[string]Interface),
		Structs:    make(map[string]Struct),
	}
}

func (c Content) IsEmpty() bool {
	return len(c.Consts) == 0 && len(c.Funcs) == 0 && len(c.Interfaces) == 0 && len(c.Structs) == 0
}

func (c *Content) addConst(pkg Package, g *ast.GenDecl) {
	for _, s := range g.Specs {
		co := Const{}
		vs := s.(*ast.ValueSpec)
		v := ""
		if vs.Type != nil {
			co.Type = vs.Type.(*ast.Ident).Name
			v = vs.Values[0].(*ast.BasicLit).Value
		} else {
			if bl, ok := vs.Values[0].(*ast.BasicLit); ok {
				co.Type = strings.ToLower(bl.Kind.String())
				v = bl.Value
			} else if ce, ok := vs.Values[0].(*ast.CallExpr); ok {
				co.Type = pkg.getText(ce.Fun.Pos(), ce.Fun.End())
				v = pkg.getText(ce.Args[0].Pos(), ce.Args[0].End())
			} else {
				panic("unhandled case for adding constant")
			}
		}
		if v[0] == '"' {
			v = v[1 : len(v)-1]
		}
		co.Value = v
		c.Consts[vs.Names[0].Name] = co
	}
}

func (c *Content) addFunc(pkg Package, f *ast.FuncDecl) {
	sig := ""
	if f.Recv != nil {
		sig = pkg.getText(f.Recv.List[0].Type.Pos(), f.Recv.List[0].Type.End())
		sig += "."
	}
	sig += f.Name.Name
	c.Funcs[sig] = pkg.buildFunc(f.Type)
}

func (c *Content) addInterface(pkg Package, name string, i *ast.InterfaceType) {
	in := Interface{
		Methods: map[string]Func{},
	}
	if i.Methods != nil {
		for _, m := range i.Methods.List {
			n := m.Names[0].Name
			f := pkg.buildFunc(m.Type.(*ast.FuncType))
			in.Methods[n] = f
		}
	}
	c.Interfaces[name] = in
}

func (c *Content) addStruct(pkg Package, name string, s *ast.StructType) {
	sd := Struct{}
	pkg.translateFieldList(s.Fields.List, func(n *string, t string) {
		if n == nil {
			sd.AnonymousFields = append(sd.AnonymousFields, t)
		} else {
			if sd.Fields == nil {
				sd.Fields = map[string]string{}
			}
			sd.Fields[*n] = t
		}
	})
	c.Structs[name] = sd
}
