package punishment

type Punishment struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	UUID           string `json:"uuid"`
	Reason         string `json:"reason"`
	Operator       string `json:"operator"`
	PunishmentType string `json:"punishment_type"`
	Start          string `json:"start"`
	End            string `json:"end"`
}
