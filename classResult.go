package main

import (
	"fmt"
	"strings"
)

func getHighestScorerPerSubject(students []Student) map[string]Student {
	highestScorers := make(map[string]Student)

	for _, subject := range subs {
		maxScore := -1
		var highestStudentPersubject Student
		for _, student := range students {
			for _, course := range student.Courses {
				if course.Name == subject && course.Score > maxScore {
					maxScore = course.Score
					highestStudentPersubject = student
				}
			}
		}

		highestScorers[subject] = highestStudentPersubject
	}

	return highestScorers
}

func main() {
	students := makeStudents()
	//for _, student := range students {
	//	fmt.Printf("Student: %s\n", student.Name)
	//	student.calculateStudentTotalScore()
	//	for _, course := range student.Courses {
	//		fmt.Printf("  %s: Score: %d, Grade: %s\n  ", course.Name, course.Score, course.Grade)
	//	}
	//	fmt.Println("-----------------------------")
	//}

	//--------------------------------

	for i := range students {
		students[i].calculateStudentTotalScore()
	}

	//newList := newListSortStudentsByTotalScore(students)
	//
	//fmt.Printf("Highest Scorer overall: %s\n", newList[0].Name)
	//for _, student := range newList {
	//	fmt.Printf("Name: %s, TotalScore: %d\n", student.Name, student.TotalScore)
	//
	//}

	//----------------------------------------subjectscore----------------------------------/

	newSubList := newListSortStudentsBySubjectScore(students, "english")

	fmt.Printf("Highest Scorer overall in Mathematics: %s\n", newSubList[0].Name)

	for _, student := range newSubList {
		var subjectScore int
		var courseName string
		for _, course := range student.Courses {
			if strings.ToLower(course.Name) == "english" {
				subjectScore = course.Score
				courseName = course.Name
				break
			}
		}
		fmt.Printf("Name: %s, %s Score: %d\n", student.Name, courseName, subjectScore)
	}

}
