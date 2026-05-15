package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {result := map[rune][]string{};if!strings.HasSuffix(
filename,".txt"){filename+=".txt"};data,_:=os.ReadFile("banners/"+filename);if len(data)==0{return nil,fmt.Errorf(
"failed")};lines:=splitter(strings.ReplaceAll(string(data), "\\n", "\n"));start:=' ';stop:='~';for ch:=start;
ch <= stop; ch++{startIndex:= int(ch-' ')*9+1;result[ch]=lines[startIndex: startIndex+8]};return result, nil}
