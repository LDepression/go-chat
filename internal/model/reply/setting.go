package reply

type GetFriendInfo struct {
	Account EasyAccount
	Setting EasySetting
}

type GetFriendsList struct {
	FriendsInfos []*GetFriendInfo
	Total        int
}

type GetFriendsByName struct {
	FriendsInfos []*GetFriendInfo
	Total        int
}
