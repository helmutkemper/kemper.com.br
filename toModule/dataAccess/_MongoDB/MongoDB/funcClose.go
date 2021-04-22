package MongoDB

func (e *MongoDB) Close() (err error) {
	err = e.Client.Disconnect(e.Ctx)
	return
}
