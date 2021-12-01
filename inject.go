package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	regTag       = regexp.MustCompile("\\s*`*(.*?):\\s*(\".*?\")")
	rInject      = regexp.MustCompile("`.+`$")
	regInjectTag = regexp.MustCompile(`#(.*?):\s*(".*?")`)

	regInjectTag_Key   = 1
	regInjectTag_Value = 2
)

func Parse(inputPattern string) {
	paths, err := filepath.Glob(inputPattern)
	if err != nil {
		log.Fatal(err)
	}

	for _, path := range paths {
		fstat, err := os.Stat(path)
		if err != nil {
			log.Fatal(err)
		}

		if fstat.IsDir() {
			continue
		}

		if !strings.HasSuffix(strings.ToLower(fstat.Name()), ".go") {
			fmt.Println("skip", fstat.Name())
			continue
		}

		tagInfos, err := getTagPosition(path)
		if err != nil {
			log.Fatal(err)
		}

		injectTags(path, tagInfos)
	}
}

func getTagPosition(inputPath string) ([]TagInfo, error) {
	tagInfos := []TagInfo{}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, inputPath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for _, decl := range f.Decls {
		// 0: check if declaration is not function decl
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		var typeSpec *ast.TypeSpec
		for _, spec := range genDecl.Specs {
			if typeSpec, ok = spec.(*ast.TypeSpec); ok {
				break
			}
		}

		// 1: check if we can get the FIRST typeSpec
		if typeSpec == nil {
			continue
		}

		// 2: check if type spec define STRUCT
		structSpec, ok := typeSpec.Type.(*ast.StructType)
		if !ok || structSpec.Incomplete {
			continue
		}

		// 3: Get comment of all fields of struct
		for _, field := range structSpec.Fields.List {

			comment := ""
			if field.Doc != nil {
				txt := field.Doc.Text()
				comment += txt
			}

			if field.Comment != nil {
				txt := field.Comment.Text()
				comment += txt
			}

			tagInfo := TagInfo{
				Start: int(field.Pos()),
				End:   int(field.End()),
			}

			// convert tag to []KeyValuePair
			if field.Tag != nil {
				allTags := regTag.FindAllStringSubmatch(field.Tag.Value, -1)
				if len(allTags) > 0 {
					for _, v := range allTags {
						tagInfo.CurrentTag = append(tagInfo.CurrentTag, StringPair{
							Key:   v[regInjectTag_Key],
							Value: v[regInjectTag_Value],
						})
					}
				}
			}

			// convert comment to []KeyValuePair
			if len(comment) > 0 {
				injTag := regInjectTag.FindAllStringSubmatch(comment, -1)
				if len(injTag) > 0 {
					for _, v := range injTag {
						tagInfo.InjectTag = append(tagInfo.InjectTag, StringPair{
							Key:   v[regInjectTag_Key],
							Value: v[regInjectTag_Value],
						})
					}

					tagInfos = append(tagInfos, tagInfo)
				}
			}
		}
	}

	return tagInfos, nil
}

func injectTags(path string, tagInfos []TagInfo) error {
	oldF, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if len(tagInfos) == 0 {
		return nil
	}

	var newF []byte
	for i := len(tagInfos) - 1; i >= 0; i-- {
		newF = []byte{}

		content := make([]byte, tagInfos[i].End-tagInfos[i].Start)
		copy(content, oldF[tagInfos[i].Start-1:tagInfos[i].End-1])

		tag := rInject.ReplaceAll(content, []byte(fmt.Sprintf("`%s`", tagInfos[i].ToTag())))
		newF = append(newF, oldF[:tagInfos[i].Start-1]...)
		newF = append(newF, tag...)
		newF = append(newF, oldF[tagInfos[i].End-1:]...)

		oldF = newF
	}

	if err = ioutil.WriteFile(path, newF, 0644); err != nil {
		return err
	}

	return nil
}
