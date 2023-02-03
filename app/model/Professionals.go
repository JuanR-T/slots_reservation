package model

type Professional struct {
	ID           uint64 `json:"idProfessional"`
	ID_USER      uint64 `json:"idUser"`
	ID_TIMETABLE uint64 `json:"idTimetable"`
	NAME         string `json:"name"`
	ADDRESS      string `json:"address"`
	SERVICE      string `json:"service"`
}

func GetAllProfessionals() ([]Professional, error) {
	var professionals []Professional

	query := `select idProfessional, idUser,idTimetable, address, service, name from professionals;`

	rows, err := db.Query(query)
	if err != nil {
		return professionals, err
	}

	defer rows.Close()

	for rows.Next() {
		var idProfessional, idUser, idTimetable uint64
		var service, address, name string

		err := rows.Scan(&idProfessional, &idUser, &idTimetable, &address, &service, &name)
		if err != nil {
			return professionals, err
		}

		professional := Professional{
			ID:           idProfessional,
			ID_USER:      idUser,
			ID_TIMETABLE: idTimetable,
			NAME:         name,
			ADDRESS:      address,
			SERVICE:      service,
		}

		professionals = append(professionals, professional)
	}

	return professionals, nil
}

func GetProfessional(id uint64) (Professional, error) {
	var professional Professional

	query := `select idProfessional, idUser,idTimetable, address, service, name from professionals where idProfessional=$1;`
	row, err := db.Query(query, id)
	if err != nil {
		return professional, err
	}

	defer row.Close()

	if row.Next() {
		var idProfessional, idUser, idTimetable uint64
		var service, address, name string

		err := row.Scan(&idProfessional, &idUser, &idTimetable, &address, &service, &name)
		if err != nil {
			return professional, err
		}

		professional = Professional{
			ID:           idProfessional,
			ID_USER:      idUser,
			ID_TIMETABLE: idTimetable,
			NAME:         name,
			ADDRESS:      address,
			SERVICE:      service,
		}
	}
	return professional, nil
}

func CreateProfessional(professional Professional) error {

	query := `insert into professionals(idUser, idTimetable, name, address, service ) values($1, $2, $3, $4, $5);`

	_, err := db.Exec(query, professional.ID_USER, professional.ID_TIMETABLE, professional.NAME, professional.ADDRESS, professional.SERVICE)

	if err != nil {
		return err
	}

	return nil
}
func UpdateProfessional(professional Professional) error {

	query := `update professionals set idUser=$1, idTimetable=$2, name=$3, address=$4, service=$5 where idProfessional=$6;`

	_, err := db.Exec(query, professional.ID_USER, professional.ID_TIMETABLE, professional.NAME, professional.ADDRESS, professional.SERVICE, professional.ID)
	if err != nil {
		return err
	}
	return nil
}
func DeleteProfessional(id uint64) error {

	query := `delete from professionals where idProfessional=$1;`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil

}
