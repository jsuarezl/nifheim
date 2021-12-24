package punishment

import (
	"database/sql"
	"fmt"
	"github.com/Beelzebu/nifheim/storage"
)

const (
	SelectPunishments    = "SELECT id, name, uuid, reason, operator, punishmentType, start, end FROM Punishments UNION SELECT id, name, uuid, reason, operator, punishmentType, start, end FROM PunishmentHistory"
	SelectPunishmentById = "SELECT id, name, uuid, reason, operator, punishmentType, start, end FROM Punishments WHERE id = ? UNION SELECT id, name, uuid, reason, operator, punishmentType, start, end FROM PunishmentHistory WHERE id = ?"
)

func GetPunishments() (punishments []Punishment) {
	punishments = []Punishment{}
	rows, err := storage.GetSQL().Query(SelectPunishments)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	for rows.Next() {
		var punishment Punishment
		if err := rows.Scan(&punishment.ID, &punishment.Name, &punishment.UUID, &punishment.Reason, &punishment.Operator, &punishment.PunishmentType, &punishment.Start, &punishment.End); err != nil {
			fmt.Println(err)
			continue
		}
		punishments = append(punishments, punishment)
	}
	return
}

func GetPunishment(id int) (punishment Punishment) {
	punishment = Punishment{}
	row := storage.GetSQL().QueryRow(SelectPunishmentById, id, id)
	err := row.Scan(&punishment.ID, &punishment.Name, &punishment.UUID, &punishment.Reason, &punishment.Operator, &punishment.PunishmentType, &punishment.Start, &punishment.End)
	if err != nil {
		return Punishment{}
	}
	return
}
