package main

import (
	"fmt"

	//"github.com/dcbaker/class_sched/rooms"
	"github.com/dcbaker/class-scheduler/schedule"
	"github.com/dcbaker/class-scheduler/subjects"
)

// XXX: need to take arguments or read config file or something
func main() {
	//aRooms := rooms.NewActualRooms([]rooms.Room{rooms.Art, rooms.Gym, rooms.Music})
	aTeachers := []schedule.Teacher{
		schedule.Teacher{"Carmen Smith", []subjects.Subject{subjects.SchoolSport}},
	}

	fmt.Println(aTeachers[0])
}
