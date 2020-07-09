package mysdk

import (
    "bytes"
    "crypto/sha1"
    "encoding/hex"
    "encoding/json"
    "errors"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "strconv"
    "strings"
    "time"
)

const (
    SMSURL = "http://api.sms.ronghub.com"
    APIURL = "https://api-cn.ronghub.com"
    APIURL2 = "https://api2-cn.ronghub.com"
    UA = "RONGSDK-GO/4.0.1"
)

type RongCloud struct {
    appKey string
    appSecret string
    host string
    url string
    data map[string]interface{}
    request *http.Request
    header
}

type header struct {
    appKey string
    nonce string
    timestamp string
    signature string
}

type Response struct {
    Code int `json:"code"`
}

func NewRongCloud(appKey, appSecret string) *RongCloud  {
    return &RongCloud{
        appKey: appKey,
        appSecret: appSecret,
        host: APIURL,
    }
}

func (rc *RongCloud) setHeader() {
    nonce := strconv.Itoa(rand.Int())
    timestamp := strconv.Itoa(time.Now().Nanosecond())
    signstr := rc.appSecret + nonce + timestamp
    h := sha1.New()
    h.Write([]byte(signstr))
    signature := hex.EncodeToString(h.Sum(nil))

    rc.request.Header.Set("Timestamp", timestamp)
    rc.request.Header.Set("Nonce", nonce)
    rc.request.Header.Set("App-Key", rc.appKey)
    rc.request.Header.Set("Signature", signature)
    rc.request.Header.Set("User-Agent", UA)
}
func (rc *RongCloud)setParams() string{
    var params []string
    for k, v := range rc.data {
        switch v.(type) {
        case int:
            params = append(params, k + "=" + strconv.Itoa(v.(int)))
        case string:
            params = append(params, k + "="+ v.(string))
        case []int:
            item := v.([]int)
            for _, _v := range item {
                params = append(params, k + "=" + strconv.Itoa(_v))
            }
        case []string:
            item := v.([]string)
            for _, _v := range item {
                params = append(params, k + "=" + _v)
            }
        default:
            params = append(params, k + "="+ v.(string))
        }
    }
    return strings.Join(params, "&")
}
func (rc *RongCloud) post() ([]byte, error)  {
    params := rc.setParams()
    url := rc.host + rc.url
    req, err := http.NewRequest("POST", url, strings.NewReader(params))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    rc.request = req
    body, err := rc.doPost()
    return body, err
}
func (rc RongCloud) postJson() ([]byte, error)  {
    params, err := json.Marshal(rc.data)
    if err != nil {
        return nil, err
    }
    url := rc.host + rc.url
    req, err := http.NewRequest("POST", url, bytes.NewReader(params))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")
    rc.request = req
    body, err := rc.doPost()
    return body, err
}
func (rc *RongCloud) doPost() ([]byte, error) {
    rc.setHeader()
    client := &http.Client{}
    resp, err := client.Do(rc.request)
    log.Println(rc.request)
    if err != nil {
        return nil, err
    }
    if resp.StatusCode != 200 {//切换域名
        rc.switchHost()
        return nil, errors.New("请求错误，错误码" + strconv.Itoa(resp.StatusCode) )
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    return body, err

}
func (rc *RongCloud) switchHost()  {
    var host string
    if rc.host == APIURL {
        host = APIURL2
    }
    if rc.host == APIURL2 {
        host = APIURL
    }
    rc.host = host
}
