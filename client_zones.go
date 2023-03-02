package luadns

// ListZones returns user zones.
//
// See: http://www.luadns.com/api.html#list-zones
func (c *Client) ListZones() ([]*Zone, error) {
	zones := []*Zone{}

	req := func() ([]byte, error) {
		return c.client.Get(c.endpoint("/zones"))
	}

	err := c.do(req, &zones)
	if err != nil {
		return nil, err
	}

	return zones, nil
}

// CreateZone creates a new zone using supplied attributes.
//
// See: http://www.luadns.com/api.html#create-a-zone
func (c *Client) CreateZone(attrs *Zone) (*Zone, error) {
	var zone Zone

	req := func() ([]byte, error) {
		return c.client.Post(c.endpoint("/zones"), attrs)
	}

	err := c.do(req, &zone)
	if err != nil {
		return nil, err
	}

	return &zone, nil
}

// GetZone get a specific zone identfied by `zoneID`.
//
// See: http://www.luadns.com/api.html#get-a-zone
func (c *Client) GetZone(zoneID int64) (*Zone, error) {
	var zone Zone

	req := func() ([]byte, error) {
		return c.client.Get(c.endpoint("/zones/%d", zoneID))
	}

	err := c.do(req, &zone)
	if err != nil {
		return nil, err
	}

	return &zone, nil
}

// UpdateZone zone identified by `zoneID` using supplied attributes.
//
// See: http://www.luadns.com/api.html#update-a-zone
func (c *Client) UpdateZone(zoneID int64, attrs *Zone) (*Zone, error) {
	var zone Zone

	req := func() ([]byte, error) {
		return c.client.Put(c.endpoint("/zones/%d", zoneID), attrs)
	}

	err := c.do(req, &zone)
	if err != nil {
		return nil, err
	}

	return &zone, nil
}

// DeleteZone delete specific zone using zone ID.
//
// See: http://www.luadns.com/api.html#delete-a-zone
func (c *Client) DeleteZone(zoneID int64) (*Zone, error) {
	var zone Zone

	req := func() ([]byte, error) {
		return c.client.Delete(c.endpoint("/zones/%d", zoneID))
	}

	err := c.do(req, &zone)
	if err != nil {
		return nil, err
	}

	return &zone, nil
}
