package books

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

func (client *Client) MakeAPIRequest(uri string, result any) error {
	resp, err := http.Get("http://" + client.addr + "/v1/" + uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return errors.New("not found")
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, result)
	if err != nil {
		return fmt.Errorf("%v in %q", err, data)
	}
	return nil
}

func (client *Client) GetBook(id string) (Book, error) {
	book := Book{}
	err := client.MakeAPIRequest("find/"+id, &book)
	if err != nil {
		return Book{}, err
	}
	return book, nil
}

func (client *Client) GetAllBooks() ([]Book, error) {
	bookList := []Book{}
	err := client.MakeAPIRequest("list", &bookList)
	if err != nil {
		return nil, err
	}
	return bookList, nil
}

func (client *Client) GetCopies(id string) (int, error) {
	copies := 0
	err := client.MakeAPIRequest("getcopies/"+id, &copies)
	if err != nil {
		return 0, err
	}
	return copies, nil
}

func (client *Client) AddCopies(id string, copies int) (int, error) {
	URI := fmt.Sprintf("/addcopies/%s/%d", id, copies)
	stock := 0
	err := client.MakeAPIRequest(URI, &stock)
	if err != nil {
		return 0, err
	}
	return stock, nil
}

func (client *Client) SubCopies(id string, copies int) (int, error) {
	URI := fmt.Sprintf("/subcopies/%s/%d", id, copies)
	stock := 0
	err := client.MakeAPIRequest(URI, &stock)
	if err != nil {
		return 0, err
	}
	return stock, nil
}
