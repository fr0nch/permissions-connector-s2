package permissions

// Generated from permissions

type Status int32

const (
	Status_Success              Status = 0
	Status_Allow                Status = 1
	Status_Disallow             Status = 2
	Status_PermNotFound         Status = 3
	Status_PermAlreadyGranted   Status = 4
	Status_CookieNotFound       Status = 5
	Status_OptionNotFound       Status = 5
	Status_GroupNotFound        Status = 6
	Status_ChildGroupNotFound   Status = 7
	Status_ParentGroupNotFound  Status = 8
	Status_ActorUserNotFound    Status = 9
	Status_TargetUserNotFound   Status = 10
	Status_GroupAlreadyExist    Status = 11
	Status_UserAlreadyExist     Status = 12
	Status_TemporalGroup        Status = 13
	Status_PermanentGroup       Status = 14
	Status_GroupNotDefined      Status = 15
	Status_CallbackInvalid      Status = 16
	Status_CallbackAlreadyExist Status = 17
	Status_CallbackNotFound     Status = 18
	Status_StorageError         Status = 20
	Status_DBNotReady           Status = 21
)

type Action int32

const (
	Action_Add         Action = 0
	Action_Remove      Action = 1
	Action_Replace     Action = 2
	Action_ReplaceToWC Action = 3
)

type PlayerState uint32

const (
	PlayerState_NotFound PlayerState = 0
	PlayerState_Online   PlayerState = 1
	PlayerState_Offline  PlayerState = 2
)

type PermSource uint32

const (
	PermSource_UserTemp  PermSource = 0
	PermSource_User      PermSource = 1
	PermSource_GroupTemp PermSource = 2
	PermSource_Group     PermSource = 3
	PermSource_NotFound  PermSource = 4
)
