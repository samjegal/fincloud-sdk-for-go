package report

import (
	"github.com/samjegal/fincloud-sdk-for-go/tools/apidiff/delta"
	"github.com/samjegal/fincloud-sdk-for-go/tools/apidiff/exports"
)

type Package struct {
	AdditiveChanges *delta.Content   `json:"additiveChanges,omitempty"`
	BreakingChanges *BreakingChanges `json:"breakingChanges,omitempty"`
}

func (r Package) HasBreakingChanges() bool {
	return r.BreakingChanges != nil && !r.BreakingChanges.IsEmpty()
}

func (r Package) HasAdditiveChanges() bool {
	return r.AdditiveChanges != nil && !r.AdditiveChanges.IsEmpty()
}

func (r Package) IsEmpty() bool {
	return (r.AdditiveChanges == nil || r.AdditiveChanges.IsEmpty()) &&
		(r.BreakingChanges == nil || r.BreakingChanges.IsEmpty())
}

type BreakingChanges struct {
	Consts     map[string]delta.Signature    `json:"consts,omitempty"`
	Funcs      map[string]delta.FuncSig      `json:"funcs,omitempty"`
	Interfaces map[string]delta.InterfaceDef `json:"interfaces,omitempty"`
	Structs    map[string]delta.StructDef    `json:"structs,omitempty"`
	Removed    *delta.Content                `json:"removed,omitempty"`
}

func (bc BreakingChanges) IsEmpty() bool {
	return len(bc.Consts) == 0 && len(bc.Funcs) == 0 && len(bc.Interfaces) == 0 && len(bc.Structs) == 0 &&
		(bc.Removed == nil || bc.Removed.IsEmpty())
}

func Generate(lhs, rhs exports.Content, onlyBreakingChanges, onlyAdditions bool) Package {
	r := Package{}
	if !onlyBreakingChanges {
		if adds := delta.GetExports(lhs, rhs); !adds.IsEmpty() {
			r.AdditiveChanges = &adds
		}
	}

	if !onlyAdditions {
		breaks := BreakingChanges{}
		breaks.Consts = delta.GetConstTypeChanges(lhs, rhs)
		breaks.Funcs = delta.GetFuncSigChanges(lhs, rhs)
		breaks.Interfaces = delta.GetInterfaceMethodSigChanges(lhs, rhs)
		breaks.Structs = delta.GetStructFieldChanges(lhs, rhs)
		if removed := delta.GetExports(rhs, lhs); !removed.IsEmpty() {
			breaks.Removed = &removed
		}
		if !breaks.IsEmpty() {
			r.BreakingChanges = &breaks
		}
	}
	return r
}

func (r Package) ToMarkdown() string {
	if r.IsEmpty() {
		return ""
	}
	return formatAsMarkdown(r)
}
