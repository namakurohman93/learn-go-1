package pkg

import (
	"encoding/json"
	"io/ioutil"
)

type Character struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Profession  string `json:"profession"`
	Affiliation string `json:"affiliation"`
	Status      string `json:"status"`
	Picture     string `json:"picture"`
}

func GetCharacterByte() ([]byte, error) {
	var bytes []byte
	var err error
	var characters []Character

	bytes, err = ioutil.ReadFile("./mr_robot.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &characters)
	if err != nil {
		return nil, err
	}

	bytes, err = json.Marshal(characters)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
