package model

//mysql存储Model层
type User struct {
	ID        int64  `db:"id"`        //用户ID
	Password  string `db:"password"`  //密码
	NickName  string `db:"nickName"`  //用户名
	AvatarURL string `db:"avatarURL"` //头像
}
