package model

// "code":1, "message":"", "data":
type Response struct {
	Code    interface{} `json:"code"`
	Message interface{} `json:"message"`
	Data    Data        `json:"data"`
}

type JobOtherLists []JobOtherList
type jobDataLists []JobDataList
type JobConditionLists []JobConditionList
type EnrollFieldsLists []EnrollFieldsList

// Data
// interview 面试
// Enroll 报名
// Entrance 就职，入口
type Data struct {
	Notice     Notice       `json:"notice"`     //❌没用：雷达的公告
	CourseList []CourseList `json:"courseList"` //❌没用：课程列表

	JobCodeBase string `json:"jobCodeBase"` //部门代码
	JobCode     string `json:"jobCode"`     //职位代码

	Info                      Info                      `json:"info"`             //岗位信息
	JobDataList               jobDataLists              `json:"jobDataList"`      //职位信息 ②
	JobConditionList          JobConditionLists         `json:"jobConditionList"` //招考条件 ③
	JobOtherList              JobOtherLists             `json:"jobOtherList"`     //岗位其他信息 ④
	JobEnrollInfo             JobEnrollInfo             `json:"jobEnrollInfo"`    //✅APP端 ☆ 报名记录
	EnrollTitle               string                    `json:"enrollTitle"`      //"报名数据" or "报名预测"
	VipAuthList               VipAuthList               `json:"vipAuthList"`
	ForecastEnrollData        ForecastEnrollData        `json:"forecastEnrollData"` //✅预测报名人数
	ForecastEnrollText        ForecastEnrollText        `json:"forecastEnrollText"` //✅报名数据，报名预测，历年报名数据
	JobInterviewInfo          JobInterviewInfo          `json:"jobInterviewInfo"`   //✅APP端 ☆ 进面分数（最低进面分数，最高进面分数，平均进面分数）
	ShareInfo                 ShareInfo                 `json:"shareInfo"`
	UserInfo                  UserInfo                  `json:"userInfo"` //个人相关 是否合适
	TagList                   []TagList                 `json:"tagList"`  //岗位标签
	HistoryPageURL            string                    `json:"historyPageUrl"`
	HistoryEntranceInfo       interface{}               `json:"historyEntranceInfo"`
	Warning                   Warning                   `json:"warning"`
	IntendJobCount            int                       `json:"intendJobCount"` //预测岗位数
	ExamForecastScheme        string                    `json:"examForecastScheme"`
	ExamForecastText          string                    `json:"examForecastText"`
	SharePointInfo            SharePointInfo            `json:"sharePointInfo"`            //❌没用：晒分相关
	HistoryInterviewScoreInfo HistoryInterviewScoreInfo `json:"historyInterviewScoreInfo"` //✅历年进面分 分数数据
	HistoryCompetitionData    HistoryCompetitionData    `json:"historyCompetitionData"`    //历年竞争者数据
	JobDataBar                JobDataBar                `json:"jobDataBar"`
	HighFlow                  int                       `json:"highFlow"`
	ExamMainBranchInfo        interface{}               `json:"examMainBranchInfo"`
	BranchArticleID           int                       `json:"branchArticleId"`      //❌没用：分支文章id
	HistoryEnrollInfo         HistoryEnrollInfo         `json:"historyEnrollInfo"`    //✅历年报名数据
	HistoryInterviewInfo      HistoryInterviewInfo      `json:"historyInterviewInfo"` //✅历史面试数据
}

