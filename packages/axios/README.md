# axios



//
//func NewAxios(req_url string, data map[string]string) *Axios {
//	axios := Axios{url: req_url, body: data, timeout: 30}
//	return &axios
//}
//

//func (a *Axios) Get() ([]byte, error) {
//	tr := &http.Transport{}
//	if a.httpsVerify {
//		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
//	}
//	if a.proxy != "" {
//		parse, err := url.Parse(a.proxy)
//		if err != nil {
//			return nil, err
//		}
//		tr.Proxy = http.ProxyURL(parse)
//	}
//
//	req_url := a.url
//	if len(a.body) > 0 {
//		params := url.Values{}
//		for k, v := range a.body {
//			params.Add(k, v)
//		}
//		req_url = fmt.Sprintf("%s?%s", a.url, params.Encode())
//	}
//
//	req, err := http.NewRequest("GET", req_url, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	if len(a.headers) > 0 {
//		for k, v := range a.headers {
//			req.Header.Set(k, v)
//		}
//	}
//
//	client := &http.Client{Transport: tr, Timeout: time.Duration(a.timeout) * time.Second}
//	response, err := client.Do(req)
//	if err != nil {
//		return nil, err
//	}
//	defer response.Body.Close()
//	if response.StatusCode != 200 {
//		return nil, errors.New(fmt.Sprintf("请求错误，错误状态码:%d", response.StatusCode))
//	}
//
//	b, err := io.ReadAll(response.Body)
//	if err != nil {
//		return nil, err
//	}
//	return b, nil
//}
//
//func (a *Axios) Post() ([]byte, error) {
//	tr := &http.Transport{}
//	if a.httpsVerify {
//		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
//	}
//	if a.proxy != "" {
//		parse, err := url.Parse(a.proxy)
//		if err != nil {
//			return nil, err
//		}
//		tr.Proxy = http.ProxyURL(parse)
//	}
//
//	marshal, err := json.Marshal(a.body)
//	if err != nil {
//		return nil, err
//	}
//	req, err := http.NewRequest("POST", a.url, bytes.NewReader(marshal))
//	if err != nil {
//		return nil, err
//	}
//	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
//
//	if len(a.headers) > 0 {
//		for k, v := range a.headers {
//			req.Header.Set(k, v)
//		}
//	}
//
//	client := http.Client{Transport: tr, Timeout: time.Duration(a.timeout) * time.Second}
//	response, err := client.Do(req)
//	if err != nil {
//		return nil, err
//	}
//	defer response.Body.Close()
//	if response.StatusCode != 200 {
//		return nil, errors.New(fmt.Sprintf("请求错误，错误状态码:%d", response.StatusCode))
//	}
//
//	b, err := io.ReadAll(response.Body)
//	if err != nil {
//		return nil, err
//	}
//	return b, nil
//}

//func (a *Axios) Dd(isPost bool) {
//	log.Println(strings.Repeat("=", 50))
//	tr := &http.Transport{}
//	if a.httpsVerify {
//		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
//	}
//	if a.proxy != "" {
//		parse, err := url.Parse(a.proxy)
//		if err != nil {
//			log.Println(err)
//			return
//		}
//		tr.Proxy = http.ProxyURL(parse)
//	}
//	log.Printf("代理地址：%s", a.proxy)
//
//	if len(a.headers) > 0 {
//		log.Printf("请求头:")
//		for k, v := range a.headers {
//			log.Printf("%s:%s", k, v)
//		}
//	}
//
//	if isPost {
//		log.Printf("请求地址：%s", a.url)
//		log.Printf("请求参数：")
//		for k, v := range a.body {
//			log.Printf("%s:%s", k, v)
//		}
//		log.Println(strings.Repeat("=", 50))
//	} else {
//		req_url := a.url
//		if len(a.body) > 0 {
//			params := url.Values{}
//			for k, v := range a.body {
//				params.Add(k, v)
//			}
//			req_url = fmt.Sprintf("%s?%s", a.url, params.Encode())
//		}
//		log.Printf("请求地址：%s", req_url)
//		log.Printf("请求参数:")
//		for k, v := range a.body {
//			log.Printf("%s:%s", k, v)
//		}
//		log.Println(strings.Repeat("=", 50))
//	}
//}

