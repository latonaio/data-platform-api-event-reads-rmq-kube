package requests

type Counter struct {
	Event					int		`json:"Event"`
	NumberOfLikes			int		`json:"NumberOfLikes"`
	NumberOfParticipations	int		`json:"NumberOfParticipations"`
	NumberOfAttendances		int		`json:"NumberOfAttendances"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	LastChangeDate			string	`json:"LastChangeDate"`
	LastChangeTime			string	`json:"LastChangeTime"`
}
