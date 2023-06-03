package factory

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"type-registry/players"
	"type-registry/players/baseball"
	"type-registry/players/football"
)

func UnmarshalBallPlayer(ballType string, baseDir string) ([]players.Player, error) {
	files, err := getFilePathsFromBaseDir(fmt.Sprintf("%s/%s", baseDir, ballType))
	if err != nil {
		return nil, err
	}
	switch ballType {
	case "baseball":
		return unmarshalBaseballPlayers(files)
	case "football":
		return unmarshalFootballPlayers(files)
	default:
		return nil, fmt.Errorf("%s is not supported", ballType)
	}

}

func unmarshalBaseballPlayers(files []string) ([]players.Player, error) {
	var players []players.Player
	for _, jsonFile := range files {
		var newPlayer baseball.Player
		jsonBytes, err := retrieveJsonFromFile(jsonFile)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(jsonBytes, &newPlayer)
		players = append(players, &newPlayer)
	}

	return players, nil
}

func unmarshalFootballPlayers(files []string) ([]players.Player, error) {
	var players []players.Player
	for _, jsonFile := range files {
		var newPlayer football.Player
		jsonBytes, err := retrieveJsonFromFile(jsonFile)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(jsonBytes, &newPlayer)
		players = append(players, &newPlayer)
	}

	return players, nil
}

func getFilePathsFromBaseDir(baseDir string) ([]string, error) {
	var results []string
	files, err := os.ReadDir(baseDir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		results = append(results, fmt.Sprintf("%s/%s", baseDir, file.Name()))
	}
	return results, nil
}

func retrieveJsonFromFile(filePath string) ([]byte, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	return io.ReadAll(jsonFile)
}
