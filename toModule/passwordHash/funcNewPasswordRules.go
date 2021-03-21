package passwordHash

func (e *Password) newPasswordRules(password []byte) (err error) {
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
