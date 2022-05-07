package main

import (
	"fmt"
	"suzaku/examples/group_chat/client"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mysql"
	"sync"
	"time"
)

const (
	TestGroupID = "666888"
)

func main() {
	var wg sync.WaitGroup
	var userIDs = []string{"123", "456", "789", "1011", "1213", "1415"}
	wg.Add(1)
	JoinGroup(userIDs, TestGroupID)
	manager := client.NewManager()
	manager.Run(userIDs, TestGroupID)
	wg.Wait()
}

func JoinGroup(userIDs []string, groupId string) {
	var (
		group  *po_mysql.Group
		gp     *po_mysql.Group
		userID string
		mb     *po_mysql.GroupMember
		member *po_mysql.GroupMember
		err    error
	)
	group = &po_mysql.Group{
		GroupId:       groupId,
		GroupName:     "朱雀",
		Notification:  "",
		Introduction:  "",
		AvatarUrl:     "",
		CreatorUserId: "",
		GroupType:     1,
		Status:        1,
		CreatedTs:     time.Now().Unix(),
		Ex:            "",
	}

	gp, err = repo_mysql.GroupRepo.GroupExist(group.GroupId)
	if err != nil {
		fmt.Println(err)
		return
	}
	if gp.GroupId == "" {
		err = repo_mysql.GroupRepo.Create(group)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for _, userID = range userIDs {
		member = &po_mysql.GroupMember{
			GroupId:        groupId,
			UserId:         userID,
			Nickname:       userID,
			UserAvatarUrl:  "",
			RoleLevel:      1,
			JoinedTs:       time.Now().Unix(),
			JoinSource:     0,
			OperatorUserId: "",
			MuteEndTs:      0,
			Ex:             "",
		}
		mb, err = repo_mysql.GroupRepo.IsJoined(member.GroupId, member.UserId)
		if err != nil {
			fmt.Println(err)
			return
		}
		if mb == nil {
			continue
		}
		if mb.UserId != "" {
			continue
		}
		err = repo_mysql.GroupRepo.Join(member)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
