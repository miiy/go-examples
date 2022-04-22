package main

import (
	"flag"
	"fmt"
	"github.com/yanyiwu/gojieba"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type sFile struct {
	path string
	fileInfo fs.FileInfo
	words []string
}

type flagParam struct {
	path string
	listFiles bool
}

func main() {
	// flag
	var p flagParam

	flag.StringVar(&p.path, "path", "", "--path=path1||path2")
	flag.BoolVar(&p.listFiles, "list-files", false, "--list-files=true")
	flag.Parse()
	if p.path == "" {
		fmt.Println("please input path")
		return
	}
	paths:= strings.Split(p.path, "||")
	log.Println(paths)

	if p.listFiles {
		for _, path := range paths {
			listFiles(path, 0)
		}
	}

	// jieba
	x := gojieba.NewJieba()
	defer x.Free()
	// walk
	sFiles := walkFiles(paths, x)
	// same check
	sFGS := sFileGroups(sFiles)
	//fmt.Println(sFGS)

	for i, group := range sFGS {
		fmt.Println(i)
		for _, f := range group {
			fmt.Println(f.path)
		}
	}
}

func walkFiles(dirNames []string, x *gojieba.Jieba) []sFile {
	var files []sFile
	for _, dirName :=range dirNames {
		err := filepath.Walk(dirName, func(path string, info fs.FileInfo, err error) error {
			fileName := info.Name()
			files = append(files, sFile{
				path: path,
				fileInfo: info,
				words: x.Cut(fileName, true),
			})
			return nil
		})
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return files
}

func sliceIntersect(s1 []string, s2 ...[]string) []string {
	kMap := make(map[string]int, len(s1))

	s2 = append(s2, s1)

	for _, arg := range s2 {
		for _ , sv := range arg {
			kMap[sv] += 1
		}
	}

	var res []string
	for k, v := range kMap {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}

func sFileGroups(sFiles []sFile) [][]sFile {
	var groups [][]sFile

	for i := 0; i < len(sFiles); i++ {
		var group []sFile
		group = append(group, sFiles[i])
		for j := i+1; j < len(sFiles); j++ {
			if sFiles[i].path == sFiles[j].path {
				continue
			}
			// similarity
			intersect := sliceIntersect(sFiles[i].words, sFiles[j].words)
			if len(intersect) > (len(sFiles[i].words) + len(sFiles[j].words)) / 3 {
				group = append(group, sFiles[j])
			}
		}

		if len(group) >= 2 {
			groups = append(groups, group)
		}
	}

	return groups
}

func listFiles(dirname string, level int) {
	// level用来记录当前递归的层次
	// 生成有层次感的空格
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|   " + s
	}

	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil{
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := dirname + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir() {
			//继续遍历fi这个目录
			listFiles(filename, level+1)
		}
	}
}
