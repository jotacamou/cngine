package actions

import (
	"encoding/json"
	"log"

	"github.com/gobuffalo/buffalo"
)

type SearchCriteria struct {
	City string `json:"city"`
}

// Search default implementation.
func Search(c buffalo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	var search SearchCriteria
	err := decoder.Decode(&search)
	if err != nil {
		log.Println(err)
	}
	//return c.Render(200, r.JSON(t))
	return c.Render(200, r.JSON(map[string]string{"message": search.City}))
}
