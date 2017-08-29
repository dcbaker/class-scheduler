package schedule

import (
	"reflect"

	"github.com/dcbaker/class_sched/rooms"
	"github.com/dcbaker/class_sched/subjects"
)

type Teacher struct {
	Name     string
	Subjects []subjects.Subject
}

func (t *Teacher) Equal(o *Teacher) bool {
	if t.Name != o.Name {
		return false
	}
	if len(t.Subjects) != len(o.Subjects) {
		return false
	}
	for i := 0; i <= len(t.Subjects); i++ {
		if t.Subjects[i] != o.Subjects[i] {
			return false
		}
	}
	return true
}

type Period struct {
	Teacher Teacher
	Subject subjects.Subject
	Room    rooms.Room
}

type Conflict int

//go:generate stringer -type=Conflict
const (
	TeacherBusy Conflict = iota
	RoomBusy
)

type scheduleConflict struct {
	message string
	reason  Conflict
}

func (e *scheduleConflict) Error() string {
	return e.message
}

func (e *scheduleConflict) Reason() Conflict {
	return e.reason
}

// The schedule for a single class
type ClassSchedule [5][4]Period

// The scehdule for a singel school, all classes
type SchoolSchedule []ClassSchedule

func (s *SchoolSchedule) isTeacherBusy(teacher Teacher, day int, period int) bool {
	for _, class := range *s {
		if class[day][period].Teacher.Equal(&teacher) {
			return true
		}
	}
	return false
}

func (s *SchoolSchedule) isRoomBusy(room rooms.Room, day int, period int) bool {
	for _, class := range *s {
		if reflect.DeepEqual(class[day][period].Room, room) {
			return true
		}
	}
	return false
}

func (s *SchoolSchedule) Schedule(session Period, class int, day int, period int) error {
	if s.isRoomBusy(session.Room, day, period) {
		return &scheduleConflict{"Room is busy", RoomBusy}
	}
	if s.isTeacherBusy(session.Teacher, day, period) {
		return &scheduleConflict{"Teacher is busy", TeacherBusy}
	}

	(*s)[class][day][period] = session

	return nil
}
