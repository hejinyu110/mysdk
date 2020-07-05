package mysdk

import (
	"encoding/json"
	"errors"
)

type RegisterResponse struct {
	Code    int     `json:"code"`
	UserId  string `json:"userId"`
	Token   string `json:"token"`
}

func (rc *RongCloud) Register(userId, name, portraitURI string) (*RegisterResponse, error)  {
	if userId ==""  {
		return nil, errors.New("userId is required")
	}

	if name ==""  {
		return nil, errors.New("name is required")
	}

	rc.url =  "/user/getToken.json"
	rc.data = map[string]interface{}{
		"userId": userId,
		"name": name,
		"portraitUri": portraitURI,
	}
	bytes, err := rc.post()
	if err != nil {
		return nil, err
	}
	resp := new(RegisterResponse)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil


}
