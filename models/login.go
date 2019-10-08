package models

import (
	"errors"
	"time"

	"erpweb/logs"
	"erpweb/util"

	"github.com/mojocn/base64Captcha"
)

var (
	TokenMap = make(map[string]string) //cardid token
	AccsMap  = make(map[string]string) //token employee's cardid
	TimeMap  = make(map[string]int64)  //token TimeMap
	LimitMap = make(map[string]int64)  //限制5次登录机会,成功后清除
)

//TODO Timer delete token
func New_Time_count() {
	//frist clear map
	TokenMap = make(map[string]string)
	AccsMap = make(map[string]string)
	TimeMap = make(map[string]int64)
	LimitMap = make(map[string]int64)
}

// 登录接口
//{username: ‘test’, password: ‘test’, vckey: ‘ekksdfssl’, verifycode: ‘3qu5’}

// url: api/login/login
// 方法: POST
// Body：{username: ‘test’, password: ‘test’, vckey: ‘ekksdfssl’, verifycode: ‘3qu5’}
// 返回：{code: 20000, data: {detailcode: 0, token: ‘’}}
// 		 或{code: 20000, data: {detailcode: 1, msg: ‘账号密码不正确’, vckey: ‘’, verifycode: ‘dsfwekfsldfklsdfkkkk’}}，其中verifycode是验证码图片的base64编码
// 		或{code: 20000, data: {detailcode: 2, msg: ‘验证码错误’, vckey: ‘新key’, verifycode: ‘新验证码’}}
// 		或{code: 20000, data: {detailcode: 3, msg: ‘您的账号已锁定，请联系管理员解除锁定’}}
// 详细说明：第一次登录不需要验证码，body的vckey和verifycode为空或不带。连续失败5次账号锁定。如果用户名不存在，则一直返回第二种情况。
func Login(cardid string, password string, vckey, verifycode string) (errorCode int64, uuid, retvckey, retverifycode string) {
	var (
		qurey      Account
		loginTimes int64
	)
	qurey.Cardid = cardid
	//qurey.Password = password
	err := OSQL.Read(&qurey, "cardid")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return 1, uuid, retvckey, retverifycode
	}

	defer func() {
		LimitMap[cardid] = loginTimes
	}()

	//校验是否需要验证码（是否是第一次登录）
	if value, ok := LimitMap[cardid]; ok {
		loginTimes = value + 1
		if loginTimes > 1 {
			//检查验证码是否正确
			if vckey == "" || verifycode == "" {
				retvckey, retverifycode = CodeCaptchaCreate()
				return 2, uuid, retvckey, retverifycode
			}

		} else if value >= 5 {
			//update status 3
			EditAccountStatusById(cardid, 3)
			return 3, uuid, retvckey, retverifycode
		}
	} else {
		loginTimes += 1
	}

	if qurey.Cardid == cardid && qurey.Password == password { //login sucess
		if qurey.Status == 1 {
			preToken := TokenMap[cardid]
			uuid = util.GetToken()
			TokenMap[cardid] = uuid
			delete(TimeMap, preToken)
			TimeMap[uuid] = time.Now().Unix()
			delete(AccsMap, preToken)
			delete(LimitMap, cardid)

			//qurey employee info
			emp, code := GetEmployeeByCardid(cardid)
			if code != util.SUCESSFUL {
				logs.FileLogs.Error("GetEmployeeByCardid failed")
				return 1, uuid, retvckey, retverifycode
			}
			AccsMap[uuid] = emp.Cardid
		} else {
			logs.FileLogs.Error("status is :%s", qurey.Status)
			return 3, uuid, retvckey, retverifycode
		}
	} else {
		logs.FileLogs.Error("cardid or password is invild")
		if loginTimes == 1 {
			//new vckey and verifycode
			retvckey, retverifycode = CodeCaptchaCreate()
		} else {
			retvckey = vckey
			retverifycode = verifycode
		}
		return 1, uuid, retvckey, retverifycode
	}

	return 0, uuid, retvckey, retverifycode
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
		logs.FileLogs.Error("token不存在 :%s", token)

		return errors.New("token不存在"), code
	}

	return nil, code
}

func Loginout(token string) (code int64) {
	code = util.SUCESSFUL
	if cardid, ok := AccsMap[token]; ok {
		//delete all map
		// TokenMap = make(map[string]string) //cardid token
		// AccsMap = make(map[string]int64)   //token employee
		// TimeMap = make(map[string]int64)   //token TimeMap
		delete(TokenMap, cardid)
		delete(AccsMap, token)
		delete(TimeMap, token)
	} else {
		code = 50008
	}

	return code
}

//创建图像验证码
func CodeCaptchaCreate() (string, string) {
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

	return idKeyC, base64stringC
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

// 9.获取登录账户的用户信息(根据token)
// url: api/user/info
// 方法: GET
// 参数：无
// 返回：
// {
// 	code: 20000,
// 	data: {
// 		userInfo: {
// 			id: 1,
// 			name: ‘测试’,
// 			sex: 0,
// 			cardid: ‘bxkc - 001’,
// 			compID: 0,
// 			deptID: 1
// 		},
// 		permission: {
// 			read: [1, 2, 3],
// 			write: [4, 5, 6]
// 		},
// 		menuList: [{
// 				menuID: 1,
// 				title: '销售管理',
// 				icon: 'el-icon-location',
// 				children: [{
// 						menuID: 2,
// 						title: '维修合同',
// 						path: '/salerepair',
// 						component: 'Sale/RepairContract',
// 						children: []
// 					]
// 				}]
// 		}
// 	}
// }
// 说明：这个接口是登录后通过登录接口返回的token信息来获取账号信息：
// userInfo的数据来源为人员信息表employee，
// permission的数据来源为权限信息表permission，
// menuList的数据来源为菜单表menu，
// 但是返回时要根据permission处理，
// permission仅包含了叶子节点（即component不为空的节点的menuID），
// 返回时仅返回过滤后的节点。

type UserInfoWeb struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Sex    int8   `json:"sex"`
	Cardid string `json:"cardid"`
	CompID int8   `json:"compID"`
	DeptID int    `json:"deptID"`
}

type PermissionWeb struct {
	Read  []int64 `json:"read"`
	Write []int64 `json:"write"`
}

type UserInfoStruct struct {
	UserInfo   UserInfoWeb   `json:"userInfo"`
	Permission PermissionWeb `json:"permission"`
	MenuList   WebMenu       `json:"menuList"`
}

func GetUserInfo(token string) (errorCode int64, userInfo UserInfoStruct) {
	if cardId, ok := TokenMap[token]; ok {
		if cardid, ok2 := AccsMap[cardId]; ok2 {
			employee, code := GetEmployeeByCardid(cardid)
			if code != util.SUCESSFUL {
				return code, userInfo
			}

			userInfo.UserInfo.Cardid = employee.Cardid
			userInfo.UserInfo.CompID = employee.CompID
			userInfo.UserInfo.DeptID = employee.DeptID
			userInfo.UserInfo.Id = employee.Id
			userInfo.UserInfo.Name = employee.Name
			userInfo.UserInfo.Sex = employee.Sex

			permis, err := QueryPermission(cardId)
			if err != nil {
				return util.SUCESSFUL, userInfo
			}

			userInfo.Permission.Read = permis.Read
			userInfo.Permission.Write = permis.Write
		}
	}
	return errorCode, userInfo
}
