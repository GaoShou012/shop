package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/app"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"github.com/proto"
	"time"
)

/*

生成订单号

*/

const (
	twepoch        = int64(1483228800000)             //开始时间截 (2017-01-01)
	workeridBits   = uint(0)                          //机器id所占的位数
	sequenceBits   = uint(12)                         //序列所占的位数
	workeridMax    = int64(-1 ^ (-1 << workeridBits)) //支持的最大机器id数量
	sequenceMask   = int64(-1 ^ (-1 << sequenceBits)) //
	workeridShift  = sequenceBits                     //机器id左移位数
	timestampShift = sequenceBits + workeridBits      //时间戳左移位数
)

type OrderNoMicroService struct {
	key []byte
	timestamp 	int64
	sequence  	int64
}

func (s *OrderNoMicroService) Encrypt(ctx context.Context,req *proto.OrderNoEncryptReq,rsp *proto.OrderNoEncryptRsp) error {
	// 把数据转成 jwt map格式
	mapClaims := jwt.MapClaims{}
	j,err := json.Marshal(req.Data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(j,&mapClaims)
	if err != nil {
		return err
	}

	// 加密数据
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,mapClaims)
	str,err := token.SignedString(s.key)
	if err != nil {
		return err
	}

	// 返回加密的数据
	rsp.Token = str
	return nil
}

func (s *OrderNoMicroService) Decrypt(ctx context.Context,req *proto.OrderNoDecryptReq,rsp *proto.OrderNoDecryptRsp) error {
	token,err := jwt.Parse(req.Token, func(token *jwt.Token) (i interface{}, err error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("Unexpected signing method: %v\n",token.Header["alg"])
		}
		return []byte(s.key),nil
	})
	if err != nil {
		return err
	}

	claims,ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("订单号解码类型转换失败")
	}
	err = mapstructure.WeakDecode(claims,&rsp.Data)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderNoMicroService) Gen(ctx context.Context, req *proto.OrderNoGenRequest, rsp *proto.OrderNoGenResponse) error {
	no,err := s.gen()
	if err != nil {
		rsp.Code = 1
		rsp.Message = err.Error()
	}
	rsp.No = no
	return nil
}

func (s *OrderNoMicroService) gen() (uint64,error) {
	now := time.Now().Unix()
	if s.timestamp == now {
		sequence := (s.sequence + 1) & sequenceMask
		if sequence == 0 {
			return 0,fmt.Errorf("系统繁忙未能生成订单号")
		}else{
			s.sequence = sequence
		}
	}else{
		s.timestamp = now
		s.sequence = 0
	}

	return uint64(now << timestampShift) | uint64(s.sequence) , nil
}

func main() {
	service := app.NewService("micro.service.order.no")
	serviceHandler := OrderNoMicroService{key:[]byte("Uc19SEWFRiCVnAZU")}
	if err := proto.RegisterOrderNoHandler(service.Server(),&serviceHandler); err != nil {
		panic(err)
	}
	if err := service.Run(); err != nil {
		panic(err)
	}
}
