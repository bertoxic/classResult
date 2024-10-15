package main

import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

type Grade string

const (
	APlus Grade = "A+"
	A     Grade = "A"
	BPlus Grade = "B+"
	B     Grade = "B"
	CPlus Grade = "C+"
	C     Grade = "C"
	D     Grade = "D"
	F     Grade = "F"
)

type Subject struct {
	Name     string
	PassMark int
	Score    int
	Grade    Grade
}

type Student struct {
	Name       string
	Courses    []Subject
	TotalScore int
}

func (s *Subject) getGrade() {
	switch {
	case s.Score >= 90 && s.Score <= 100:
		s.Grade = APlus
	case s.Score >= 80 && s.Score < 90:
		s.Grade = A
	case s.Score >= 75 && s.Score < 80:
		s.Grade = BPlus
	case s.Score >= 70 && s.Score < 75:
		s.Grade = B
	case s.Score >= 65 && s.Score < 70:
		s.Grade = CPlus
	case s.Score >= 60 && s.Score < 65:
		s.Grade = C
	case s.Score >= 50 && s.Score < 60:
		s.Grade = D
	default:
		s.Grade = F
	}
}

var subs = []string{
	"Mathematics",
	"Physics",
	"Chemistry",
	"Biology",
	"History",
	"Geography",
	"English",
	"Computer Science",
	"Economics",
	"Physical Education",
	"Verbal Reasoning",
	"Literature",
	"French",
	"Fulani",
	"Basic Science",
	"vocational aptitude",
	"Cultural and Creative Education",
}

func makeStudents() []Student {
	studentNames := []string{"Alice", "Bob", "Charlie", "David", "Eva", "Frank", "Grace", "Helen", "Ivy", "John"}
	var students []Student

	for _, name := range studentNames {
		subjects := createSubjects()
		assignRandomScores(subjects)
		students = append(students, Student{Name: name, Courses: subjects})
	}

	return students
}

func createSubjects() []Subject {
	var subjects []Subject
	for _, subName := range subs {
		subjects = append(subjects, Subject{Name: subName, PassMark: 50})
	}
	return subjects
}

func assignRandomScores(subjects []Subject) {

	for i := range subjects {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
		subjects[i].Score = r.Intn(91) + 10
		subjects[i].getGrade()
	}
}
func (s *Student) calculateStudentTotalScore() {
	total := 0
	for _, subject := range s.Courses {
		total += subject.Score
	}
	s.TotalScore = total
}

func sortStudentsByTotalScore(students []Student) {
	sort.Slice(students, func(a, b int) bool {
		return students[a].TotalScore > students[b].TotalScore
	})
}

func newListSortStudentsByTotalScore(students []Student) []Student {
	newStudentList := make([]Student, len(students))
	copy(newStudentList, students)

	sort.Slice(newStudentList, func(a, b int) bool {
		return newStudentList[a].TotalScore > newStudentList[b].TotalScore
	})
	return newStudentList
}

func newListSortStudentsBySubjectScore(students []Student, subjectName string) []Student {
	newStudentList := make([]Student, len(students))
	copy(newStudentList, students)

	sort.Slice(newStudentList, func(a, b int) bool {
		var scoreA, scoreB int
		for _, course := range newStudentList[a].Courses {
			if strings.ToLower(course.Name) == strings.ToLower(subjectName) {
				scoreA = course.Score
				break
			}
		}
		for _, course := range newStudentList[b].Courses {
			if strings.ToLower(course.Name) == strings.ToLower(subjectName) {
				scoreB = course.Score
				break
			}
		}
		return scoreA > scoreB
	})

	return newStudentList
}