type Notice struct {
	Text      string `json:"text"`
	SchemeURL string `json:"schemeUrl"`
	Type      int    `json:"type"`
	PopDesc   string `json:"popDesc"`
}
type CourseList struct {
	Cid    string `json:"cid"`
	Icon   string `json:"icon"`
	Name   string `json:"name"`
	Slogan string `json:"slogan"`
	URL    string `json:"url"`
}
type ArticleInfo struct {
	ExamCategoryID int `json:"examCategoryId"`
	Province       int `json:"province"`
	City           int `json:"city"`
	Region         int `json:"region"`
}
type Info struct {
	ExamTitle        string      `json:"examTitle"` //来自公告
	Position         string      `json:"position"`  //岗位名称
	JobID            int         `json:"jobId"`     //职位ID
	ViewsCount       int         `json:"viewsCount"`
	IntentionCount   int         `json:"intentionCount"` //感觉是 匹配度
	Hires            int         `json:"hires"`          //招录人数
	ArticleID        int         `json:"articleId"`
	ExamID           int         `json:"examId"` //公告ID
	Tags             []string    `json:"tags"`
	WorkCompany      string      `json:"workCompany"`     //招考单位
	WorkCompanyLink  string      `json:"workCompanyLink"` //招考单位具有链接功能，可以app内跳转查询到该单位的历年招录情况
	Education        string      `json:"education"`       //学历要求
	Description      string      `json:"description"`     //公考雷达免责声明
	RecommendDesc    string      `json:"recommendDesc"`
	IsRecommendTag   int         `json:"isRecommendTag"`
	RecommendTagName string      `json:"recommendTagName"`
	EnrollStatusTag  string      `json:"enrollStatusTag"`  // 结束报名/即将截止/即将报名
	EnrollStatusCode int         `json:"enrollStatusCode"` // -1
	IsNotRecommend   int         `json:"isNotRecommend"`
	NotRecommendTag  string      `json:"notRecommendTag"`
	ArticleInfo      ArticleInfo `json:"articleInfo"`
	HiresShow        string      `json:"hiresShow"`
}
type JobConditionList struct { //招考条件(大致包含：专业要求、学历要求、学历性质、学位要求、应届、性别要求、最低服务工作年限、服务基层项目、年龄要求等等)
	//				{
	//                "name": "专业要求",
	//                "value": "本科：法学类（0301）研究生：法学（0301），法律（0351）",
	//                "sort": 1
	//            },
	//            {
	//                "name": "学历要求",
	//                "value": "本科及以上",
	//                "sort": 2
	//            },
	//            {
	//                "name": "学位要求",
	//                "value": "取得相应学位",
	//                "sort": 4
	//            },
	//            {
	//                "name": "政治面貌",
	//                "value": "不限",
	//                "sort": 5
	//            },
	//            {
	//                "name": "应届",
	//                "value": "2023应届毕业生",
	//                "sort": 6
	//            },
	//            {
	//                "name": "性别要求",
	//                "value": "不限",
	//                "sort": 7
	//            },
	//            {
	//                "name": "服务基层项目",
	//                "value": "不限",
	//                "sort": 13
	//            },
	//            {
	//                "name": "年龄要求",
	//                "value": "18周岁以上、35周岁以下（1986年11月至2004年11月期间出生）",
	//                "sort": 14
	//            },
	//            {
	//                "name": "资格证书",
	//                "value": "通过国家司法考试或国家统一法律职业资格考试（A类）",
	//                "sort": 16
	//            },
	//            {
	//                "name": "备注",
	//                "value": "通过国家司法考试或国家统一法律职业资格考试（A类）。",
	//                "sort": 17
	//            }
	Name  string `json:"name"`
	Value string `json:"value"`
	Sort  int    `json:"sort"`
}
type JobDataList struct { //职位信息（大致包含：职位名称、报考地区、考试类型、职位介绍、职位代码、工作单位、单位性质等）
	Name  string `json:"name"`
	Value string `json:"value"`
	Sort  int    `json:"sort"`
}
type EnrollFieldsList struct {
	Name            string `json:"name"`  //缴费人数、最低进面分数、最高进面分数、平均进面分数
	Value           string `json:"value"` // 对应缴费人数：例如25
	DifferenceData  string `json:"differenceData"`
	Sort            int    `json:"sort"`
	UpdateStatus    string `json:"updateStatus"` //
	DifferenceText  string `json:"differenceText"`
	DifferenceValue int    `json:"differenceValue"`
}
type FloatingWindow struct {
	MessageText string `json:"messageText"` //考试竞争比17:1,总缴费人数可查看报名数据专题>
	SchemeURL   string `json:"schemeUrl"`
	Remark      string `json:"remark"`
}
type JobEnrollInfo struct {
	EnrollFieldsList   EnrollFieldsLists `json:"enrollFieldsList"` //缴费情况
	EnrollStatus       int               `json:"enrollStatus"`
	EnrollURL          string            `json:"enrollUrl"`      //跳转链接，公考报名大数据“https://www.gongkaoleida.com/v-activity/sign-up-for-big-data/index.html#/bigdata?examId=588063”
	LastUpdateTime     string            `json:"lastUpdateTime"` //最近一次更新时间
	Remark             string            `json:"remark"`         //报名数据更新时间为每日10点、14点、17点，数据来源于网络，仅供参考！
	SourceName         string            `json:"sourceName"`     //数据来源：来源网络（公考雷达整理发布）
	SourceURL          string            `json:"sourceUrl"`      //云岭先锋网网站报名系统
	SchemeURL          string            `json:"schemeUrl"`
	EntranceName       string            `json:"entranceName"`
	Tag                string            `json:"tag"`
	UpdateNote         string            `json:"updateNote"`
	LastUpdateTimeName string            `json:"lastUpdateTimeName"` //更新时间
	UpdateStatusTag    string            `json:"updateStatusTag"`
	FloatingWindow     FloatingWindow    `json:"floatingWindow"` //报名数据窗口下的一条数据
	FieldsTag          string            `json:"fieldsTag"`
	FieldsTagNotice    string            `json:"fieldsTagNotice"`
}
type VipAuthList struct {
	JobEnrollForecast         int `json:"jobEnrollForecast"`
	HistoryInterviewScoreInfo int `json:"historyInterviewScoreInfo"`
	HistoryCompetitionData    int `json:"historyCompetitionData"`
}
type Legend struct { //报名预测中的折线图
	SplitValue  int `json:"splitValue"`  //间隔值：30
	Max         int `json:"max"`         //最高值:150
	SplitNumber int `json:"splitNumber"` //分为5部分值。最终纵轴结果为：0,30,60,90,120,150
}

