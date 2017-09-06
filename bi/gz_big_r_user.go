package main

import (
	"bytes"
	"compress/flate"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/astaxie/beego/httplib"
)

func main() {
	pullUrl := `http://tank-android-12.raysns.com/tank-server/public/index.php/getstat?type=get_user_audit_infos&is_compress=1&userids=[{"userids":["74000292","74000307","74000308","74000328","74000366","74000424","74000443","74000487","74000489","74000522","74000538","74000757","74000793","74000806","74000820","74000829","74000853","74000858","74000864","74000897","74000941","74000983","74001032","74001033","74001146","74001192","74001283","74001286","74001324","74001346","74001445","74001480","74001543","74001559","74001564","74001588","74001694","74001742","74001829","74001836","74001969","74002007","74002069","74002192","74002242","74002248","74002322","74002326","74002327","74002341"],"zid":"71"}]&sign=fa6b380db34e776ac280e04f4835770b`
	req := httplib.Get(pullUrl).SetTimeout(30*time.Second, 30*time.Second)

	retStr, err := req.String()
	if err != nil {
		fmt.Println("request err", err.Error())
	}

	retJson := GZinflate(retStr)
	if retJson == "[]" || retJson == "" {
		fmt.Println("%s", " data is empty or []")
	}

	ret := make(map[string]interface{})
	if err := json.Unmarshal([]byte(retJson), &ret); err != nil {
		fmt.Println("pull data Unmarshal error :" + pullUrl)
	}

	fmt.Println(ret)
	return
	items := ret["data"].([]interface{})
	datas := make([]map[string]interface{}, 0)
	for _, item := range items {
		if itemMap, ok := item.(map[string]interface{}); ok {
			datas = append(datas, itemMap)
		}
	}
	fmt.Println(datas)
}

func GZinflate(str string) string {
	b := bytes.NewReader([]byte(str))
	r := flate.NewReader(b)
	b2 := new(bytes.Buffer)
	_, err := io.Copy(b2, r)
	if err != nil {
		return err.Error()
	}
	defer r.Close()
	byts := b2.Bytes()
	return string(byts)
}
