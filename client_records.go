package luadns

// ListRecords returns zone records.
//
// See: http://www.luadns.com/api.html#list-records
func (c *Client) ListRecords(zone *Zone) ([]*Record, error) {
	records := []*Record{}

	req := func() ([]byte, error) {
		return c.client.Get(c.endpoint("/zones/%d/records", zone.ID))
	}

	err := c.do(req, &records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

// CreateRecord creates a zone record using supplied attributes.
//
// See: http://www.luadns.com/api.html#create-a-record
func (c *Client) CreateRecord(zone *Zone, attrs *Record) (*Record, error) {
	var record Record

	req := func() ([]byte, error) {
		return c.client.Post(c.endpoint("/zones/%d/records", zone.ID), attrs)
	}

	err := c.do(req, &record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// GetRecord returns a zone record identified by `recordID`.
//
// See: http://www.luadns.com/api.html#get-a-record
func (c *Client) GetRecord(zone *Zone, recordID int64) (*Record, error) {
	var record Record

	req := func() ([]byte, error) {
		return c.client.Get(c.endpoint("/zones/%d/records/%d", zone.ID, recordID))
	}

	err := c.do(req, &record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// UpdateRecord updates a zone record identfied by `recordID` using supplied attributes.
//
// See: http://www.luadns.com/api.html#update-a-record
func (c *Client) UpdateRecord(zone *Zone, recordID int64, attrs *Record) (*Record, error) {
	var record Record

	req := func() ([]byte, error) {
		return c.client.Put(c.endpoint("/zones/%d/records/%d", zone.ID, recordID), attrs)
	}

	err := c.do(req, &record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// DeleteRecord deletes a zone record identfied by `recordID`.
//
// See: http://www.luadns.com/api.html#delete-a-record
func (c *Client) DeleteRecord(zone *Zone, recordID int64) (*Record, error) {
	var record Record

	req := func() ([]byte, error) {
		return c.client.Delete(c.endpoint("/zones/%d/records/%d", zone.ID, recordID))
	}

	err := c.do(req, &record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}
