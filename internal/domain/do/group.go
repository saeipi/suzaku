package do

import "suzaku/internal/domain/po_mysql"

type JoinGroupResult struct {
	Member       *po_mysql.GroupMember
	Group        *po_mysql.Group
	GroupRequest *po_mysql.GroupRequest
	HandleResult int32
}
