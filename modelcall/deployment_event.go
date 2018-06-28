package modelcall

const DEPLOYMENT_ITEM = "Deployment"
const POD_ITEM = "Pod"
const DEP_CREATE = "Dep_Create"
const DEP_UPGRADE_CREATE = "Dep_Upgrade_Create"
const DEP_STOP = "Dep_Stop"
const DEP_UPDATE = "Dep_Update"
const DEP_PAUSE = "Dep_Pause"
const DEP_START = "Dep_Start"

type DeploymentEvent struct {
}

//主动给对外调用生成一个实例
func newDeploymentEvent() *DeploymentEvent {
	&DeploymentEvent{}
}

//类调用
func (this *DeploymentEvent) OperationLog() {
}
func (this *DeploymentEvent) _checkData() {
}
func (this *DeploymentEvent) _log() {
}
func (this *DeploymentEvent) coredumpLog() {

}

//或者写成，包调用,不用生成类
func coredumpLog() {
}
