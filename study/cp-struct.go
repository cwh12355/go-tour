package main

import f "fmt"

type persion struct {
	name  string
	age   int
	hobby string
	dead  bool
	time  int
}

type Cao struct {
	persion
	honor string
	time  int
	money int
	unit  string
}

func main() {
	caocao := Cao{persion{"CaoWenHao", 12, "running", false, 2032}, "bussiness", 2011, 100, "RMB WA"}
	f.Println(caocao.name)
	f.Println(caocao.time)

	f.Println(caocao.persion.hobby)

}
