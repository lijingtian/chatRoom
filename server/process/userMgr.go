package process

import (
	"fmt"
)

type  UserMgr struct {
	onlineUsers map[int]*UserProcess
}
/*	UserMgr实例在服务器端有且只有一个
 *	因为要在很多地方用到，所以声明为全局的
*/
var(
	userMgr *UserMgr
)

func init(){
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (this *UserMgr) AddOnlineUser(up *UserProcess){
	this.onlineUsers[up.UserID] = up
}

func (this *UserMgr) DeleteOnlineUser(up *UserProcess){
	delete(this.onlineUsers, up.UserID)
}

//返回当前所有的在线用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess{
	return this.onlineUsers
}

//根据ID返回对应的值
func (this *UserMgr) GetOnlineUserByID(id int) (up *UserProcess, err error){
	up, ok := this.onlineUsers[id]
	if ok{
		return up, nil
	} else {
		return nil, fmt.Errorf("用户%d不存在", id)
	}
}