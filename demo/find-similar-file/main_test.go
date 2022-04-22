package main

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"strings"
	"testing"
)

func TestJieba(t *testing.T)  {
	// jieba
	x := gojieba.NewJieba()
	defer x.Free()

	s := "你喜欢吃水果吗"
	words := x.Cut(s, true)
	fmt.Println(s)
	fmt.Println("精确模式:", strings.Join(words, "/"))
}

func TestSliceIntersect(t *testing.T)  {
	s1 := []string{"苹果", "橘子", "香蕉"}
	s2 := []string{"苹果", "橘子", "橙子"}
	s3 := []string{"菠萝", "香蕉"}
	intersect := sliceIntersect(s1, s2, s3)
	fmt.Println(intersect)
}

func TestWalkFiles(t *testing.T)  {
	// jieba
	x := gojieba.NewJieba()
	defer x.Free()

	dirNames := []string{
		"./data/1",
		"./data/2",
	}
	sFiles := walkFiles(dirNames, x)
	for _, sF := range sFiles {
		fmt.Println(sF)
	}
}

func TestDiff(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < len(s1); i++ {
		for j := i+1; j < len(s1); j++ {
			fmt.Println(s1[i], s1[j])
		}
	}
}

func TestListFiles(t *testing.T)  {
	dirname := "./data"
	listFiles(dirname, 0)
}
