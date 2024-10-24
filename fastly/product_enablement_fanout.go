package fastly

const ProductID = "fanout"

// ProductEnablementInput is used as input to the various product API functions.
type ProductFanoutEnablementInput struct {
	// ServiceID is the ID of the service (required).
	ServiceID string
}

// GetProductFanout retrieves the details of the product on the service.
func (c *Client) GetProductFanout(i *ProductFanoutEnablementInput) (bool, error) {
	if i.ServiceID == "" {
		return false, ErrMissingServiceID
	}

	path := ToSafeURL("enabled-products", ProductID, "services", i.ServiceID)

	resp, err := c.Get(path, nil)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var h *ProductEnablement
	if err := decodeBodyMap(resp.Body, &h); err != nil {
		return false, err
	}

	return true, nil
}

// EnableProductFanout enables the product on the service.
func (c *Client) EnableProductFanout(i *ProductFanoutEnablementInput) (bool, error) {
	if i.ServiceID == "" {
		return false, ErrMissingServiceID
	}

	path := ToSafeURL("enabled-products", ProductID, "services", i.ServiceID)

	resp, err := c.Put(path, nil)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var h *ProductEnablement
	if err := decodeBodyMap(resp.Body, &h); err != nil {
		return false, err
	}
	return true, nil
}

// DisableProductFanout disables the specified product on the service.
func (c *Client) DisableProductFanout(i *ProductFanoutEnablementInput) error {
	if i.ServiceID == "" {
		return ErrMissingServiceID
	}

	path := ToSafeURL("enabled-products", ProductID, "services", i.ServiceID)

	_, err := c.Delete(path, nil)
	return err
}
