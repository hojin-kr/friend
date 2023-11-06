// BEGIN: 4d5f6a8b3c2d
package friend_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/hojin-kr/friend"
	"github.com/stretchr/testify/assert"
)

func TestRequestFriend(t *testing.T) {
	uuid := uuid.New().String()
	_friend := &friend.Friend{
		UUID_Send:    "test",
		UUID_Receive: "test2",
		Status:       friend.PENDING,
		Friend_ID:    uuid,
	}
	_friend.PendingFriend()

	assert.Equal(t, friend.PENDING, _friend.Status)
}

func TestAcceptFriend(t *testing.T) {
	uuid := "a20119db-c9f0-4a33-9322-5db6095e181b"
	_friend := &friend.Friend{
		UUID_Send:    "test",
		UUID_Receive: "test2",
		Status:       friend.PENDING,
		Friend_ID:    uuid,
	}
	_friend.AcceptFriend()

	assert.Equal(t, friend.ACCEPTED, _friend.Status)
}

func TestFiilterdList(t *testing.T) {
	_friend := &friend.Friend{}
	friends := _friend.FilteredList("Status", "=", friend.PENDING, 100)
	t.Logf("%+v\n", friends)
	assert.NotNil(t, friends)
}

// END: 4d5f6a8b3c2d
