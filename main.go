package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

func(s student) String() string {
	return fmt.Sprintf("%s %s from %s scored %d, %d, %d and %d marks in tests respectively", s.firstName, s.lastName, s.university, s.test1Score, s.test2Score, s.test3Score, s.test4Score)
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {
	scanner, cleanup, err := readFile(filePath)
	if err != nil {
		fmt.Printf("Error reading the file: %s", err.Error())
		return nil
	}
	defer cleanup()

	students, err := extractStudents(scanner)
	if err != nil {
		fmt.Printf("Error parsing the test score: %s", err.Error())
		return nil
	}
	
	return students
}

func readFile(filePath string) (*bufio.Scanner, func(), error) {
	file, err := os.Open(filePath)
	if err != nil {
		// fmt.Printf("Error reading the file: %s", err.Error())
		return nil, nil, err
	}
	
	cleanup := func() {
		file.Close()
	}

	return bufio.NewScanner(file), cleanup, nil
}

func extractStudents(scanner *bufio.Scanner) ([]student, error) {
	var students []student

	// discarding the header of csv file
	scanner.Scan()

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")

		test1Score, err := strconv.Atoi(line[3])
		if err != nil {
			// fmt.Printf("Error parsing the test score: %s", err.Error())
			return nil, err
		}
		test2Score, err := strconv.Atoi(line[4])
		if err != nil {
			// fmt.Printf("Error parsing the test score: %s", err.Error())
			return nil, err
		}
		test3Score, err := strconv.Atoi(line[5])
		if err != nil {
			// fmt.Printf("Error parsing the test score: %s", err.Error())
			return nil, err
		}
		test4Score, err := strconv.Atoi(line[6])
		if err != nil {
			// fmt.Printf("Error parsing the test score: %s", err.Error())
			return nil, err
		}

		students = append(students, student{
			firstName: line[0],
			lastName: line[1],
			university: line[2],
			test1Score: test1Score,
			test2Score: test2Score,
			test3Score: test3Score,
			test4Score: test4Score,
		})
	}

	return students, nil
}


func calculateGrade(students []student) []studentStat {
	var studentStats []studentStat

	for _, student := range students {
		var finalScore float32
		var grade Grade

		finalScore = float32(student.test1Score + student.test2Score + student.test3Score + student.test4Score)/4
		
		switch {
		case finalScore >= 70:
			grade = A
		case finalScore >= 50 && finalScore < 70:
			grade = B
		case finalScore >= 35 && finalScore < 50:
			grade = C
		default:
			grade = F
		}

		studentStats = append(studentStats, studentStat{
			student: student,
			finalScore: finalScore,
			grade: grade,
		})
	}

	return studentStats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	var topStudent studentStat

	for _, gradedStudent := range gradedStudents {
		if gradedStudent.finalScore > topStudent.finalScore {
			topStudent = gradedStudent
		}
	}
	
	return topStudent
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	topStudentPerUni := make(map[string]studentStat)

	for _, student := range gs {
		val, ok := topStudentPerUni[student.university]

		if ok {
			if student.finalScore > val.finalScore {
				topStudentPerUni[student.university] = student
			}
		} else {
			topStudentPerUni[student.university] = student
		}
	}

	return topStudentPerUni
}

// func parseCSV(filePath string) []student {
// 	var students []student

// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		fmt.Printf("Error reading the file: %s", err.Error())
// 		return nil
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	// scanner.Split(bufio.ScanLines)

// 	// discarding the header of csv file
// 	scanner.Scan()

// 	for scanner.Scan() {
// 		line := strings.Split(scanner.Text(), ",")

// 		test1Score, err := strconv.Atoi(line[3])
// 		if err != nil {
// 			fmt.Printf("Error parsing the test score: %s", err.Error())
// 			return nil
// 		}
// 		test2Score, err := strconv.Atoi(line[4])
// 		if err != nil {
// 			fmt.Printf("Error parsing the test score: %s", err.Error())
// 			return nil
// 		}
// 		test3Score, err := strconv.Atoi(line[5])
// 		if err != nil {
// 			fmt.Printf("Error parsing the test score: %s", err.Error())
// 			return nil
// 		}
// 		test4Score, err := strconv.Atoi(line[6])
// 		if err != nil {
// 			fmt.Printf("Error parsing the test score: %s", err.Error())
// 			return nil
// 		}

// 		students = append(students, student{
// 			firstName: line[0],
// 			lastName: line[1],
// 			university: line[2],
// 			test1Score: test1Score,
// 			test2Score: test2Score,
// 			test3Score: test3Score,
// 			test4Score: test4Score,
// 		})
// 	}

// 	return students
// }