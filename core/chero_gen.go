package core

import "fmt"

type CheroGen struct {
	Dir      string
	FileName string
	ScanDir  string
}

func (cg *CheroGen) TODO() {
	shmfs := ScanHttpMethodFunc(cg.Dir)
	for _, v := range shmfs {
		v.parse()
		for _, c := range v.mwps {
			fmt.Println(c)
		}
		fmt.Println("====")
	}
	cg.flush()
}

func (cg *CheroGen) flush() {
}
