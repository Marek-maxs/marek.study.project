package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// ScoreFunction 定义评分函数类型
type ScoreFunction func(*ast.File) float64

// scoreByLines 根据代码行数评分
func scoreByLines(file *ast.File) float64 {
	return float64(len(file.Decls)) / 100.0
}

// scoreByComments 根据注释比例评分
func scoreByComments(file *ast.File) float64 {
	var commentCount int
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CommentGroup); ok {
			commentCount += len(c.List)
		}
		return true
	})
	return float64(commentCount) / float64(len(file.Decls)+commentCount) * 100.0
}

// ScoreFile 评分文件
func ScoreFile(filename string, scoreFuncs ...ScoreFunction) (float64, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return 0, err
	}

	var totalScore float64
	for _, scoreFunc := range scoreFuncs {
		totalScore += scoreFunc(file)
	}

	return totalScore / float64(len(scoreFuncs)), nil
}

func ar() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: codescore <file.go>")
		return
	}

	filename := os.Args[1]
	score, err := ScoreFile(filename, scoreByLines, scoreByComments)
	if err != nil {
		fmt.Printf("Error scoring file: %v\n", err)
		return
	}

	fmt.Printf("Score for %s: %.2f\n", filename, score)
}