package state

type ClientState struct {
	token        string
	isAuthorized bool
}

func NewClientState() *ClientState {
	return &ClientState{}
}

func (c *ClientState) IsAuthorized() bool {
	return c.isAuthorized
}

func (c *ClientState) SetIsAuthorized(isAuthorized bool) {
	c.isAuthorized = isAuthorized
}

func (c *ClientState) GetToken() string {
	return c.token
}

func (c *ClientState) SetToken(token string) {
	c.token = token
}
