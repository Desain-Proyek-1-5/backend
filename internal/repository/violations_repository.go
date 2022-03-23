package repository

import "distancing-detect-backend/internal/entity"

func (r *Repository) Create(violation *entity.ViolationData) error {
	stmt, err := r.db.Prepare(`INSERT INTO violations
	(class, totalviolations, timeofdetections, photolink)
	VALUES ($1, $2, $3, $4)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(violation.Class, violation.TotalViolations, violation.Time, violation.ImageLink)
	return err
}

func (r *Repository) List() ([]*entity.ViolationData, error) {
	var violations []*entity.ViolationData
	stmt, err := r.db.Prepare(`SELECT photolink, totalviolations, class, timeofdetection 
	FROM violations ORDER BY timeofdetections LIMIT 10`)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var violation entity.ViolationData
		err = rows.Scan(&violation.ImageLink, &violation.TotalViolations, &violation.Class, &violation.Time)
		if err != nil {
			return nil, err
		}
		violations = append(violations, &violation)
	}
	return violations, nil
}

func (r *Repository) GetByClass(class string) ([]*entity.ViolationData, error) {
	var violations []*entity.ViolationData
	stmt, err := r.db.Prepare(`SELECT photolink, totalviolations, class, timeofdetection 
	FROM violations WHERE class = $1 ORDER BY timeofdetections LIMIT 10`)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(class)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var violation entity.ViolationData
		err = rows.Scan(&violation.ImageLink, &violation.TotalViolations, &violation.Class, &violation.Time)
		if err != nil {
			return nil, err
		}
		violations = append(violations, &violation)
	}
	return violations, nil
}
