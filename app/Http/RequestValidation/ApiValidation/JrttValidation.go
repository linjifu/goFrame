package ApiValidation

type JrttValidation struct {
	AdvertiserId  string `form:"advertiser_id"`  // 广告主id
	Aid           string `form:"aid"`            // 广告计划id
	AidName       string `form:"aid_name"`       // 广告计划名称
	Cid           string `form:"cid"`            // 广告创意id，长整型
	CidName       string `form:"cid_name"`       // 广告创意名称
	CampaignId    string `form:"campaign_id"`    // 广告组id
	Csite         string `form:"csite"`          // 广告投放位置
	Ip            string `form:"ip"`             // 媒体投放系统获取的用户终端的公共IP地址
	Model         string `form:"model"`          // 手机型号
	CampaignName  string `form:"campaign_name"`  // 广告组名称
	Ctype         int16  `form:"ctype"`          // 创意样式
	ConvertId     int64  `form:"convert_id"`     // 转化id
	RequestId     string `form:"request_id"`     // 请求下发的id
	Sl            string `form:"sl"`             // 这次请求的语言
	Imei          string `form:"imei"`           // (原始值做md5)安卓的设备 ID 的 md5 摘要，32位
	Idfa          string `form:"idfa"`           // IOS 6+的设备id字段，32位
	Androidid     string `form:"androidid"`      // (原始值做md5)安卓id原值的md5，32位
	Oaid          string `form:"oaid"`           // (原始值)Android Q及更高版本的设备号，32位
	OaidMd5       string `form:"oaid_md5"`       // (oaid原始值做md5)Android Q及更高版本的设备号的md5摘要，32位
	Os            uint8  `form:"os"`             // 操作系统平台
	Mac           string `form:"mac"`            // 移动设备mac地址,转换成大写字母,去掉“:”，并且取md5摘要后的结果
	Mac1          string `form:"mac1"`           // 移动设备 mac 地址,转换成大写字母,并且取md5摘要后的结果，32位
	Ua            string `form:"ua"`             // 用户代理(User Agent)，一个特殊字符串头，使得服务器能够识别客户使用的操作系统及版本、CPU类型、浏览器及版本、浏览器渲染引擎、浏览器语言、浏览器插件等。
	Geo           string `form:"geo"`            // 位置信息，包含三部分:latitude（纬度），longitude（经度）以及precise（确切信息,精度）
	Ts            string `form:"ts"`             // 客户端发生广告点击事件的时间，以毫秒为单位时间戳
	CallbackParam string `form:"callback_param"` // 一些跟广告信息相关的回调参数，内容是一个加密字符串，在调用事件回传接口的时候会用到
	CallbackUrl   string `form:"callback_url"`   // 直接把调用事件回传接口的url生成出来，广告主可以直接使用
	UnionSite     string `form:"union_site"`     // 对外广告位编码
	Caid1         string `form:"caid1"`          // 不同版本的中国广告协会互联网广告标识，CAID1是20201230版，暂无CAID2
	PromotionId   int64  `form:"promotion_id"`   //巨量广告体验版中特有的宏参，代表巨量广告体验版的广告ID
	ProjectId     int64  `form:"project_id"`     //巨量广告体验版中特有的宏参，代表巨量广告体验版的项目ID
	PromotionName string `form:"promotion_name"` //巨量广告体验版中的广告名称
	ProjectName   string `form:"project_name"`   //巨量广告体验版中的项目名称
	Mid1          string `form:"mid1"`           //针对巨量广告体验版，图片素材宏参数（下发原始素材id）
	Mid2          string `form:"mid2"`           //针对巨量广告体验版，标题素材宏参数（下发原始素材id）
	Mid3          string `form:"mid3"`           //针对巨量广告体验版，视频素材宏参数（下发原始素材id）
	Mid4          string `form:"mid4"`           //针对巨量广告体验版，搭配试玩素材宏参数（下发原始素材id）
	Mid5          string `form:"mid5"`           //针对巨量广告体验版，落地页素材宏参数（下发原始素材id）
	Mid6          string `form:"mid6"`           //针对巨量广告体验版，安卓下载详情页素材宏参数（下发原始素材id）
	IdfaMd5       string `form:"idfa_md5"`       //IOS 6+的设备id的md5摘要，32位	注意，用户关闭读取idfa权限，0值也会进行MD5加密
	Ipv4          string `form:"ipv4"`           //优先使用上报请求的对端 IP 地址。如果该IP为 IPv6, 则使用客户端获取的 client_ipv4 地址
	Ipv6          string `form:"ipv6"`           //优先使用上报请求的对端 IP 地址。如果该IP为 IPv4, 则使用客户端获取的 client_ipv6 地址
	Caid          string `form:"caid"`           //中国广告协会互联网广告标识，包含最新两个版本的CAID和版本号，url encode之后的json字符串	(【CAID】和【CAID1、CAID2】的信息一致，使用一种即可；建议使用【CAID】，参数中包含多个信息，后续维护成本低）
	Caid2         string `form:"caid2"`          //不同版本的中国广告协会互联网广告标识，CAID1是20220111版，CAID2是20211207版
	Caid1Md5      string `form:"caid1_md5"`      //不同版本的中国广告协会互联网广告标识，CAID1是20220111版，CAID2是20211207版
	Caid2Md5      string `form:"caid2_md5"`      //不同版本的中国广告协会互联网广告标识，CAID1是20220111版，CAID2是20211207版
}

//func (this *JrttValidation) Validation(c *gin.Context) JrttValidation {
//	var jrttValidation JrttValidation
//	if err := c.ShouldBind(&jrttValidation); err != nil {
//		errs, ok := err.(validator.ValidationErrors)
//		if !ok {
//			panic(new(Tools.ExceptionHandle).ValidationException(err.Error()))
//		}
//		for _, v := range errs.Translate(Tools.Trans) {
//			panic(new(Tools.ExceptionHandle).ValidationException(v))
//		}
//	}
//	return jrttValidation
//}
