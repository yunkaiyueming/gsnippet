package pod_event

type EVENT_ITEM string

const (
	DEPLOYMENT_ITEM = "Deployment"
	POD_ITEM        = "Pod"

	DEP_CREATE         = "Dep_Create"
	DEP_UPGRADE_CREATE = "Dep_Upgrade_Create"
	DEP_STOP           = "Dep_Stop"
	DEP_UPDATE         = "Dep_Update"
	DEP_PAUSE          = "Dep_Pause"
	DEP_START          = "Dep_Start"

	POD_START    EVENT_ITEM = "Pod_Start"
	POD_STOP     EVENT_ITEM = "Pod_Stop"
	POD_PAUSE    EVENT_ITEM = "Pod_Pause"
	POD_FAILED   EVENT_ITEM = "Pod_Failed"
	POD_TIMEOUT2 EVENT_ITEM = "Pod_Timeout2"
	POD_TASKADD  EVENT_ITEM = "Pod_TaskAdd"
	POD_TASKDEL  EVENT_ITEM = "Pod_TaskDel"
)

const LOG_BASE_PATH = "E:/GO_PATH/src/gsnippet/pod_event" // "/data/log/iaas"
const F_DATE = "20060102"                                 //长日期格式
const F_DATETIME = "2006-01-02 15:04:05"                  //日期时间格式
const FILENAME = "deployment_event.log."
