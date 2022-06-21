package test_model

//通用返回信息
type ResponseInfo struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

//=================================请求信息====================================
type LoginRequestInfo struct {
	Name string `json:"name"`
	Pswd string `json:"pswd"`
}

type WhichAllianceRequestInfo struct {
	Name string `json:"name"`
}

type AllianceList struct {
}

type CreateAllianceRequestInfo struct {
	CName string `json:"cname"`
	AName string `json:"aname"`
}

type JoinAllianceRequestInfo struct {
	CName string `json:"cname"`
	AName string `json:"aname"`
}

type DismissAllianceRequestInfo struct {
	CName string `json:"cname"`
}

type IncreaseCapacityRequestInfo struct {
	CName string `json:"cname"`
}

type StoreItemRequestInfo struct {
	CName   string `json:"cname"`
	ItemId  int32  `json:"item_id"`
	ItemNum int32  `json:"item_num"`
	Index   int32  `json:"index"`
}

type DestoryItemRequestInfo struct {
	CName string `json:"cname"`
	Index int32  `json:"index"`
}

type ClearUpRequestInfo struct {
	CName string `json:"cname"`
}
