package luadns

func (c *Client) Me() (*User, error) {
	var user User

	req := func() ([]byte, error) {
		return c.client.Get(c.endpoint("/users/me"))
	}

	err := c.do(req, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