type JobInterviewInfo struct { //进面分数
	EnrollFieldsList   EnrollFieldsLists `json:"enrollFieldsList"`
	EnrollStatus       int               `json:"enrollStatus"`
	EnrollURL          string            `json:"enrollUrl"`      //公考雷达进面分数查询链接：https://ylxf.1237125.cn/Html/News/2024/4/8/444149.html
	LastUpdateTime     string            `json:"lastUpdateTime"` //更新时间
	Remark             string            `json:"remark"`         //备注：呈现为资格审查名单笔试成绩，部分职位为进入体能测评人员名单笔试成绩，仅供参考！
	SourceName         string            `json:"sourceName"`     //来源：云岭先锋（公考雷达整理发布）
	SourceURL          string            `json:"sourceUrl"`      //云岭先锋网址链接 https://ylxf.1237125.cn/Html/News/2024/4/8/444149.html
	SchemeURL          string            `json:"schemeUrl"`
	EntranceName       string            `json:"entranceName"`
	Tag                string            `json:"tag"`
	InterviewTitle     string            `json:"interviewTitle"` //进面分数线、考生晒分、理念进面分
	InterviewTips      string            `json:"interviewTips"`
	NamelistType       int               `json:"namelistType"`
	UpdateNote         string            `json:"updateNote"`
	LastUpdateTimeName string            `json:"lastUpdateTimeName"`
	UpdateStatusTag    string            `json:"updateStatusTag"`
	FieldsTag          string            `json:"fieldsTag"`
	FloatingWindow     interface{}       `json:"floatingWindow"` //进面分数线(同报名数据)下方的一段文字：本考试平均进面分：125:1,职位分数线可查看进面专题>
	FieldsTagNotice    string            `json:"fieldsTagNotice"`
}

type HistoryInterviewScoreInfo struct {
	HistoryInterviewScoreTitle  string         `json:"historyInterviewScoreTitle"` //eg:历年进面分
	HistoryInterviewScore       string         `json:"historyInterviewScore"`      //eg:49
	HistoryInterviewScoreText   string         `json:"historyInterviewScoreText"`  // eg:平均最低进面分数
	Remark                      string         `json:"remark"`
	FloatingWindow              FloatingWindow `json:"floatingWindow"`
	HistoryInterviewScoreMosaic int            `json:"historyInterviewScoreMosaic"`
}

