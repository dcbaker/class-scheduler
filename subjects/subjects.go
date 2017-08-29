package subjects

type Subject int

//go:generate stringer -type=Subject
const (
	Anatomy Subject = iota
	Art
	Biology
	StudyHall
	Chemistry
	ComputerScience
	Economics
	English
	Geography
	History
	Math
	Music
	Philosophy
	Physics
	Religion
	SchoolSport
	Science
	Swimming
	TheoreticalEd
	PracticalEd
)
