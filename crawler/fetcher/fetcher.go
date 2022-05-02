package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(10* time.Millisecond)
func Fetch(url string) ([]byte, error) {
	// Every work will use rateLimiter
	<-rateLimiter
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	//request, err := http.NewRequest(http.MethodGet, url, nil)
	//if err!=nil {
	//	return nil,err
	//}
	//
	//request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.137 Safari/537.36 LBBROWSER")
	//resp, err :=http.DefaultClient.Do(request)
	//if  err!=nil{
	//	return nil,err
	//}

	//client := &http.Client{}
	//newUrl := strings.Replace(url, "http://", "https://", 1)
	//req, err := http.NewRequest("GET", newUrl, nil)
	//if err != nil {
	//	panic(err)
	//	return nil, err
	//}
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36")
	//cookie1 := "sid=3078aafa-e6ea-459b-a78d-a4e32254c167; ec=ffZoCV9y-1650080667470-311105da86e0b-1987680242; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1650080644,1650171373,1650466132; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1650802937; _exid=FIHjW1siUeB2f3X5ZOjK33%2FUHUcl5V%2BLguO47lQ4bp4oqPw4Vjvk%2FFgsyPIPoPP2C4X1mc7YmQMVo5QXVJfbzQ%3D%3D; _efmdata=phN5PQKUYHHyiRGFgnNFfZfuZYIVCn1IqMgzhCl1vr%2FEl4qzA%2BaFuFOd88mefoePv2W2%2FFCMd2Z2jrJSFI66UX4Y8tcrqhuMLKp2tvrsn1g%3D"
	//req.Header.Add("cookie", cookie1)
	//resp, err := client.Do(req)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer resp.Body.Close()
	//if resp.StatusCode != http.StatusOK {
	//	return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	//}

	defer resp.Body.Close()



	if resp.StatusCode != http.StatusOK {
		// fmt.Println("Error: status code", resp.StatusCode)
		return nil,
			fmt.Errorf("wrong status code %d", resp.StatusCode)
	}

	// return all, err
	return ioutil.ReadAll(resp.Body)

}
