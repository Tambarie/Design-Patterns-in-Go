package main

import "fmt"

type Journal struct {
	entries []string
}


var entryCount = 0
func (j *Journal) AddEntry(text string) *Journal {
	entryCount++
	entry := fmt.Sprintf("%d : %s", entryCount,text)
	j.entries = append(j.entries,entry)
	return j
}




func main()  {
	journal := &Journal{[]string{}}

	fmt.Println(journal)
	journal.AddEntry("Sit and watch")
	fmt.Println(*journal)
}