package middleware

import (
	"testing"

	check "gopkg.in/check.v1"
	"io/ioutil"
	"os"
)

func Test(t *testing.T) { check.TestingT(t) }

type MiddlewareSuite struct{}

var _ = check.Suite(&MiddlewareSuite{})

func (s *MiddlewareSuite) TestNoConfig(c *check.C) {
	options := make(map[string]interface{})
	_, err := newCloudFrontStorageMiddleware(nil, options)
	c.Assert(err, check.ErrorMatches, "no baseurl provided")
}


func TestCloudFrontStorageMiddleware_GenerateKey(t *testing.T) {

	options := make(map[string]interface{})
	options["baseurl"] = "example.com"

	var privk = "-----BEGIN PRIVATE KEY-----\n"+
		"MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDKlvdvjS7EEkYm"+
		"RcMYBemMNvIbMRoJPBuwtRHph8pehi/BUXmTUsGUSDelFAl6tH2eMpHD6FpMJkiu"+
		"PAm875yJYo9nxgnH8SNzS1FSR05qoFkJGSvkQwLd++V4Z17kqg9LyzQf/XGSiGGm"+
		"VbvPS8rkJUdrFxgTjqhw8KIfNXGG4zLNcB6i0YgW2lZFgaFB16J3/O4KDm21bgdi"+
		"qI6F5xKel/YSQl+AfF2NlFt5xI3wuZ4YhD7wrp3WWgiS/JIxu0ESDmJTYExdb3K0"+
		"Dh6bmQl+wYJbkBpO2csNRDehtM2YEU83N8mmlSbIbSSinZYT3KwT9JQbd8QqJZCP"+
		"/mOPFacrAgMBAAECggEAMqGWR3fWd0RF6ezHfGqF2vgke+1Cn4o5NWmbh2zbg9Iv"+
		"fzYYl1w4axG9bnFaiSMwvefPjFG2t49d3MW+fUy5J5DNXFcfPKwkev0Y3uJZU8at"+
		"WdvDn3Gr9sSsrfHPwoBKAFxRs6kIyGFzXjnRDVbY5zn15mrIJqMhr9BEBF6798TA"+
		"qTG0FYTqkGK0D+FVfaWXvQ0u9Jw0KootS4kKHNwDmbZK2xYI2Ilt1ikeN9MMt8ZD"+
		"tXV4shTnQaYPty5Atr9Dzh052FTnlwsclVo33XHF2N2dfe7TaJYTaf5uXuh7Vkj3"+
		"99bKuvA8iLmEVe+i86L9K9LD5QEywO/sNcqQkNU8sQKBgQDkycMSv1dfSAgSZywU"+
		"hcJqFWJ3TpAP5aoWqw8Svill3Qs6/2zatU4XC4tzRX9KW437M4ORNee/XrCi9z4L"+
		"VOeOtR+gp/zn9DebBzaEdTfMlof+znPfx9fVIpBkpjezDgiFNkeeMnKTQrT/kULl"+
		"Zso29pCfgO/57L9Vi3pjiqKt0wKBgQDir4EDVesZTGcBstSRLIGrUQFmvxICPaMm"+
		"0PogvQhUv8UFMxx6nBl/ZB7AZWQH0TRlpVL7iqS2clO1dQOgrtgwgUV4m3Ml9Ivr"+
		"vI2fgCWzFudPst3oSu9Udc9Pq995Pl2nnwV612p3P6p3wgCaFjbouyyJZxNmHQj/"+
		"JPtv65nSSQKBgQDVcRnJorLLlHLbYF9yYfunZn3vWl7yRcvxy/KLBNewTZENoIAY"+
		"Zm8M9ttJVjvTziheg4ep8EVdduSJlOnQPoysyXNROYeril5aBlepKYY+Gu2THV5j"+
		"FpjYIZ/eFmf+Zwgx5xrXjq7vjZs4lnd3dvcOYec4t1yqqGE0WKR8uzjbuwKBgAYL"+
		"0FEadY7TLtwovOqyWTMMkhD/f6d3pWZfpIxC/nnkM4kT9+p9R2DSds+C5MwglFkx"+
		"s6jp5cLIAduRJ2udvj5s9EFnRAb7ItBC0zQx4s+ICNtjVe/gL8n86m6hkvBU7YKP"+
		"B0JjhH9xv0Y6cnGprgU/GM0BZs8ObzL+9YXirtOhAoGACk1TGT6I5oMHLDB4yOOe"+
		"a7dT+kAY6a2W2Sq2e0VN70EkVmXGC6ODpRzPcH7nojcMN5jk8QHHosWNK8DECwAj"+
		"uCGYn8G/0yAlhkddzE1+y1f5nVm+GCQTrdMMuqOwJnosifdoNDbWg4oGiBRt1uwI"+
		"aoxbWonlDAZaLC+8Bxe1Hss=\n"+
		"-----END PRIVATE KEY-----"


	file, err := ioutil.TempFile("", "pkey")
	if err != nil{
		t.Fatal("File cannot be created")
	}
	file.WriteString(privk)
	defer os.Remove(file.Name())
	t.Log(file.Name())
	options["privatekey"] = file.Name()
	options["keypairid"] = "test"
	newCloudFrontStorageMiddleware(nil, options)

}