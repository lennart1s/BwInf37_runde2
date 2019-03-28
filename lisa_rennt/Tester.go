package main

import "BwInf37_runde2/lisa_rennt/lib"

func main2() {
	v := &lib.Vertex{1, 2}
	v2 := v
	println(v == v2)
}
