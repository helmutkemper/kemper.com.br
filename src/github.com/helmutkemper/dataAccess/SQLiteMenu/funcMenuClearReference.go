package SQLiteMenu

func (e *SQLiteMenu) clearReference(ref *[]menuRef) {
	if ref == nil {
		return
	}

	for {
		pass := false
		for k := range *ref {
			if (*ref)[k].pass == true {
				*ref = append((*ref)[:k], (*ref)[k+1:]...)
				pass = true
				break
			}
		}

		if pass == false {
			break
		}
	}
}
