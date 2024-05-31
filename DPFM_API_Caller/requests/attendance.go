package requests

type Attendance struct {
	Event							int		`json:"Event"`
	Attender						int		`json:"Attender"`
	Attendance						int		`json:"Attendance"`
	Participation					int		`json:"Participation"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	IsCancelled						*bool	`json:"IsCancelled"`
}
