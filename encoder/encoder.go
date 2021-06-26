package encoder

import (
	"encoding/json"

	"github.com/MatejLach/1password-linux-to-bitwarden/decoder"
)

func (bwd *BitwardenPersonal) EncodeBitwardenImport() ([]byte, error) {
	return json.MarshalIndent(bwd, "", "    ")
}

func Encode1PasswordExportAsBitWardenImport(onePassExport decoder.OnePassword) (BitwardenPersonal, error) {
	btwFolders := make([]Folder, len(onePassExport.Accounts[0].Vaults))
	btwItems := make([]Entry, 0)

	for _, account := range onePassExport.Accounts {
		for _, vault := range account.Vaults {
			btwFolder := Folder{
				ID:   vault.Attrs.Uuid,
				Name: vault.Attrs.Name,
			}
			btwFolders = append(btwFolders, btwFolder)

			for _, entry := range vault.Items {
				var entryType int
				var entryUsername string
				var entryPassword string
				entryURIS := make([]URIEntry, len(entry.Item.Overview.Urls))

				for _, uri := range entry.Item.Overview.Urls {
					entryURIS = append(entryURIS, URIEntry{URI: uri.Url})
				}

				if len(entry.Item.Details.LoginFields) > 0 {
					for _, loginField := range entry.Item.Details.LoginFields {
						if loginField.Designation == "username" && loginField.Value != "" {
							entryType = 1
							entryUsername = loginField.Value
						} else if loginField.Designation == "password" && loginField.Value != "" {
							entryType = 1
							entryPassword = loginField.Value
						}
					}
				} else if len(entry.Item.Details.Sections) > 0 {
					for _, section := range entry.Item.Details.Sections {
						for _, loginField := range section.Fields {
							if loginField.Title == "email" && loginField.Value.Email != "" {
								entryType = 1
								entryUsername = loginField.Value.Email
							}
						}
					}
					if entry.Item.Details.Password != "" {
						entryPassword = entry.Item.Details.Password
					}
				}

				if entryUsername == "" || entryPassword == "" {
					continue
				}

				btwItems = append(btwItems, Entry{
					ID:       entry.Item.Uuid,
					FolderID: btwFolder.ID,
					Type:     entryType,
					Name:     entry.Item.Overview.Title,
					Login: LoginEntry{
						Uris:     entryURIS,
						Username: entryUsername,
						Password: entryPassword,
					},
				})
			}
		}
	}

	return BitwardenPersonal{
		Folders: btwFolders,
		Items:   btwItems,
	}, nil
}
