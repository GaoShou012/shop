package lib

//func PostFormNum(context *gin.Context,key string,defaultValue uint) (uint,error) {
//	val,ok := context.GetPostForm(key)
//
//	// 适应前端，如果是undefined就匹配为未定义
//	if val == "undefined" {
//		ok = false
//	}
//
//	if val == "" {
//		ok = false
//	}
//
//	if !ok {
//		return defaultValue,nil
//	}
//
//	num,err := strconv.Atoi(val)
//	if err != nil {
//		return 0,err
//	}
//
//	return num,err
//}