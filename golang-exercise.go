package main

import (
	"fmt"
	"strings"
)

type Student struct {
	No         int
	Name       string
	Exercise   float64
	Homework   float64
	Exam       float64
	FinalScore float64
	Grade      string
}

func calculateFinal(exercise, homework, exam float64) float64 {
	return (exercise * 0.2) + (homework * 0.3) + (exam * 0.5)
}

func grade(value float64) string {
	switch {
	case value >= 85:
		return "A"
	case value >= 75:
		return "B"
	case value >= 65:
		return "C"
	case value >= 50:
		return "D"
	default:
		return "E"
	}
}

func showReport(data []Student) {
	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("%-3s %-10s %-8s %-8s %-8s %-12s %-5s\n", "No", "Name", "Exercise", "Homework", "Exam", "Final Score", "Grade")
	fmt.Println(strings.Repeat("-", 70))
	for _, s := range data {
		fmt.Printf("%-3d %-10s %-8.1f %-8.1f %-8.1f %-12.1f %-5s\n", s.No, s.Name, s.Exercise, s.Homework, s.Exam, s.FinalScore, s.Grade)
	}
}

func (s *Student) calculateFinalScore() {
	s.FinalScore = calculateFinal(s.Exercise, s.Homework, s.Exam)
	s.Grade = grade(s.FinalScore)
}

func main() {
	studentList := []Student{
		{No: 1, Name: "Andi", Exercise: 85, Homework: 90, Exam: 80},
		{No: 2, Name: "Rara", Exercise: 90, Homework: 90, Exam: 95},
		{No: 3, Name: "Okta", Exercise: 78, Homework: 80, Exam: 87},
	}
	for i, _ := range studentList {
		// studentList[i].FinalScore = calculateFinal(studentList[i].Exercise, studentList[i].Homework, studentList[i].Exam)
		// studentList[i].Grade = grade(studentList[i].FinalScore)
		studentList[i].calculateFinalScore()
	}
	showReport(studentList)
}
