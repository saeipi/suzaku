// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
)

// InboxReadChatMediaUnreadToInbox
// inbox.readChatMediaUnreadToInbox from_id:long peer_chat_id:long id:Vector<int> = Void;
func (c *InboxCore) InboxReadChatMediaUnreadToInbox(in *inbox.TLInboxReadChatMediaUnreadToInbox) (*mtproto.Void, error) {
	c.svcCtx.Dao.MessagesDAO.SelectPeerDialogMessageIdListWithCB(
		c.ctx,
		in.FromId,
		in.Id,
		func(i int, v *dataobject.MessagesDO) {
			c.svcCtx.Dao.MessagesDAO.UpdateMediaUnread(c.ctx, v.UserId, v.UserMessageBoxId)

			// TODO: batch handle
			pts := c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, v.UserId)
			c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
				UserId: v.UserId,
				Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadMessagesContents(&mtproto.Update{
					Messages:  []int32{v.UserMessageBoxId},
					Pts_INT32: pts,
					PtsCount:  1,
				}).To_Update()),
			})
		},
	)

	return mtproto.EmptyVoid, nil
}
