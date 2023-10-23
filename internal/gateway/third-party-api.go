package gateway

import (
	"encoding/json"
	"net/http"
	logger2 "testTask/pkg/logger"
)

type CountryProbability struct {
	Country     string  `json:"country_id"`
	Probability float32 `json:"probability"`
}

const (
	agify       = "https://api.agify.io/?name="
	nationalize = "https://api.nationalize.io/?name="
	genderize   = "https://api.genderize.io/?name="
)

type UserThirdPartyApi interface {
	GetAge(name string) (int, error)
	GetGender(name string) (string, error)
	GetCountry(name string) ([]CountryProbability, error)
}

type UserThirdPartyApiImpl struct {
	client *http.Client
	logger *logger2.Logger
}

func (u UserThirdPartyApiImpl) getJson(url string, target interface{}) error {
	r, err := u.client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type AgifyJSON struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u UserThirdPartyApiImpl) GetAge(name string) (int, error) {

	u.logger.Info.Println("Getting Age of ", name)
	res := AgifyJSON{}

	err := u.getJson(agify+name, &res)

	if err != nil {
		return 0, err
	}

	return res.Age, nil
}

type GenderizeJSON struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func (u UserThirdPartyApiImpl) GetGender(name string) (string, error) {
	u.logger.Info.Println("Getting Gender of ", name)

	res := GenderizeJSON{}

	err := u.getJson(genderize+name, &res)

	if err != nil {
		return "", err
	}

	return res.Gender, nil
}

type NationalizeJSON struct {
	Name    string               `json:"name"`
	Country []CountryProbability `json:"country"`
}

func (u UserThirdPartyApiImpl) GetCountry(name string) ([]CountryProbability, error) {
	u.logger.Info.Println("Getting Country of ", name)

	res := NationalizeJSON{}

	err := u.getJson(nationalize+name, &res)

	if err != nil {
		return nil, err
	}

	return res.Country, nil
}