type List struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Count     int    `json:"count"`
	Date      string `json:"date"`
	OrderTime int    `json:"orderTime"`
}
type OriginDailyData struct {
	LineText string `json:"lineText"` //缴费人数
	Color    string `json:"color"`    //颜色：例如#FF6044
	List     []List `json:"list"`
}
type ForecastDailyData struct { //每日预测
	LineText string `json:"lineText"` //预测过审人数:
	Color    string `json:"color"`
	List     []List `json:"list"`
}
type LastForecastData struct { //预测最终过审人数
	Name                   string `json:"name"` //预测最终过审人数
	Value                  string `json:"value"`
	Sort                   int    `json:"sort"`
	LeftName               string `json:"leftName"` //预测热度排名
	LeftValue              string `json:"leftValue"`
	DifferenceForecast     int    `json:"differenceForecast"`
	DifferenceRank         int    `json:"differenceRank"`
	DifferenceForecastText string `json:"differenceForecastText"`
	DifferenceRankText     string `json:"differenceRankText"`
}
type ForecastEnrollData struct {
	Legend                  Legend            `json:"legend"`          //报名预测中的折线图
	OriginDailyData         OriginDailyData   `json:"originDailyData"` //缴费人数 每天的详细数据
	IntervalEnrollDailyData interface{}       `json:"intervalEnrollDailyData"`
	ForecastDailyData       ForecastDailyData `json:"forecastDailyData"` //预测缴费人数
	LastForecastData        LastForecastData  `json:"lastForecastData"`  //预测最终缴费人数
	EnrollDays              int               `json:"enrollDays"`
	ForecastStatus          string            `json:"forecastStatus"` //更新中
	LastUpdateTime          string            `json:"lastUpdateTime"` //更新时间
	Title                   string            `json:"title"`          //每日报名趋势
	Remark                  string            `json:"remark"`         //备注
	Description             string            `json:"description"`
	FloatingWindow          interface{}       `json:"floatingWindow"`
}
type ForecastEnrollText struct {
	EnrollText             string `json:"enrollText"`             //报名数据标题
	ForecastText           string `json:"forecastText"`           //报名预测标题
	HistoryCompetitionText string `json:"historyCompetitionText"` //历年报名数据
}
type JobOtherList struct { //其他信息：例如单位级别、选调生职位类型、人民警察职位类型、
	Name  string `json:"name"`
	Value string `json:"value"`
	Sort  int    `json:"sort"`
}
type ShareInfo struct { //分享至微信、QQ、朋友圈等链接携带的信息
	ShareDesc  string `json:"shareDesc"`  //这个职位感觉很适合你，公考雷达为您分析，该职位你报考的成功率是……
	ShareTitle string `json:"shareTitle"` //职位标题：第一师十三团社保所招聘：公务员（事业单位工作人员） 职位2人
	ShareURL   string `json:"shareUrl"`   //分享链接：https://www.gongkaoleida.com/exam/job/180585
}

// 对应职位详情中第一个卡片中的信息
type UserInfo struct { //个人匹配情况
	CompeteScore           int    `json:"competeScore"`
	CompeteTag             string `json:"competeTag"` //匹配度：偏弱、中等
	FeedbackEntranceStatus int    `json:"feedbackEntranceStatus"`
	IsIntention            int    `json:"isIntention"`
	IsMatch                int    `json:"isMatch"`
	IsSubscribe            int    `json:"isSubscribe"`
	Star                   string `json:"star"` //几颗星
}
type TagList struct { //标签
	TagID     int    `json:"tagId"`
	SchemeURL string `json:"schemeUrl"`
	TagName   string `json:"tagName"` //有编制、行政编
}

