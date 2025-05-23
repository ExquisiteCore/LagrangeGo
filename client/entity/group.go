package entity

import "fmt"

type (
	EventState          uint32
	EventType           uint32
	GroupRequestOperate uint8 // 处理加群请求
)

const (
	GroupRequestOperateAllow  GroupRequestOperate = 1 // 同意
	GroupRequestOperateDeny   GroupRequestOperate = 2 // 拒绝
	GroupRequestOperateIgnore GroupRequestOperate = 3 // 忽略
)

const (
	/*
		1:未处理
		2:已同意
		3:已拒绝
		4:已忽略
		5:已处理
	*/
	NoNeed EventState = iota
	Unprocessed
	Processed
)

const (
	/*
		1:“用户昵称”申请加入“群名称”群。
		2:“用户昵称”邀请您加入“群名称”群。
		3：成为管理员，通知全员。
		6：UIN被T 出群，通知管理员。
		7：UIN被T 出群，通知UIN。
		10:您的好友“用户昵称”拒绝加入“群名称”群。拒绝理由：XXX
		11:“群名称”群管理员“用户昵称”拒绝了您的加群请求。拒绝理由：XXX
		12:您的好友“用户昵称”已经同意加入“群名称”群。
		13：UIN退群，通知管理员。
		15：管理员身份被取消，通知被取消人。
		16：管理员身份被取消，通知其他管理员。
		20(同2，已废弃）:	您的好友“用户昵称”邀请您加入“群名称”群。
		21（同12，已废弃）:	您的好友“用户昵称”已经同意加入“群名称”群。附加信息：正在等待管理员验证。
		22:“用户昵称”申请加入“群名称”群。附加信息：来自群成员XXX的邀请。
		23（同10，已废弃）:	您的好友“用户昵称”拒绝加入“群名称”群。拒绝理由：XXX。
		35:群“群名称”管理员已同意您的加群申请
	*/
	// UserJoinRequest 用户申请加群
	UserJoinRequest EventType = 1
	// GroupInvited 被邀请加群
	GroupInvited EventType = 2
	// AssignedAsAdmin 被设置为管理员
	AssignedAsAdmin EventType = 3
	// Kicked 被踢出群聊
	Kicked EventType = 7
	// RemoveAdmin 被取消管理员
	RemoveAdmin EventType = 15
	// UserInvited 群员邀请其他人
	UserInvited EventType = 22
)

type (
	Group struct {
		GroupUin        uint32
		GroupName       string
		GroupOwner      uint32
		GroupCreateTime uint32
		GroupMemo       string
		GroupLevel      uint32
		MemberCount     uint32
		MaxMember       uint32
		LastMsgSeq      uint32
	}

	UserJoinGroupRequest struct {
		GroupName   string     `json:"group_name"`
		GroupUin    uint32     `json:"group_id"`
		InvitorUin  uint32     `json:"invitor_uin"`
		InvitorUID  string     `json:"-"`
		TargetNick  string     `json:"requester_nick"`
		TargetUin   uint32     `json:"requester_uin"`
		TargetUID   string     `json:"-"`
		OperatorUin uint32     `json:"actor"`
		OperatorUID string     `json:"-"`
		Sequence    uint64     `json:"request_id"`
		Checked     bool       `json:"checked"`
		State       EventState `json:"-"`
		EventType   EventType  `json:"-"`
		Comment     string     `json:"message"`
		IsFiltered  bool       `json:"is_filtered"`
	}

	GroupInvitedRequest struct {
		GroupUin    uint32     `json:"group_id"`
		GroupName   string     `json:"group_name"`
		InvitorNick string     `json:"invitor_nick"`
		InvitorUin  uint32     `json:"invitor_uin"`
		InvitorUID  string     `json:"-"`
		Sequence    uint64     `json:"request_id"`
		Checked     bool       `json:"checked"`
		State       EventState `json:"-"`
		EventType   EventType  `json:"-"`
		IsFiltered  bool       `json:"is_filtered"`
	}

	GroupSystemMessages struct {
		InvitedRequests []*GroupInvitedRequest  `json:"invited_requests"`
		JoinRequests    []*UserJoinGroupRequest `json:"join_requests"`
	}
)

func (g *Group) Avatar() string {
	return fmt.Sprintf("https://p.qlogo.cn/gh/%d/%d/0/", g.GroupUin, g.GroupUin)
}
