package main

import (
	"fmt"
	"net/http"
	"strings"
)

//func getHighestScorerPerSubject(students []Student) map[string]Student {
//	highestScorers := make(map[string]Student)
//
//	for _, subject := range subs {
//		maxScore := -1
//		var highestStudentPersubject Student
//		for _, student := range students {
//			for _, course := range student.Courses {
//				if course.Name == subject && course.Score > maxScore {
//					maxScore = course.Score
//					highestStudentPersubject = student
//				}
//			}
//		}
//
//		highestScorers[subject] = highestStudentPersubject
//	}
//
//	return highestScorers
//}

func DisplayResultsForScores(w http.ResponseWriter, r *http.Request) {
	students := makeStudents()
	for _, student := range students {
		fmt.Fprintf(w, "Student: %s\n", student.Name)
		student.calculateStudentTotalScore()
		for _, course := range student.Courses {

			fmt.Fprintf(w, "  %s: Score: %d, Grade: %s\n  ", course.Name, course.Score, course.Grade)
			fmt.Fprintf(w, "\n")
		}
	}
}

func HandlerResultForHighestScorer(w http.ResponseWriter, r *http.Request) {
	students := makeStudents()

	for i := range students {

		students[i].calculateStudentTotalScore()
	}

	newList := newListSortStudentsByTotalScore(students)

	fmt.Fprintf(w, "Highest Scorer overall: %s", newList[0].Name)
	for _, student := range newList {
		fmt.Fprintf(w, "\n")

		fmt.Fprintf(w, "Name: %s, TotalScore: %d", student.Name, student.TotalScore)

	}

}

func DisplayResultsForSortBySubject(w http.ResponseWriter, r *http.Request) {
	students := makeStudents()
	newSubList := newListSortStudentsBySubjectScore(students, "english")

	fmt.Fprintf(w, "Highest Scorer overall in english: ", newSubList[0].Name)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, " %s\n", newSubList[0].Name)
	fmt.Fprintf(w, "\n")
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

		fmt.Fprintf(w, "\n")

		_, err := fmt.Fprintf(w, fmt.Sprintf("Name: %s, %s Score: %d\n", student.Name, courseName, subjectScore))
		if err != nil {
			return
		}

	}
}
func main() {

	http.HandleFunc("/", DisplayResultsForScores)
	http.HandleFunc("/sub", HandlerResultForHighestScorer)
	http.HandleFunc("/top", DisplayResultsForSortBySubject)
	// Start the server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}
