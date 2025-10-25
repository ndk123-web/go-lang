// Package user contains teacher-related functionality
package user

// TeacherLogs stores log entries for teacher activities
var TeacherLogs []string

// Teacher represents a teacher with name and age
type Teacher struct {
	Name string // Teacher's name
	Age  int    // Teacher's age
}

// AddTeacher adds a teacher's name to the logs
func (t *Teacher) AddTeacher(name string) {
	TeacherLogs = append(TeacherLogs, name)
}
