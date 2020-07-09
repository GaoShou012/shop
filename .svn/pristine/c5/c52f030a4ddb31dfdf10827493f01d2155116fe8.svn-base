package models

type Admin struct {
	Model
	ThumbId *int					`json:"thumbId"`
	Username *string				`json:"username"`
	Password *string				`json:"password"`
	Nickname *string				`json:"nickname"`
}

//var ModelAdmin = modelAdmin{
//	Model{TableName:"admin"},
//}
//type modelAdmin struct {
//	Model
//}

// bool 返回校验结果
// error 返回是否有执行错误
//func (m *modelAdmin) AuthVerify(username string , password string) (*Admin,error) {
//	user := &Admin{}
//
//	res := DB.Table(m.TableName).Where("username= ? and password = ?",username,password).First(user)
//	if res.RecordNotFound() {
//		return nil,fmt.Errorf("校验失败，账号或密码错误")
//	}
//	if res.Error != nil {
//		return nil,res.Error
//	}
//
//	return user,nil
//}

