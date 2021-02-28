package password

func (e *Password) NewPasswordRules(password []byte) (err error) {
	err = e.ruleLength(password)
	if err != nil {
		return
	}

	err = e.ruleOneSpecialChars(password)
	if err != nil {
		return
	}

	err = e.ruleUpperLetter(password)
	if err != nil {
		return
	}

	err = e.ruleLowerCase(password)
	return
}