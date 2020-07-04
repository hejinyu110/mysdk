package mysdk

import "encoding/json"

type RegisterResponse struct {
	Code    int     `json:"code"`
	UserId  string `json:"userId"`
	Token   string `json:"token"`
}

func (rc *RongCloud) Register(userId, name, portraitURI string) (*RegisterResponse, error)  {
	checkEmpty(userId,"userId")
	checkEmpty(name, "name")
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
