package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		dir   string
		topic string
		num   int
	)

	flag.StringVar(&dir, "f", "../problems", "生成的目录")
	flag.IntVar(&num, "n", 0, "题号")
	flag.StringVar(&topic, "t", "code", "题目")

	flag.Parse()
	path := getPath(dir, topic, num)

	Create(path)
}

func getPath(f, t string, n int) string {
	return fmt.Sprintf("%s/%s.%s", f, qnum(n), t)
}

func qnum(n int) string {
	return fmt.Sprintf("%04d", n)
}

var (
	err error
)

func Create(path string) {

	if of, oerr := os.Stat(path); oerr == nil && of.IsDir() {
		log.Printf("%s is dir", path)
		return
	} else if os.IsExist(oerr) {
		log.Printf("%s is exist", path)
		return
	}

	log.Printf("path is %s", path)
	ps := strings.Split(path, "/")
	if len(ps) != 3 {
		return
	}

	rp := ps[2]
	pnames := strings.Split(rp, ".")[len(strings.Split(rp, "."))-1]
	pname := strings.Split(pnames, "-")[len(strings.Split(rp, "-"))-1]

	funcn := "Run"

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	fname := fmt.Sprintf("%s/run.go", path)

	_, err = os.Create(fname)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	rf, err := os.OpenFile(
		fname,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666)
	if err != nil {
		log.Fatal(err)
	}
	defer rf.Close()

	byteF := []byte(fmt.Sprintf(runs, pname, funcn))
	_, err = rf.Write(byteF)
	if err != nil {
		log.Fatal(err)
	}

	tname := fmt.Sprintf("%s/run_test.go", path)
	err = ioutil.WriteFile(tname, []byte(fmt.Sprintf(runt, pname, funcn, funcn)), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	rname := fmt.Sprintf("%s/README.md", path)

	err = ioutil.WriteFile(rname, []byte(fmt.Sprintf(readme, rp)), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

var runs = `
package %s

func %s() {

}
`
var runt = `
package %s

import "testing"

func Test%s(t *testing.T) {
	%s()
}
`

var readme = `
# %s

`
