package SQLite3

func (e *SQLite3) Init() (err error) {
	err = e.userCreate()
	if err != nil {
		return
	}

	return
}
