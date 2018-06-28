package modelcall

type modelcall struct {
	reportData map[string]interface{}
}

func (this *modelcall) startModCall() {
}

func (this *modelcall) endModCall(resp map[string]interface{}) {
}

func (this *modelcall) getHostName() {

}

/* 外部结构体
type UserModel struct {
	Mocall modelcall
	OtherData xxxx
}

u:=UserModel{}
u.startModCall
resp := getResp()

u.endModCall(resp)

*/
