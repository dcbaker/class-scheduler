package rooms

type Room int

//go:generate stringer -type=Room
const (
	Room1 Room = iota
	Room2
	Room3
	Room4
	Room5
	Room6
	Room7
	Room8
	Art
	Biology
	Chemistry
	Gym
	Music
	Swimming
)

type ActualRooms []Room

// ActualRooms is odd since it's not known at compile time what the length will
// be, but it is known at runtime and will be constant at runtime
func NewActualRooms(special []Room) *ActualRooms {
	constRooms := [8]Room{Room1, Room2, Room3, Room4, Room5, Room6, Room7, Room8}
	a := make(ActualRooms, 8, 16)
	for index, value := range constRooms {
		a[index] = value
	}
	for index, value := range special {
		a[index+len(constRooms)] = value
	}

	return &a
}
