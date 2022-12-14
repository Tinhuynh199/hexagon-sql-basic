package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "hexrestapi/internal/user/domain"
)

type UserClient struct {
	Client *http.Client
	URL string
}

func NewUserClient(client *http.Client, url string) *UserClient {
	return &UserClient{Client: client, URL: url}
}

func (c *UserClient) Load(ctx context.Context, id string) (*User, error) {
	requestURL := fmt.Sprintf("%s/%s", c.URL, id)

	request, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		panic(err)
	}

	response, err := c.Client.Do(request) 
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var res User
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	return &res, err
}

func (c *UserClient) Create(ctx context.Context, user *User) (int64, error) {
	requestURL := c.URL

	data, err := json.Marshal(user)
	request, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	response, err := c.Client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	if response.StatusCode == 201 {
		return 1, err
	} else {
		return 0, err
	}
}

func (c *UserClient) Update(ctx context.Context, user *User) (int64, error) {
	requestURL := fmt.Sprintf("%s/%s", c.URL, user.ID)

	data, err := json.Marshal(user)
	req, err := http.NewRequest("PUT", requestURL, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 200 {
		return 1, err
	} else {
		return 0, err
	}
}

func (c *UserClient) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	id := user["id"]
	requestURL := fmt.Sprintf("%s/%s", c.URL, id)

	data, err := json.Marshal(user)
	req, err := http.NewRequest("PATCH", requestURL, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 200 {
		return 1, err
	} else {
		return 0, err
	}
}

func (c *UserClient) Delete(ctx context.Context, id string) (int64, error) {
	requestURL := fmt.Sprintf("%s/%s", c.URL, id)

	req, err := http.NewRequest("DELETE", requestURL, nil)
	if err != nil {
		panic(err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 200 {
		return 1, err
	} else {
		return 0, err
	}
}

