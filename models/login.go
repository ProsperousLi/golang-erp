package models

import (
	"errors"
	"time"

	"erpweb/logs"
	"erpweb/util"

	"github.com/mojocn/base64Captcha"
)

var (
	TokenMap = make(map[string]string)      //cardid token
	AccsMap  = make(map[string]interface{}) //token employee
	TimeMap  = make(map[string]int64)       //token TimeMap
)

//TODO Timer delete token
func New_Time_count() {
	//frist clear map
	TokenMap = make(map[string]string)
	AccsMap = make(map[string]interface{})
	TimeMap = make(map[string]int64)
}

// 登录接口
//{username: ‘test’, password: ‘test’, vckey: ‘ekksdfssl’, verifycode: ‘3qu5’}
func Login(cardid string, password string) (errorCode int64, uuid string) {
	var (
		qurey Account
	)
	qurey.Cardid = cardid
	//qurey.Password = password
	err := OSQL.Read(&qurey, "cardid")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return 1, uuid
	}

	if qurey.Cardid == cardid && qurey.Password == password { //login sucess
		if qurey.Status == 1 {
			preToken := TokenMap[cardid]
			uuid = util.GetToken()
			TokenMap[cardid] = uuid
			delete(TimeMap, preToken)
			TimeMap[uuid] = time.Now().Unix()
			delete(AccsMap, preToken)

			//qurey employee info
			emp, code := GetEmployeeByCardid(cardid)
			if code != util.SUCESSFUL {
				logs.FileLogs.Error("GetEmployeeByCardid failed")
				return 1, uuid
			}
			AccsMap[uuid] = emp
		} else {
			logs.FileLogs.Error("status is :%s", qurey.Status)
			return 1, uuid
		}
	} else {
		logs.FileLogs.Error("cardid or password is invild")
		return 1, uuid
	}

	return 0, uuid
}

// 单点登录
func SSOLogin(token string) (err error, code int64) {
	code = util.SUCESSFUL
	if _, ok := AccsMap[token]; ok {
		//pandding time
		if lastTime, ok := TimeMap[token]; ok {
			if lastTime+1*60*60 < time.Now().Unix() {
				//token过期了
				code = 50014
				logs.FileLogs.Error("token过期了")
				return errors.New("token过期了"), code
			}
		}
	} else {
		code = 50008
		logs.FileLogs.Error("this token not exist :%s", token)

		return errors.New("token not exist"), code
	}

	return nil, code
}

func Loginout(token string) (code int64) {
	code = util.SUCESSFUL
	if emp, ok := AccsMap[token]; ok {
		//delete all map
		TokenMap = make(map[string]string)     //cardid token
		AccsMap = make(map[string]interface{}) //token employee
		TimeMap = make(map[string]int64)       //token TimeMap

		tempEmp := emp.(Employee)
		delete(TokenMap, tempEmp.Cardid)
		delete(AccsMap, token)
		delete(TimeMap, token)
	} else {
		code = 50008
	}

	return code
}

//创建图像验证码
func demoCodeCaptchaCreate() string {
	//config struct for digits
	// //数字验证码配置
	// var configD = base64Captcha.ConfigDigit{
	// 	Height:     80,
	// 	Width:      240,
	// 	MaxSkew:    0.7,
	// 	DotCount:   80,
	// 	CaptchaLen: 5,
	// }
	// //config struct for audio
	// //声音验证码配置
	// var configA = base64Captcha.ConfigAudio{
	// 	CaptchaLen: 6,
	// 	Language:   "zh",
	// }
	//config struct for Character
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumberAlphabet,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	// //create a audio captcha.
	// idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	// //以base64编码
	// base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
	//create a characters captcha.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	// //create a digits captcha.
	// idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	// //以base64编码
	// base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

	//fmt.Println(idKeyA, base64stringA, "\n")
	logs.FileLogs.Info(idKeyC, base64stringC, "\n")
	//fmt.Println(idKeyD, base64stringD, "\n")

	return base64stringC
}

//验证图像验证码
func verfiyCaptcha(idkey, verifyValue string) {
	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
	if verifyResult {
		//success
	} else {
		//fail
	}
}
