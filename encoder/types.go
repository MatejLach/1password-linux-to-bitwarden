package encoder

type Entry struct {
	ID             string      `json:"id"`
	OrganizationID interface{} `json:"organizationId"`
	FolderID       string      `json:"folderId"`
	Type           int         `json:"type"`
	Name           string      `json:"name"`
	Notes          string      `json:"notes"`
	Favorite       bool        `json:"favorite"`
	Fields         []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
		Type  int    `json:"type"`
	} `json:"fields"`
	SecureNote struct {
		Type int `json:"type"`
	} `json:"secureNote,omitempty"`
	CollectionIds []interface{} `json:"collectionIds"`
	Card          struct {
		CardholderName string `json:"cardholderName"`
		Brand          string `json:"brand"`
		Number         string `json:"number"`
		ExpMonth       string `json:"expMonth"`
		ExpYear        string `json:"expYear"`
		Code           string `json:"code"`
	} `json:"card,omitempty"`
	Identity struct {
		Title          string      `json:"title"`
		FirstName      string      `json:"firstName"`
		MiddleName     string      `json:"middleName"`
		LastName       string      `json:"lastName"`
		Address1       string      `json:"address1"`
		Address2       interface{} `json:"address2"`
		Address3       interface{} `json:"address3"`
		City           string      `json:"city"`
		State          string      `json:"state"`
		PostalCode     string      `json:"postalCode"`
		Country        string      `json:"country"`
		Company        string      `json:"company"`
		Email          string      `json:"email"`
		Phone          string      `json:"phone"`
		Ssn            string      `json:"ssn"`
		Username       string      `json:"username"`
		PassportNumber string      `json:"passportNumber"`
		LicenseNumber  string      `json:"licenseNumber"`
	} `json:"identity,omitempty"`
	Login LoginEntry `json:"login,omitempty"`
}

type LoginEntry struct {
	Uris     []URIEntry `json:"uris"`
	Username string     `json:"username"`
	Password string     `json:"password"`
	Totp     string     `json:"totp"`
}

type URIEntry struct {
	Match interface{} `json:"match"`
	URI   string      `json:"uri"`
}

type Folder struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BitwardenPersonal struct {
	Folders []Folder `json:"folders"`
	Items   []Entry  `json:"items"`
}
