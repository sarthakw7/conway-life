package utils

import (
	"conway-life/game"
	"encoding/json"
	"os"
)

const SaveFileName = "grid.json"

func SaveGrid(grid game.Grid) error {
	data, err := json.Marshal(grid)
	if err != nil {
		return err
	}
	return os.WriteFile(SaveFileName, data, 0644)
}

func LoadGrid() (game.Grid, error) {
	data, err := os.ReadFile(SaveFileName)
	if err != nil {
		return nil, err
	}

	var grid game.Grid
	err = json.Unmarshal(data, &grid)
	if err != nil {
		return nil, err
	}

	return grid, nil
}
