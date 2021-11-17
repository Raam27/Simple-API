package model

type Keycloak struct {
	Id              *string             `json:"id"`
	Username        string              `json:"username"`
	Email           string              `json:"email"`
	FirstName       string              `json:"firstName"`
	LastName        string              `json:"lastName"`
	Enabled         bool                `json:"enabled"`
	EmailVerified   bool                `json:"emailVerified"`
	Attributes      map[string][]string `json:"attributes"`
	Roles           []string            `json:"roles"`
	Groups          []string            `json:"groups"`
	RequiredActions []string            `json:"requiredActions"`
}

func NewKeycloak(username string, firstname string, email string) *Keycloak {
	kc := new(Keycloak)
	kc.Id = nil
	kc.Email = email
	kc.FirstName = firstname
	kc.LastName = "keycloak"
	kc.Username = username
	kc.EmailVerified = true
	kc.Enabled = true
	kc.Attributes = map[string][]string{"atribut 1": {"dari API", "Go"}}
	kc.Roles = []string{"role 1 dari API", "role 2 API"}
	kc.Groups = []string{"group 1 dari API", "group 2 API"}
	kc.RequiredActions = []string{"UPDATE_PROFILE"}
	return kc
}
