package main

import (
	"fmt"
)

func main() {
	name := "Santosh"
	course := "Docker"

	fmt.Println("\nHi,", name, "you are currently watching", course, "course")

	changeCourse(course) //pass by value
	fmt.Println("\nYou are watching course", course)

	changeCourseRef(&course) //pass by reference
	fmt.Println("\nYou are watching course", course)
}

func changeCourse(course string) {
	course = "First Look: Native Docker Clustering"
	fmt.Println("\nTrying to change your course to", course)

}

func changeCourseRef(course *string) {
	*course = "First Look: Native Docker Clustering"
	fmt.Println("\nTrying to change your course to", *course)
}
