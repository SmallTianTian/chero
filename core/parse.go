package core

import (
	"path/filepath"
	"strings"
)

func (sc *ScanFunc) parse() {
	sc.methodNamePath()
	sc.parseMethods()
}

func (sc *ScanFunc) parseMethods() {
	switch sc.F.Name.Name[:2] {
	case "Ge":
		sc.mwps = append(sc.mwps, &MethodWithPath{Method: GET, Path: sc.methodHttpPath})
	case "Po":
		sc.mwps = append(sc.mwps, &MethodWithPath{Method: POST, Path: sc.methodHttpPath})
	case "Pu":
		sc.mwps = append(sc.mwps, &MethodWithPath{Method: PUT, Path: sc.methodHttpPath})
	case "Pa":
		sc.mwps = append(sc.mwps, &MethodWithPath{Method: PATCH, Path: sc.methodHttpPath})
	case "De":
		sc.mwps = append(sc.mwps, &MethodWithPath{Method: DELETE, Path: sc.methodHttpPath})
	case "Op":
		sc.mwps = append(sc.mwps, &MethodWithPath{Method: OPTIONS, Path: sc.methodHttpPath})
	}

	if sc.F.Doc == nil {
		return
	}

	for _, comment := range sc.F.Doc.List {
		realComment := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))
		switch realComment {
		case "@HttpGet":
			sc.mwps = append(sc.mwps, &MethodWithPath{Method: GET, Path: sc.methodHttpPath})
		case "@HttpPost":
			sc.mwps = append(sc.mwps, &MethodWithPath{Method: POST, Path: sc.methodHttpPath})
		case "@HttpPut":
			sc.mwps = append(sc.mwps, &MethodWithPath{Method: PUT, Path: sc.methodHttpPath})
		case "@HttpPatch":
			sc.mwps = append(sc.mwps, &MethodWithPath{Method: PATCH, Path: sc.methodHttpPath})
		case "@HttpDelete":
			sc.mwps = append(sc.mwps, &MethodWithPath{Method: DELETE, Path: sc.methodHttpPath})
		case "@HttpOptions":
			sc.mwps = append(sc.mwps, &MethodWithPath{Method: OPTIONS, Path: sc.methodHttpPath})
		}
	}
}

func (sc *ScanFunc) methodNamePath() {
	funcName := sc.F.Name.Name
	switch funcName[:2] {
	case "Ge":
		funcName = strings.TrimLeft(funcName, "Get")
	case "Po":
		funcName = strings.TrimLeft(funcName, "Post")
	case "Pu":
		funcName = strings.TrimLeft(funcName, "Put")
	case "Pa":
		funcName = strings.TrimLeft(funcName, "Patch")
	case "De":
		funcName = strings.TrimLeft(funcName, "Delete")
	case "Op":
		funcName = strings.TrimLeft(funcName, "Options")
	}
	sc.methodHttpPath = UrlUnderscored(filepath.Join(sc.AbsPath, funcName))
}
