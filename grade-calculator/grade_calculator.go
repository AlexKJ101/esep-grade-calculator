package esepunittests

type GradeCalculator struct {
	grades []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func filterGradesByType(all []Grade, t GradeType) []Grade {
	filtered := make([]Grade, 0)
	for _, g := range all {
		if g.Type == t {
			filtered = append(filtered, g)
		}
	}
	return filtered
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignmentGrades := filterGradesByType(gc.grades, Assignment)
	examGrades := filterGradesByType(gc.grades, Exam)
	essayGrades := filterGradesByType(gc.grades, Essay)

	assignmentAverage := computeAverage(assignmentGrades)
	examAverage := computeAverage(examGrades)
	essayAverage := computeAverage(essayGrades)

	weightedGrade := float64(assignmentAverage)*0.5 + float64(examAverage)*0.35 + float64(essayAverage)*0.15

	return int(weightedGrade)
}

func computeAverage(grades []Grade) int {
	sum := 0

	for _, g := range grades {
		sum += g.Grade
	}

	return sum / len(grades)
}
