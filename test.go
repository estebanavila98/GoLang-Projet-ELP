package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func check(e error) {
	if e != nil {
	panic(e)
	}
}
func main(){
	name := os.Args[1]

	file,err:= os.Open(name)
	check(err)

	rd := bufio.NewReader(file)

	var lien []string

	for {
		lines, err := rd.ReadString('\n')
		if err != nil {
			break
		}
		line := strings.TrimSuffix(lines, "\r\n")
		line = strings.TrimSuffix(line, "\n")

		lien = append(lien,strings.Split(line, " ")...)
	}

	var tab []int
	fmt.Printf("%v",tab)
	for i := 0 ; i<len(lien)-1 ; i ++ {
		tab[i], err = strconv.Atoi(lien[i])
	}
	fmt.Printf("%v",tab)


}