type Warning struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	SchemeURL string `json:"schemeUrl"`
}
type ShowConfig struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Color       string `json:"color"`
	LinkURL     string `json:"linkUrl"`
	ErrMsg      string `json:"errMsg"`
	TrackingKey string `json:"trackingKey"`
	Remark      string `json:"remark"`
}
type SharePointInfo struct { //邀请好友晒分
	NoticeText     string       `json:"noticeText"`
	InviteText     string       `json:"inviteText"`
	ButtonText     string       `json:"buttonText"`
	ButtonURL      string       `json:"buttonUrl"`
	JobCount       string       `json:"jobCount"`
	NoShowText     string       `json:"noShowText"`
	ShowConfig     []ShowConfig `json:"showConfig"`
	TrackingKey    string       `json:"trackingKey"`
	RankingCheckOn int          `json:"rankingCheckOn"`
	FloatingWindow interface{}  `json:"floatingWindow"`
}
type HistoryCompetitionData struct {
	EnrollRatio          string         `json:"enrollRatio"`
	EnrollRatioText      string         `json:"enrollRatioText"`
	SubscribeCount       string         `json:"subscribeCount"`
	SubscribeCountText   string         `json:"subscribeCountText"`
	Remark               string         `json:"remark"`
	FloatingWindow       FloatingWindow `json:"floatingWindow"`
	EnrollRatioMosaic    int            `json:"enrollRatioMosaic"`
	SubscribeCountMosaic int            `json:"subscribeCountMosaic"`
}
type JobDataBar struct {
	Title     string `json:"title"`
	BtnText   string `json:"btnText"`
	SchemeURL string `json:"schemeUrl"`
}

type HistoryEnrollInfo struct {
	TabTitle       string                  `json:"tabTitle"` //历年报名数据
	Remark         string                  `json:"remark"`   //（1）【历年报名数据】反映了本职位对应历年职位的竞争比中位数，【历年相似职位】指与本职位极其相似的职位，通常根据职位名称+工作单位来匹配。由于单位更名、职位新增等因素，会存在找不到历年职位的情况，此时无法提供相应数据。\n（2）竞争比中位数：是指对一组职位的竞争比（竞争比=报名数据/招录人数）求中位数，中位数是指从小到大排列的一组数据中，位于中间位置的数，如果该组数据为偶数个，则取中间2个数的平均值作为中位数
	Describe       string                  `json:"describe"` //以下数据为：职位竞争比中位数
	FloatingWindow FloatingWindow          `json:"floatingWindow"`
	List           []HistoryEnrollInfoList `json:"list"` // []{ {所属工作单位,1:1},{所属报考地区,1:1} }
}
type HistoryEnrollInfoList struct {
	Name              string `json:"name"`
	Value             string `json:"value"`
	EnrollRatioMosaic int    `json:"enrollRatioMosaic"`
}

type HistoryInterviewInfo struct {
	TabTitle       string         `json:"tabTitle"` //历年进面分
	Remark         string         `json:"remark"`   // eg:（1）【历年进面分/入围分】反映了本职位对应历年职位的最低进面/入围分数均值，【历年相似职位】指与本职位极其相似的职位，通常根据职位名称+工作单位来匹配。由于单位更名、职位新增等因素，会存在找不到历年职位的情况，此时无法提供相应数据。 \n（2）入围分数说明：官方有可能不公布进面人员名单，而公布面试后其他流程（如入围体检、入围体侧环节的）人员名单，此时公考雷达会根据其他人员名单进行整理发布，此类分数统称为“入围分数”。
	Describe       string         `json:"describe"` // eg:以下数据为：职位最低进面分均值
	FloatingWindow FloatingWindow `json:"floatingWindow"`
	List           []List         `json:"list"` // eg:{"name": "所属工作单位", "value": "56.8", "enrollRatioMosaic": 0},{"name": "所属报考地区", "value": "57.0", "enrollRatioMosaic": 0}
}

func (j JobOtherLists) Get(name string) string {
	for _, obj := range j {
		if obj.Name == name {
			//fmt.Println("get1", obj.Value)
			return obj.Value
		}
	}
	//fmt.Println("get1", "-")
	return ""
}

func (j jobDataLists) Get(name string) string {
	for _, obj := range j {
		if obj.Name == name {
			//fmt.Println("get2", obj.Value)
			return obj.Value
		}
	}
	//fmt.Println("get2", "-")
	return ""
}
func (j JobConditionLists) Get(name string) string {
	for _, obj := range j {
		if obj.Name == name {
			//fmt.Println("get3", obj.Value)
			return obj.Value
		}
	}
	//fmt.Println("get3", "-")
	return ""
}

func (j EnrollFieldsLists) Get(name string) string {
	for _, obj := range j {
		if obj.Name == name {
			return obj.Value
		}
	}
	return ""
}
