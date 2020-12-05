package main

import (
	"bufio"
    "fmt"
    "os"
    "io"
	"strings"
	"strconv"
	"log"
	"bytes"
)

func printDistr(array []int, text string) {
	
	fmt.Println(text)
	strArr := strings.Split(text,"")

	for i, c := range array {
		fmt.Println(strArr[0],i+1, "->", c)
	}
}

func assignTasks(tasks [][]string, nGroups []int, nTasks []int, ord string) []int {

	groupsAssignents := make([]int, len(nGroups))

	for tid, groups := range tasks {

		gid := -1;
		gcount := -1
		if ord == "min" {
			gcount = len(nGroups)+1
		}

		for _, group := range groups {
			// get group id
			arr := strings.Split(group, "")
			gdx, _ := strconv.Atoi(arr[1])


			condition := nGroups[gdx-1] > gcount

			if ord == "min" {
				condition = nGroups[gdx-1] < gcount
			}

			if condition {
				gid = gdx
				gcount = nGroups[gdx-1] 
			}
		}

		groupsAssignents[tid] = gid-1;

	}

	return groupsAssignents

}

func lineCounter(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
        log.Fatal(err)
	}
	defer file.Close()

	r := file

    buf := make([]byte, 32*1024)
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], lineSep)

        switch {
        case err == io.EOF:
            return count, nil

        case err != nil:
            return count, err
        }
	}
}

func main() {

	arg := os.Args[1:]
	filename := arg[0]
	citerion := arg[1]

	totalGroups := 0
	totalTasks,_ := lineCounter(filename)

	tasksData := make([][]string, totalTasks-1)

	file, err := os.Open(filename)
	if err != nil {
        log.Fatal(err)
	}
	defer file.Close()

	count := -1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if count == -1{
			totalGroups,_ = strconv.Atoi(scanner.Text())
		} else {
			tasksData[count] = strings.Split(scanner.Text(), ",")
		}
		count++
	}
	
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }


	nGroups := make([]int, totalGroups)
	nTasks := make([]int, totalTasks-1)

	for tid, groups := range tasksData {
		for _, group := range groups {
			arr := strings.Split(group,"")
			idx, _ := strconv.Atoi(arr[1])
			nGroups[idx-1] += 1
			nTasks[tid] += 1
		}
	}

	tasksAssignment := assignTasks(tasksData, nGroups, nTasks, citerion)

	for tid, gid := range tasksAssignment {
		fmt.Printf("T%d -> G%d\n", tid+1, gid+1)
	}

}