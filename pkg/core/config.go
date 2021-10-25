package core

type Config struct {
	Auth struct {
		ClientID        string `yaml:"ClientID"`
		UserAccessToken string `yaml:"UserAccessToken"`
		AppAccessToken  string `yaml:"AppAccessToken"`
	}
}
