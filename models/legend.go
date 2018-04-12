package models

import "api.meet.the/components/database"

type Legend struct {
	ID       int    `json:"legendId" form:"legendId"`
	Name     string `json:"name" form:"name"`
	PhotoURL string `json:"photo"`
}

func (p *Legend) GetLegends() ([]Legend, error) {

	var legends []Legend

	results, err := database.DB.Query("SELECT p.id, p.name, i.url FROM People p, Images i WHERE p.photoID = i.id AND p.active = 1 ORDER BY name ASC")
	if err != nil {
		return nil, err
	}

	for results.Next() {

		var legend Legend

		err = results.Scan(&legend.ID, &legend.Name, &legend.PhotoURL)

		if err != nil {
			return nil, err
		}

		legends = append(legends, legend)

	}

	return legends, nil

}
