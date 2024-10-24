package fastly

import (
	"testing"
)

func TestClient_Product_Enablement_Fanout(t *testing.T) {
	t.Parallel()

	var err error

	// Enable product
	record(t, "product_enablement_fanout/enable", func(c *Client) {
		_, err = c.EnableProductFanout(&ProductFanoutEnablementInput{
			ServiceID: testServiceID,
		})
	})
	if err != nil {
		t.Fatal(err)
	}

	// Get product status
	record(t, "product_enablement_fanout/get", func(c *Client) {
		_, err = c.GetProductFanout(&ProductFanoutEnablementInput{
			ServiceID: testServiceID,
		})
	})
	if err != nil {
		t.Fatal(err)
	}

	// Disable product
	record(t, "product_enablement_fanout/disable", func(c *Client) {
		err = c.DisableProductFanout(&ProductFanoutEnablementInput{
			ServiceID: testServiceID,
		})
	})
	if err != nil {
		t.Fatal(err)
	}

	// Get product status again to check disabled
	record(t, "product_enablement_fanout/get_disabled", func(c *Client) {
		_, err = c.GetProductFanout(&ProductFanoutEnablementInput{
			ServiceID: testServiceID,
		})
	})

	// The API returns status code 400 if the product is not enabled
	if err == nil {
		t.Fatal("expected a 400 from the API but got a 2xx")
	}
}

func TestClient_GetProductFanout_validation(t *testing.T) {
	var err error

	_, err = testClient.GetProductFanout(&ProductFanoutEnablementInput{})
	if err != ErrMissingServiceID {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_EnableProductFanout_validation(t *testing.T) {
	var err error

	_, err = testClient.EnableProductFanout(&ProductFanoutEnablementInput{})
	if err != ErrMissingServiceID {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_DisableProductFanout_validation(t *testing.T) {
	var err error

	err = testClient.DisableProductFanout(&ProductFanoutEnablementInput{})
	if err != ErrMissingServiceID {
		t.Errorf("bad error: %s", err)
	}
}
