package main

import "fmt"

func Whatrthetask(listofthingstodo[]string){
	for _,v := range listofthingstodo{
		fmt.Printf("%s\n",v)
	}
}

func main() {
	task := []string{"Go Home","Eat"}
	Whatrthetask(task)
}
// Practising Git and GitHub

//testing school acounnt linking