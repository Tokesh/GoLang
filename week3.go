package main

// Student schedule and realization

import (
	"fmt"
	"time"
)

type School struct {
	SchoolAdding
	Schedule
}

type Lesson struct {
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Teacher   Teacher
}

type Schedule struct {
	UniqueID int
	Lessons  []Lesson
	Pupils   []Pupil
	Teacher  Teacher
}

type SchoolAdding struct {
	PupilsNumber   int
	TeachersNumber int
	Pupils         []Pupil
	Teachers       []Teacher
}

type Pupil struct {
	Name string
	Age  int
	Id   string
}

type Teacher struct {
	Name        string
	Age         int
	PhoneNumber string
}

type teacherInterface interface {
	addTeacher(Teacher)
	removeTeacher(Teacher)
}

func (school *SchoolAdding) addPupil(pupil Pupil) {
	school.PupilsNumber += 1
	school.Pupils = append(school.Pupils, pupil)
}

func (school *SchoolAdding) removePupil(pupil Pupil) {
	for i, a := range school.Pupils {
		if a.Id == pupil.Id {
			//fmt.Println(a.Id)
			school.PupilsNumber -= 1
			school.Pupils = append(school.Pupils[:i], school.Pupils[i+1:]...)
		}
	}
}

func (school *SchoolAdding) addTeacher(teacher Teacher) {
	school.Teachers = append(school.Teachers, teacher)
	school.TeachersNumber += 1
}

func (school *SchoolAdding) removeTeacher(teacher Teacher) {
	for i, a := range school.Teachers {
		if a.Name == teacher.Name {
			//fmt.Println(a.Name)
			school.TeachersNumber -= 1
			school.Teachers = append(school.Teachers[:i], school.Teachers[i+1:]...)
		}
	}
}

func main() {
	Tokesh := Pupil{"Tokesh", 19, "20B030106"}
	Mels := Teacher{"Mels", 32, "T030245"}
	//School35 := SchoolAdding{0, 0, []Pupil{}, []Teacher{}}
	//NewSchoool := School{School35}
	School35 := School{}

	School35.addTeacher(Mels)
	School35.addPupil(Tokesh)
	fmt.Println(School35)
	School35.removePupil(Tokesh)
	School35.removeTeacher(Mels)

	fmt.Println(School35)
}
