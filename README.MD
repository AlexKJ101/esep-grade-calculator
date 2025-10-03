# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```
Project Layout
- "grade_calculator.go" - Library implementation (GradeCalculator, Grade/GradeType, weighting logic).
- "grade_calculator_test.go" - Unit tests for the letter grade output.
- "go.mod" - Module name and Go version.
```

Calculates final letter grades from three categories:
- Assignments (50%)
- Exams (35%)
- Essays (15%)
```

Letter thresholds:
- A >= 90
- B >= 80
- C >= 70
- D >= 60
- F < 60
```

The library calculates a weighted average and converts it to a letter grade.
```