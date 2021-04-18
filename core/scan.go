package core

import (
	"go/ast"
	"os"

	"log"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

type MethodWithPath struct {
	Method HttpMethod
	Path   string
}

type ScanFunc struct {
	AbsPath        string
	PkgPath        string
	F              *ast.FuncDecl
	methodHttpPath string
	mwps           []*MethodWithPath
}

func ScanHttpMethodFunc(dir string) (result []*ScanFunc) {
	// fix dir is `.`, cloudn't load child dir.
	if len(dir) > 0 && dir[0] != os.PathSeparator {
		dir = "./" + dir
	}
	return scanMethodFunc(dir, dir)
}

func scanMethodFunc(dir, baseDir string) (result []*ScanFunc) {
	pkgs, err := packages.Load(&packages.Config{Mode: packages.NeedTypes | packages.NeedName | packages.NeedFiles | packages.NeedModule | packages.NeedImports | packages.NeedSyntax}, dir)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	pkg := pkgs[0]

	// 循环每一个 Go ast 文件
	for _, v := range pkg.Syntax {
		// 循环 go ast 文件的主体部分
		for _, fc := range v.Scope.Objects {
			f, ok := fc.Decl.(*ast.FuncDecl)
			// 如果不是 func 则直接跳过
			if !ok {
				continue
			}

			// 如果方法是 http 请求方法开头，将被加入扫描结果列表
			switch {
			case strings.HasPrefix(f.Name.Name, "Get"):
				fallthrough
			case strings.HasPrefix(f.Name.Name, "Post"):
				fallthrough
			case strings.HasPrefix(f.Name.Name, "Put"):
				fallthrough
			case strings.HasPrefix(f.Name.Name, "Patch"):
				fallthrough
			case strings.HasPrefix(f.Name.Name, "Delete"):
				fallthrough
			case strings.HasPrefix(f.Name.Name, "Options"):
				result = append(result, &ScanFunc{
					AbsPath: AbsHttpPath(dir, baseDir),
					PkgPath: pkg.PkgPath,
					F:       f,
				})
				continue
			}

			if f.Doc != nil {
			Loop:
				// 如果方法的注释中包含指定的 http 请描述，也将被加入扫描结果列表
				for _, comment := range f.Doc.List {
					realComment := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))
					switch realComment {
					case "@HttpGet", "@HttpPost", "@HttpPut",
						"@HttpPatch", "@HttpDelete", "@HttpOptions":
						result = append(result, &ScanFunc{
							AbsPath: AbsHttpPath(dir, baseDir),
							PkgPath: pkg.PkgPath,
							F:       f,
						})
						break Loop
					}
				}
			}
		}
	}

	// 读取当前目录所有文件
	dirFiles, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}
	for _, df := range dirFiles {
		// 如果还有子文件夹，再依次加载子文件夹
		if df.IsDir() {
			childDir := filepath.Join(dir, df.Name())
			childResult := scanMethodFunc(childDir, baseDir)
			result = append(result, childResult...)
		}
	}
	return
}
