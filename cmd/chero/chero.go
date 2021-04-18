package main

import (
	"crypto/md5"
	"fmt"
	"go/ast"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/SmallTianTian/chero/core"
	"golang.org/x/tools/go/packages"
)

func newCheros(pkg *packages.Package) []*chero {
	if len(pkg.GoFiles) == 0 {
		log.Fatalln("No any go file.")
	}
	// 没有导入这个包的，一定不需要自动生成代码
	if _, exist := pkg.Imports["github.com/SmallTianTian/chero"]; !exist {
		log.Fatalln("Not need chero compile.")
	}

	// 获取当前位置
	dir := filepath.Dir(pkg.GoFiles[0])

	var cheros []*chero
	// 循环每一个文件，创建 chero 待执行实例
	for _, eachFile := range pkg.Syntax {
		// 如果没有评论，则不会由当前文件发起 generated 请求
		if len(eachFile.Comments) == 0 {
			continue
		}

	Loop:
		for _, comment := range eachFile.Comments {
			for _, eachComment := range comment.List {
				// 找到发起 generated 请求的文件，加入待执行实例列表
				if strings.HasPrefix(eachComment.Text, "//go:generate chero") {
					cheros = append(cheros, &chero{
						path:    dir,
						astFile: eachFile,
						fresh:   &ast.File{},
					})
					break Loop
				}
			}
		}
	}
	if len(cheros) == 0 {
		log.Fatalln("Not need chero compile. Or no specile comments.")
	}
	return cheros
}

type chero struct {
	path    string
	astFile *ast.File
	fresh   *ast.File
}

// 查找需要进行 generated 的目录地址
func (c *chero) searchGeneratPath() []string {
	// 找到别名，没有别名，直接使用 chero
	var alias string
	for _, v := range c.astFile.Imports {
		if v.Path.Value == `"github.com/SmallTianTian/chero"` {
			if v.Name == nil {
				alias = "chero"
			} else {
				alias = v.Name.Name
			}
		}
	}

	// 找到调用 Scan 的地方
	var paths []string
	for _, v := range c.astFile.Scope.Objects {
		for _, line := range v.Decl.(*ast.FuncDecl).Body.List {
			var expr *ast.ExprStmt
			var call *ast.CallExpr
			var selector *ast.SelectorExpr
			var ident *ast.Ident
			var basic *ast.BasicLit
			var ok bool
			if expr, ok = line.(*ast.ExprStmt); !ok {
				continue
			}
			if call, ok = expr.X.(*ast.CallExpr); !ok {
				continue
			}
			if selector, ok = call.Fun.(*ast.SelectorExpr); !ok {
				continue
			}
			// 如果参数不为一个，并且方法不是 Scan 则跳过
			if len(call.Args) != 1 || selector.Sel.Name != "Scan" {
				continue
			}
			if ident, ok = selector.X.(*ast.Ident); !ok || ident.Name != alias {
				continue
			}
			if basic, ok = call.Args[0].(*ast.BasicLit); !ok {
				continue
			}
			paths = append(paths, strings.TrimRight(strings.TrimPrefix(basic.Value, `"`), `"`))
		}
	}
	return paths
}

// Render 渲染用户的代码。
func (c *chero) Render() error {
	if c.astFile == nil {
		log.Println("No such go file.")
		return nil
	}

	paths := c.searchGeneratPath()
	if len(paths) == 0 {
		log.Println("No any generat path.")
		return nil
	}
	for _, v := range paths {
		if len(v) > 0 && v[0] == os.PathSeparator {
			fmt.Println(v)
		} else {
			bs := md5.Sum([]byte(filepath.Join(c.path, v)))
			fileName := fmt.Sprintf("`%X.go`", bs)
			g := &core.CheroGen{
				Dir:      c.path,
				FileName: fileName,
				ScanDir:  filepath.Join(c.path, v),
			}
			g.TODO()
		}
	}
	return nil
}
