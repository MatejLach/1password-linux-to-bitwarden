package decoder

type Entry struct {
	Uuid         string `json:"uuid"`
	FavIndex     int    `json:"favIndex"`
	CreatedAt    int    `json:"createdAt"`
	UpdatedAt    int    `json:"updatedAt"`
	Trashed      bool   `json:"trashed"`
	CategoryUuid string `json:"categoryUuid"`
	Details      struct {
		LoginFields []struct {
			Value       string `json:"value"`
			Id          string `json:"id"`
			Name        string `json:"name,omitempty"`
			FieldType   string `json:"fieldType"`
			Designation string `json:"designation,omitempty"`
		} `json:"loginFields"`
		NotesPlain string `json:"notesPlain,omitempty"`
		Sections   []struct {
			Title  string `json:"title"`
			Name   string `json:"name,omitempty"`
			Fields []struct {
				Title string `json:"title"`
				Id    string `json:"id"`
				Value struct {
					Email     string      `json:"email,omitempty"`
					String    string      `json:"string,omitempty"`
					Concealed string      `json:"concealed,omitempty"`
					Url       string      `json:"url,omitempty"`
					Menu      string      `json:"menu,omitempty"`
					Date      interface{} `json:"date"`
					Address   struct {
						Street  string `json:"street"`
						City    string `json:"city"`
						Country string `json:"country"`
						Zip     string `json:"zip"`
						State   string `json:"state"`
					} `json:"address,omitempty"`
					Phone string `json:"phone,omitempty"`
				} `json:"value"`
				IndexAtSource int  `json:"indexAtSource"`
				Guarded       bool `json:"guarded"`
				Multiline     bool `json:"multiline"`
				DontGenerate  bool `json:"dontGenerate"`
				InputTraits   struct {
					Keyboard       string `json:"keyboard"`
					Correction     string `json:"correction"`
					Capitalization string `json:"capitalization"`
				} `json:"inputTraits"`
			} `json:"fields"`
		} `json:"sections"`
		PasswordHistory []struct {
			Value string `json:"value"`
			Time  int    `json:"time"`
		} `json:"passwordHistory"`
		HtmlForm struct {
			HtmlMethod string `json:"htmlMethod"`
			HtmlAction string `json:"htmlAction,omitempty"`
			HtmlId     string `json:"htmlId,omitempty"`
			HtmlName   string `json:"htmlName,omitempty"`
		} `json:"htmlForm,omitempty"`
		Password string `json:"password,omitempty"`
	} `json:"details"`
	Overview struct {
		Subtitle string `json:"subtitle"`
		Urls     []struct {
			Label string `json:"label"`
			Url   string `json:"url"`
		} `json:"urls,omitempty"`
		Title      string   `json:"title"`
		Url        string   `json:"url"`
		Ps         int      `json:"ps"`
		Pbe        float64  `json:"pbe"`
		Pgrng      bool     `json:"pgrng"`
		Tags       []string `json:"tags,omitempty"`
		B5UserUuid string   `json:"b5UserUuid,omitempty"`
	} `json:"overview"`
}

type OnePassword struct {
	Accounts []struct {
		Attrs struct {
			AccountName string `json:"accountName"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			Email       string `json:"email"`
			Uuid        string `json:"uuid"`
			Domain      string `json:"domain"`
		} `json:"attrs"`
		Vaults []struct {
			Attrs struct {
				Uuid   string `json:"uuid"`
				Desc   string `json:"desc"`
				Avatar string `json:"avatar"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"attrs"`
			Items []struct {
				Item Entry `json:"item"`
			} `json:"items"`
		} `json:"vaults"`
	} `json:"accounts"`
}
