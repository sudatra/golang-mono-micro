package orders

import (
	"errors"
	"fmt"
	"net/http"
)

type HTTPClient struct {
	address string
}

func NewHTTPClient(address string) HTTPClient {
	return HTTPClient{
		address: address,
	}
}

func (h HTTPClient) MarkOrderAsPaid(orderID string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/orders/%s/paid", h.address, orderID), nil)
	if err != nil {
		return errors.New("unable to create request")
	}

	res, err := http.DefaultClient.Do(req);
	if err != nil {
		return errors.New("request to orders failed");
	}

	return res.Body.Close();
}