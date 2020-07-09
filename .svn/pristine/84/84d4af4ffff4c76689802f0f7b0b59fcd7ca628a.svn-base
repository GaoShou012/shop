package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	// 加密key
	key :=  []byte("aadlfkjaslkdfjlk")

	type User struct {
		Id int `json:"id"`
		Username string `json:"username"`
	}

	u := User{
		Id:       123,
		Username: "小小小",
	}
	m := jwt.MapClaims{}
	j,err := json.Marshal(&u)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(j,&m)
	if err != nil {
		panic(err)
	}
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
	//	"Id":"123",
	//	"Name":"小虎",
	//	"Exp":time.Now().Add(time.Hour * 1).Unix(),
	//})
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,m)
	tokenEncrypt,err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	fmt.Printf("token=%s\n",tokenEncrypt)

	tokenDecrypt,err := jwt.Parse(tokenEncrypt, func(token *jwt.Token) (i interface{}, err error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v\n", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return key, nil
	})
	if err != nil {
		panic(err)
	}

	// 把解密后的token转换类型
	if claims,ok := tokenDecrypt.Claims.(jwt.MapClaims); ok{
		fmt.Println(claims)
		fmt.Println(claims["Id"])
	}else{
		fmt.Println(err)
	}
}
