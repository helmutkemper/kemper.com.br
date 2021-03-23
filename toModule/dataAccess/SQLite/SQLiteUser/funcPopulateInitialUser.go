package SQLiteUser

import (
	"log"
)

// populateInitialMenu (Português): popula o menu com os primeiros dados após a instalação.
func (e *SQLiteUser) populateInitialUser() (err error) {

	err = e.Set(
		"24867707-3a21-4368-8058-3d7b1ddb8c06",
		"5996b891-9d3c-4038-af37-cb07f5f0f72d",
		1,
		"Helmut Kemper",
		"Kemper",
		"helmut.kemper@gmail.com",
		"admin",
	)
	if err != nil {
		log.Printf("SQLiteUser.populateInitialUser().error: %v", err.Error())
	}
	return
}
