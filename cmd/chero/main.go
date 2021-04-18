package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/SmallTianTian/chero/core"
	"golang.org/x/tools/go/packages"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("chero: ")

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	// 获取待执行目录地址
	var dir string
	if len(args) == 1 && core.IsDir(args[0]) {
		dir = args[0]
	} else {
		dir = filepath.Dir(args[0])
	}
	if !core.IsExist(dir) {
		log.Fatalln("No such path.", dir)
	}

	// 加载文件夹
	pkgs, err := packages.Load(&packages.Config{Mode: packages.NeedFiles | packages.NeedModule | packages.NeedImports | packages.NeedSyntax}, dir)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	pkg := pkgs[0]

	for _, v := range newCheros(pkg) {
		if err := v.Render(); err != nil {
			log.Println("ERR:", err)
		}
	}
	// fmt.Println(newCheros(pkg)[0].astFile)
}
