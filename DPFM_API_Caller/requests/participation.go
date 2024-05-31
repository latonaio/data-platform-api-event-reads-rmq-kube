package requests

type Participation struct {
	Event							int		`json:"Event"`
	Participator					int		`json:"Participator"`
	Participation					int		`json:"Participation"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	IsCancelled						*bool	`json:"IsCancelled"`
}
