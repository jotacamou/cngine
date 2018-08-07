package actions

import (
	"github.com/gobuffalo/buffalo"
)

type AssetLists struct {
	Cities     []string `bson:"cities"`
	Activities []string `bson:"activities"`
}

// AssetsAssets default implementation.
func Assets(c buffalo.Context) error {
	var result AssetLists

	collection, err := GetCollection("indexdata")
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	err = collection.Find(nil).One(&result)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(200, r.JSON(result))
}
