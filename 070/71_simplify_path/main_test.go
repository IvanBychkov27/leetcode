package main

import (
	"fmt"
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	type data struct {
		Path     string
		Expected string
	}

	ds := []data{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
		{"/a//b////c/d//././/..", "/a/b/c"},
	}

	for _, d := range ds {
		res := simplifyPath(d.Path)
		if res != d.Expected {
			fmt.Println("error: path =", d.Path)
			fmt.Println("result  :", res)
			fmt.Println("expected:", d.Expected)

		}
	}

}
