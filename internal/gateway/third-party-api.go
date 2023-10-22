package gateway

type CountryProbability struct {
	Country     string
	Probability float32
}

const (
	agify       = "https://api.agify.io/"
	nationalize = "https://api.nationalize.io/"
	genderize   = "https://api.genderize.io/"
)

type UserThirdPartyApi interface {
	GetAge(name string) (int, error)
	GetGender(name string) (string, error)
	GetCountry(name string) ([]CountryProbability, error)
}

type UserThirdPartyApiImpl struct {
}

func (u UserThirdPartyApiImpl) GetAge(name string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserThirdPartyApiImpl) GetGender(name string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserThirdPartyApiImpl) GetCountry(name string) ([]CountryProbability, error) {
	//TODO implement me
	panic("implement me")
}
