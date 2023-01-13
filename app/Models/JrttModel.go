package Models

type JrttModel struct {
	ID            uint   `gorm:"primaryKey;autoIncrement:true;column:id;type:int(int) unsigned AUTO_INCREMENT;not null"`
	AdvertiserId  string `gorm:"column:advertiser_id;type:varchar(100);not null;default:'';comment:广告主id" json:"advertiser_id"`
	Aid           string `gorm:"column:aid;type:varchar(100);not null;default:'';comment:广告计划id" json:"aid"`
	AidName       string `gorm:"column:aid_name;type:varchar(100);not null;default:'';comment:广告计划名称" json:"aid_name"`
	Cid           string `gorm:"column:cid;type:varchar(100);not null;default:'';comment:广告创意id" json:"cid"`
	CidName       string `gorm:"column:cid_name;type:varchar(100);not null;default:'';comment:广告创意名称" json:"cid_name"`
	CampaignId    string `gorm:"column:campaign_id;type:varchar(100);not null;default:'';comment:广告组id" json:"campaign_id"`
	Csite         string `gorm:"column:csite;type:varchar(100);not null;default:'';comment:广告投放位置" json:"csite"`
	Ip            string `gorm:"column:ip;type:varchar(100);not null;default:'';comment:媒体投放系统获取的用户终端的公共IP地址" json:"ip"`
	Model         string `gorm:"column:model;type:varchar(100);not null;default:'';comment:手机型号" json:"model"`
	CallbackUrl   string `gorm:"column:callback_url;type:varchar(3000);not null;default:'';comment:调用事件回传接口的url" json:"callback_url"`
	Imei          string `gorm:"column:imei;type:varchar(100);not null;default:'';comment:安卓手机 imei 的 md5 摘要" json:"imei"`
	IDfa          string `gorm:"column:idfa;type:varchar(100);not null;default:'';comment:ios 手机的 idfa 原值" json:"idfa"`
	Androidid     string `gorm:"column:androidid;type:varchar(100);not null;default:'';comment:安卓id原值的md5，32位" json:"androidid"`
	Oaid          string `gorm:"column:oaid;type:varchar(100);not null;default:'';comment:Android Q 版本的 oaid 原值" json:"oaid"`
	OaidMd5       string `gorm:"column:oaid_md5;type:varchar(100);not null;default:'';comment:Android Q 版本的 oaid 原值的md5摘要" json:"oaid_md5"`
	Os            uint8  `gorm:"column:os;type:tinyint unsigned;not null;default:0;comment:客户端的操作系统类型 0: android 1: ios" json:"os"`
	Channel       string `gorm:"column:channel;type:varchar(120);not null;default:'';comment:渠道来源" json:"channel"`
	CampaignName  string `gorm:"column:campaign_name;type:varchar(100);not null;default:'';comment:广告组名称" json:"campaign_name"`
	Ctype         int16  `gorm:"column:ctype;type:tinyint(4);not null;default:0;comment:广告组名称" json:"ctype"`
	ConvertId     int64  `gorm:"column:convert_id;type:varchar(100);not null;default:0;comment:转化id" json:"convert_id"`
	RequestId     string `gorm:"column:request_id;type:varchar(100);not null;default:'';comment:请求下发的id" json:"request_id"`
	Sl            string `gorm:"column:sl;type:varchar(100);not null;default:'';comment:这次请求的语言" json:"sl"`
	Mac           string `gorm:"column:mac;type:varchar(100);not null;default:'';comment:移动设备mac地址,转换成大写字母,去掉“:”，并且取md5摘要后的结果" json:"mac"`
	Mac1          string `gorm:"column:mac1;type:varchar(100);not null;default:'';comment:移动设备 mac 地址,转换成大写字母,并且取md5摘要后的结果，32位" json:"mac1"`
	Ua            string `gorm:"column:ua;type:varchar(1000);not null;default:'';comment:用户代理(User Agent)，一个特殊字符串头，使得服务器能够识别客户使用的操作系统及版本、CPU类型、浏览器及版本、浏览器渲染引擎、浏览器语言、浏览器插件等" json:"ua"`
	Geo           string `gorm:"column:geo;type:varchar(100);not null;default:'';comment:位置信息，包含三部分:latitude（纬度），longitude（经度）以及precise（确切信息,精度）" json:"geo"`
	Ts            string `gorm:"column:ts;type:int(11);not null;default:0;comment:客户端发生广告点击事件的时间，以毫秒为单位时间戳" json:"ts"`
	CallbackParam string `gorm:"column:callback_param;type:varchar(100);not null;default:'';comment:一些跟广告信息相关的回调参数，内容是一个加密字符串，在调用事件回传接口的时候会用到" json:"callback_param"`
	UnionSite     string `gorm:"column:union_site;type:varchar(100);not null;default:'';comment:对外广告位编码" json:"union_site"`
	Caid1         string `gorm:"column:caid1;type:varchar(255);not null;default:'';comment:不同版本的中国广告协会互联网广告标识，CAID1是20220111版，CAID2是20211207版" json:"caid1"`
	PromotionId   int64  `gorm:"column:promotion_id;type:varchar(50);not null;default:0;comment:巨量广告体验版中特有的宏参，代表巨量广告体验版的广告ID" json:"promotion_id"`
	ProjectId     int64  `gorm:"column:project_id;type:varchar(50);not null;default:0;comment:巨量广告体验版中特有的宏参，代表巨量广告体验版的项目ID" json:"project_id"`
	PromotionName string `gorm:"column:promotion_name;type:varchar(100);not null;default:'';comment:巨量广告体验版中的广告名称" json:"promotion_name"`
	ProjectName   string `gorm:"column:project_name;type:varchar(100);not null;default:'';comment:巨量广告体验版中的项目名称" json:"project_name"`
	Mid1          string `gorm:"column:mid1;type:varchar(100);not null;default:'';comment:针对巨量广告体验版，图片素材宏参数（下发原始素材id）" json:"mid1"`
	Mid2          string `gorm:"column:mid2;type:varchar(100);not null;default:'';comment:针对巨量广告体验版，标题素材宏参数（下发原始素材id）" json:"mid2"`
	Mid3          string `gorm:"column:mid3;type:varchar(100);not null;default:'';comment:针对巨量广告体验版，视频素材宏参数（下发原始素材id）" json:"mid3"`
	Mid4          string `gorm:"column:mid4;type:varchar(100);not null;default:'';comment:针对巨量广告体验版，搭配试玩素材宏参数（下发原始素材id）" json:"mid4"`
	Mid5          string `gorm:"column:mid5;type:varchar(100);not null;default:'';comment:针对巨量广告体验版，落地页素材宏参数（下发原始素材id）" json:"mid5"`
	Mid6          string `gorm:"column:mid6;type:varchar(100);not null;default:'';comment:针对巨量广告体验版，安卓下载详情页素材宏参数（下发原始素材id）" json:"mid6"`
	IdfaMd5       string `gorm:"column:idfa_md5;type:varchar(100);not null;default:'';comment:IOS 6+的设备id的md5摘要，32位,注意，用户关闭读取idfa权限，0值也会进行MD5加密" json:"idfa_md5"`
	Ipv4          string `gorm:"column:ipv4;type:varchar(50);not null;default:'';comment:优先使用上报请求的对端 IP 地址。如果该IP为 IPv6, 则使用客户端获取的 client_ipv4 地址" json:"ipv4"`
	Ipv6          string `gorm:"column:ipv6;type:varchar(50);not null;default:'';comment:优先使用上报请求的对端 IP 地址。如果该IP为 IPv4, 则使用客户端获取的 client_ipv6 地址" json:"ipv6"`
	Caid          string `gorm:"column:caid;type:varchar(255);not null;default:'';comment:中国广告协会互联网广告标识，包含最新两个版本的CAID和版本号，url encode之后的json字符串" json:"caid"`
	Caid2         string `gorm:"column:caid2;type:varchar(255);not null;default:'';comment:不同版本的中国广告协会互联网广告标识，CAID1是20220111版，CAID2是20211207版" json:"caid2"`
	Caid1Md5      string `gorm:"column:caid1_md5;type:varchar(100);not null;default:'';comment:不同版本的中国广告协会互联网广告标识，CAID1是20220111版，CAID2是20211207版" json:"caid1_md5"`
	Caid2Md5      string `gorm:"column:caid2_md5;type:varchar(100);not null;default:'';comment:不同版本的中国广告协会互联网广告标识，CAID1是20220111版，CAID2是20211207版" json:"caid2_md5"`
	CreatedAt     uint   `gorm:"autoCreateTime;column:created_at;type:int(11);comment:创建时间" json:"created_at"`
	UpdatedAt     uint   `gorm:"autoUpdateTime;column:updated_at;type:int(11);comment:更新时间" json:"updated_at"`
}
