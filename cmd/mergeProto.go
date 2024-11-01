package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

const protoHeader = `syntax = "proto3";
package core;
option go_package = "./core";`

func main() {
	baseDir := flag.String("base", "./rpc/desc/common", "base dir")
	dir := flag.String("dir", "./rpc/desc", ".proto dir")
	outDir := flag.String("out", "./rpc/desc", "all.proto out dir")
	serviceName := flag.String("n", "Core", "service name")
	flag.Parse()

	allPaths := getAllFilePaths(*baseDir)
	allPaths = append(allPaths, getAllFilePaths(*dir)...)

	genAllProto(*serviceName, *outDir, allPaths)
}

func getAllFilePaths(dir string) []string {
	matchs, err := filepath.Glob(path.Join(dir, "*.proto"))
	if err != nil {
		panic(err)
	}

	for _, match := range matchs {
		info, err := os.Stat(match)
		if err != nil {
			panic(err)
		}
		if info.IsDir() {
			matchs = append(matchs, getAllFilePaths(match)...)
		}
	}

	return matchs
}

func genAllProto(sn string, outDir string, allPaths []string) {
	outFile, err := os.Create(outDir)
	serviceReg := regexp.MustCompile(fmt.Sprintf(`(?ms)service\s+%s\s+\{(.*)\}`, sn))
	serviceSb := strings.Builder{}

	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	buf := make([]byte, 1024)
	defer writer.Flush()
	writer.WriteString(protoHeader)

	for _, path := range allPaths {
		protoFile, err := os.Open(path)
		sb := strings.Builder{}
		if err != nil {
			panic(err)
		}
		defer protoFile.Close()

		reader := bufio.NewReader(protoFile)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					fmt.Println("gen proto err:\n", err)
				}
				break
			}

			if isBegin(line) {
				sb.WriteString(line)
				io.CopyBuffer(&sb, reader, buf)
				str := sb.String()

				tmpStr := serviceReg.FindStringSubmatch(str)
				if len(tmpStr) > 1 {
					serviceSb.WriteString(tmpStr[1])
					str = serviceReg.ReplaceAllString(str, "")
				}

				writer.WriteString(str)
				break
			}
		}
	}

	writer.WriteString(fmt.Sprintf("service %s {\n%s\n}", sn, serviceSb.String()))
}

func isBegin(txt string) bool {
	return !(strings.HasPrefix(txt, "syntax") || strings.HasPrefix(txt, "package") || strings.HasPrefix(txt, "option"))
}
