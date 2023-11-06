package friend

import (
	"github.com/google/uuid"
	"github.com/hojin-kr/datastore"
)

type Friend struct {
	UUID_Send    string `json:"uuid_send"`
	UUID_Receive string `json:"uuid_receive"`
	Status       string `json:"status"`
	Friend_ID    string `json:"friend_id"`
}

const (
	PENDING  = "pending"
	ACCEPTED = "accepted"
	REJECTED = "rejected"
	BLOCKED  = "blocked"
)

func (friend *Friend) PendingFriend() {
	friend.Status = PENDING
	// uuid
	friend.Friend_ID = uuid.New().String()
	// put datastore
	datastore := datastore.GcpDatastore{}
	datastore.Init()
	datastore.PutEntity(friend.Friend_ID, friend)
}

func (friend *Friend) AcceptFriend() {
	friend.Status = ACCEPTED
	// get datastore
	datastore := datastore.GcpDatastore{}
	datastore.Init()
	datastore.PutEntity(friend.Friend_ID, friend)
}

func (friend *Friend) RejectFriend() {
	friend.Status = REJECTED
	// get datastore
	datastore := datastore.GcpDatastore{}
	datastore.Init()
	datastore.PutEntity(friend.Friend_ID, friend)
}

func (friend *Friend) BlockFriend() {
	friend.Status = BLOCKED
	// get datastore
	datastore := datastore.GcpDatastore{}
	datastore.Init()
	datastore.PutEntity(friend.Friend_ID, friend)
}

// filterdlist
func (friend *Friend) FilteredList(colume string, operation string, value string, limit int) (friends interface{}) {
	datastore := datastore.GcpDatastore{}
	datastore.Init()
	var _friend Friend
	friends = datastore.FilteredList(&_friend, colume, operation, value, limit)
	return friends
}
