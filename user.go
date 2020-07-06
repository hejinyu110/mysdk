package mysdk

import (
	"encoding/json"
	"errors"
	"math"
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
//用户更新
func (rc RongCloud)Refresh(userId, name, portraitURI string) (*Response, error)  {
	if userId ==""  {
		return nil, errors.New("userId is required")
	}
	if name ==""  {
		return nil, errors.New("name is required")
	}
	rc.url =  "/user/refresh.json"
	rc.data = map[string]interface{}{
		"userId": userId,
		"name": name,
		"portraitUri": portraitURI,
	}
	bytes, err := rc.post()
	if err != nil {
		return nil, err
	}
	resp := new(Response)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
//添加黑名单
func (rc RongCloud)AddBlackList(userId string, blackuserId ...string) (*Response, error) {
	if userId == "" {
		return nil, errors.New("userId is required")
	}
	if len(blackuserId) > 20 || len(blackuserId) <=0 {
		return nil, errors.New("blackuserId length in 1 - 20 ")
	}
	rc.url = "/user/blacklist/add.json"
	rc.data = map[string]interface{}{
		"userId": userId,
		"blackUserId":blackuserId,
	}
	bytes, err := rc.post()
	if err != nil {
		return nil, err
	}
	resp := new(Response)
	if err := json.Unmarshal(bytes,&resp); err != nil {
		return nil, err
	}
	return resp, err
}

type UsersResponse struct {
	Code  int `json:"code"`
	Users string `json:"users"`
}
func (rc RongCloud)GetBlackList(userId string) (*UsersResponse, error) {
	if userId == "" {
		return nil, errors.New("userId is required")
	}
	rc.url = "/user/blacklist/qurey.json"
	rc.data = map[string]interface{}{
		"userId" : userId,
	}
	bytes, err := rc.post()
	if err != nil {
		return nil, err
	}
	resp := new(BlackListResponse)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
//移除黑名单
func (rc RongCloud) RemoveBlackList(userId string, blackUserId ... string) (*Response, error)  {
	if userId == "" {
		return nil, errors.New("userId is required")
	}
	if len(blackUserId) ==0 || len(blackUserId) > 20 {
		return nil, errors.New("blackUserId len in 1 -20 ")
	}
	rc.url = "/user/blacklist/remove.json"
	rc.data = map[string]interface{}{
		"userId":userId,
		"blackUserId":blackUserId,
	}
	bytes, err := rc.post()
	if err != nil {
		return nil, err
	}
	resp := new(Response)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (rc RongCloud) AddWhiteList(userId string, whiteUserId ... string) (*UsersResponse, error)  {
	if userId == "" {
		return nil, errors.New("userId is required")
	}
	if len(whiteUserId) ==0 || len(whiteUserId) > 20 {
		return nil, errors.New("whiteUserId length in 1 - 20 ")
	}
	rc.url = "/user/whitelist/add.json"
	rc.data = map[string]interface{}{
		"userId" : userId,
		"whiteUserId" : whiteUserId,
	}
	bytes, err := rc.post()
	if err != nil {
		return nil, err
	}
	resp := new(UsersResponse)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (rc RongCloud) RemoveWhiteList(userId string, whiteUserId ... string) (*Response, error)  {
	if userId == "" {
		return nil, errors.New("userId is required")
	}
	if len(whiteUserId) ==0 || len(whiteUserId) > 20 {
		return nil, errors.New(" whiteUserId length in 1 -20 ")
	}
	rc.url = "/user/whitelist/remove.json"
	rc.data = map[string]interface{}{
		"userId" : userId,
		"whiteUserId": whiteUserId,
	}
	bytes, err := rc.post()
	if err != nil {
		return nil, err
	}
	resp := new(Response)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
//用户白名单
func (rc RongCloud) GetWhiteList(userId string) (*UsersResponse, error)  {
	if userId == "" {
		return nil, errors.New("userId is required")
	}
	rc.url = "/user/whitelist/query.json"
	rc.data = map[string]interface{}{
		"userId" :userId,
	}
	bytes, err := rc.post()
	if err != nil {
		return nil, err
	}
	resp := new(UsersResponse)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//禁用
func (rc RongCloud) AddBlock(minute int, userId ... string) (*Response, error) {
	if minute ==0 || minute > 43200 {
		return nil, errors.New("minute in 1 - 43200")
	}
	if len(userId) ==0 || len(userId) > 20 {
		return nil, errors.New("user id length in 1 - 20")
	}
	rc.url = "/user/block.json"
	rc.data = map[string]interface{}{
		"userId":userId,
		"minute": minute,
	}
	bytes, err := rc.post()
	if err != nil {
		return nil, err
	}
	resp := new(Response)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (rc RongCloud) RemoveBlock(userId ... string) (*Response, error)  {
	if len(userId) ==0 || len(userId) > 20 {
		return nil, errors.New("userId len in 1 - 20 ")
	}
	rc.url = "/user/unblock.json"
	rc.data = map[string]interface{}{
		"userId":userId,
	}
	bytes, err := rc.post()
	if err != nil {
		return nil, err
	}
	resp := new(Response)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
type BlockListResponse struct {
	Code int `json:"code"`
	Users []struct{
		UserId string `json:"userId"`
		BlockEndTime string `json:"blockEndTime"`
	} `json:"users"`
}
func (rc RongCloud) BlockList(page int, size int) (*BlockListResponse, error)  {
	rc.url = "/user/block/query.json"
	rc.data = map[string]interface{}{
		"page":page,
		"size":size,
	}
	bytes, err := rc.post()
	if err != nil {
		return nil, err
	}
	resp := new(BlockListResponse)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
//标签
func (rc RongCloud) TagSet(userId string, tags ... string) (*Response, error) {
	if userId == "" {
		return nil, errors.New("userId is required")
	}
	if len(tags) ==0 || len(tags) >20 {
		return nil, errors.New("tags length in 1 - 20")
	}
	rc.url = "/user/tag/set.json"
	rc.data = map[string]interface{}{
		"userId": userId,
		"tags": tags,
	}
	bytes, err := rc.postJson()
	if err != nil {
		return nil, err
	}
	resp := new(Response)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (rc RongCloud) TagBatchSet(userId, tags []string) (*Response, error)  {
	if len(userId) ==0 || len(userId) >20 {
		return nil, errors.New("userId length in 1 - 20 ")
	}
	if len(tags) ==0 || len(tags) >20 {
		return nil, errors.New("tags length in 1 - 20")
	}
	rc.url = "/user/tag/batch/set.json"
	rc.data = map[string]interface{}{
		"userIds": userId,
		"tags": tags,
	}
	bytes, err := rc.postJson()
	if err != nil {
		return nil, err
	}
	resp := new(Response)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type TagGetResponse struct {
	Code int `json:"code"`
	Result map[string][]string `json:"result"`
}

func (rc RongCloud) TagGet(userId string) (*TagGetResponse, error)  {
	if userId == "" {
		return nil, errors.New("userId is required ")
	}
	rc.url = "/user/tag/get.json"
	rc.data = map[string]interface{}{
		"userId": userId,
	}
	bytes, err := rc.postJson()
	if err != nil {
		return nil, err
	}
	resp := new(TagGetResponse)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}