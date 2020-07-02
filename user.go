package mysdk

func (rc *RongCloud)register(userId, name, portraitURI string) ([]byte, error)  {
	checkEmpty(userId,"userId")
	checkEmpty(name, "name")
	rc.url =  "/user/getToken.json"
	rc.data = rmap{
		"userId": userId,
		"name": name,
		"portraitUri": portraitURI,
	}
	req, err := rc.post()
	if err != nil {
		return nil, err
	}


}
