package mysdk

import (
	"encoding/json"
	"errors"
)

type MessageHistoryResponse struct {
	Code int `json:"code"`
	Url  string `json:"url"`
	Date string `json:"date"`
}
func (rc RongCloud) MessageHistory(date string) (*MessageHistoryResponse, error) {
	if date == "" {
		return  nil, errors.New("date not null")
	}
	rc.url = "/message/history.json"
	rc.data = map[string]interface{}{
		"date": date,
	}
	bytes, err := rc.post()
	if err != nil {
		return nil, err
	}
	resp := new(MessageHistoryResponse)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (rc RongCloud) MessageHistoryDelete(date string)(*Response, error) {
	if date == "" {
		return nil, errors.New("date not null")
	}
	rc.url = "message/history/delete.json"
	rc.data = map[string]interface{}{
		"date": date,
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
func (rc RongCloud) MessagePrivatePush(msg RCMessage, fromUserId, targetUserId, objectName, pushContent, pushData string, isIncloudSender, count, verifyBlackList, isPersised, contentAvailable int) (*Response, error) {
	if fromUserId == "" {
		return nil, errors.New("fromUserId is required")
	}
	if targetUserId == "" {
		return nil, errors.New("targetUserId is required")
	}
	if objectName == "" {
		return nil, errors.New("objectName is required")
	}
	
	message, err := msg.JsonFormat()
	if err != nil {
		return nil, err
	}
	rc.url = "/message/private/push.json"
	rc.data = map[string]interface{}{
		"fromUserId": fromUserId,
		"targetUserId": targetUserId,
		"objectName": objectName,
		"pushContent": pushContent,
		"contet": string(message),
		"pushData": pushData,
		"isIncloudSender": isIncloudSender,
		"count": count,
		"isPersised": isPersised,
		"contentAvailable":contentAvailable,
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






