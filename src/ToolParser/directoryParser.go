// Package toolParser ,package for dealing with parsing and testing of files
//
// file for dealing with parsing of directory
package toolParser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func getGoFileNode(path string) (*ast.File, error) {
	if filepath.Ext(path) == ".go" {
		fset := token.NewFileSet()

		node, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return nil, err
		}
		return node, nil
	}
	return nil, nil
}

// ParseDir function
//
// Parses given directory and
func ParseDir(dir string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		node, err := getGoFileNode(path)
		if node != nil {
			packageName := getPackageName(node)
			packageInfo := getMapPackage(packageName)
			newFile := packageFiles{
				Name: filepath.Base(path),
			}
			packageInfo.Files = append(packageInfo.Files, &newFile)
			parseCode(packageInfo, node)

		}
		return err
	})
	if err != nil {
		return err
	}
	return nil
}
