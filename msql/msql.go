package msql

import (
	"container/list"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"
	"time"
)

const (
	MysqlComponentErrorPrefix = "msql error:"
)

const (
	USERNAME = "root"
	PASSWORD = "Forever634312."
	NETWORK  = "tcp"
	SERVER   = "192.168.0.200"
	PORT     = 3306
	DATABASE = "server"
)

var DB *sql.DB

func init() {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	db,err := sql.Open("mysql",dsn)
	if err != nil{
		fmt.Printf("Open mysql failed,err:%v\n",err)
		return
	}
	db.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
	db.SetMaxOpenConns(100)				//设置最大连接数
	db.SetMaxIdleConns(16) 				//设置闲置连接数

	DB = db
	fmt.Println("sql connected")
}

type Where struct {
	Key string
	Operator string
	Val interface{}
}

type Condition struct {
	Key string
	Operator string
	Val interface{}
}

type Conditions struct {
	init sync.Once
	li *list.List
}

func(c *Conditions) Push(key string,operator string,val interface{}) {
	c.init.Do(func() {
		c.li = list.New()
	})

	w := &Condition{
		Key:      key,
		Operator: operator,
		Val:      val,
	}

	c.li.PushBack(w)
}

func (c *Conditions) GetList() *list.List {
	c.init.Do(func() {
		c.li = list.New()
	})
	return c.li
}

type Msql struct {
	builderCache 		*list.List
	builderParams 		*list.List

	buildingFailed bool
	buildingFailedError error

	QueryString 		string
}

func (m *Msql) Select(val interface{}) *Msql {
	msql := &Msql{builderCache:list.New(),builderParams:list.New()}
	msql.builderCache.PushBack("select")

	switch val.(type) {
	case string:
		msql.builderCache.PushBack(val.(string))
		break
	case []string:
		tmp := strings.Join(val.([]string),"`,`")
		str := fmt.Sprintf("`%s`",tmp)
		msql.builderCache.PushBack(str)
		break
	default:
		log.Println("msql select unknown type")
		break
	}

	return msql
}
func (m *Msql) Insert() {}
func (m *Msql) Update() {}
func (m *Msql) Delete() {}

func (m *Msql) From(table string) *Msql {
	m.builderCache.PushBack("from")
	m.builderCache.PushBack(fmt.Sprintf("`%s`",table))
	return m
}

func (m *Msql) condition(c *Condition) error {
	switch c.Operator {
	case "in":
		str := fmt.Sprintf("`%s` in (?)",c.Key)
		//m.cache.PushBack(str)
		m.builderCache.PushBack(str)
		//m.params.PushBack(join(v.Val0))
		break
	case "not in":
		//str := fmt.Sprintf("`%s` not in (?)",v.Key)
		//m.cache.PushBack(str)
		//m.params.PushBack(join(v.Val0))
		break
	case "between":
		str := fmt.Sprintf("`%s` between ? and ?",c.Key)
		m.builderCache.PushBack(str)
		switch c.Val.(type) {
		case [2]int:
			params := c.Val.([2]int)
			m.builderParams.PushBack(params[0])
			m.builderParams.PushBack(params[1])
			break
		case [2]string:
			params := c.Val.([2]string)
			m.builderParams.PushBack(params[0])
			m.builderParams.PushBack(params[1])
			break
		default:
			return fmt.Errorf("%s between转换类型失败",MysqlComponentErrorPrefix)
		}
		break
	case "not between":
		str := fmt.Sprintf("`%s` not between ? and ?",c.Key)
		m.builderCache.PushBack(str)
		switch c.Val.(type) {
		case [2]int:
			params := c.Val.([2]int)
			m.builderParams.PushBack(params[0])
			m.builderParams.PushBack(params[1])
			break
		case [2]string:
			params := c.Val.([2]string)
			m.builderParams.PushBack(params[0])
			m.builderParams.PushBack(params[1])
			break
		default:
			return fmt.Errorf("%s not between转换类型失败",MysqlComponentErrorPrefix)
		}
		break
		// =,>,<,>=,<=
	default:
		str := fmt.Sprintf("`%s` %s ?",c.Key,c.Operator)
		m.builderCache.PushBack(str)
		switch c.Val.(type) {
		case int:
			m.builderParams.PushBack(c.Val.(int))
			break
		case string:
			m.builderParams.PushBack(c.Val.(string))
			break
		default:
			return fmt.Errorf("%s 运算符%s 转换类型识别",MysqlComponentErrorPrefix,c.Operator)
		}
		break
	}

	return nil
}

func (m *Msql) Where(conditions interface{}) *Msql {
	m.builderCache.PushBack("where")

	switch conditions.(type) {
	case string:
		m.builderCache.PushBack(conditions.(string))
		break
	case *Condition:
		err := m.condition(conditions.(*Condition))
		if err != nil {
			log.Println(err)
		}
	case *Conditions:
		li := conditions.(*Conditions).GetList()
		count := li.Len()

		for i := li.Front(); i != nil; i=i.Next() {
			err := m.condition(i.Value.(*Condition))
			if err != nil {
				log.Println(err)
			}
			if count--;count >0 {
				m.builderCache.PushBack(" and ")
			}
		}
		break
	default:
		log.Println("不支持的conditions类型")
	}
	return m
}

func (m *Msql) QueryRow(container interface{}) {
	//cache := make([]string,0,m.builderCache.Len())
	//for i := m.builderCache.Front(); i != nil; i = i.Next() {
	//	cache = append(cache,i.Value.(string))
	//}
	//
	//queryString :=  strings.Join(cache," ")
	//
	//params := make([]interface{},0,m.builderParams.Len())
	//for i := m.builderParams.Front(); i != nil ; i = i.Next() {
	//	params = append(params,i.Value)
	//}
	//
	//// 查询数据
	//rows,err := DB.Query(queryString,params...)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//// 查询的列
	//columns,err := rows.Columns()
	//fields := make([]interface{},0,len(columns))
	//fieldsIndex := reflectIndex(container)
}

func (m *Msql) Query(container interface{}) {
	cache := make([]string,0,m.builderCache.Len())
	for i := m.builderCache.Front(); i != nil; i = i.Next() {
		cache = append(cache,i.Value.(string))
	}

	queryString := strings.Join(cache," ")
	fmt.Println(queryString)

	params := make([]interface{},0,m.builderParams.Len())
	for i := m.builderParams.Front(); i != nil ; i = i.Next() {
		params = append(params,i.Value)
	}

	// 反射出结构


	rows,err := DB.Query(queryString,params...)
	if err != nil {
		log.Println(err)
		return
	}

	//res := make(map[string]interface{})
	//var id interface{}
	//fmt.Printf("i am len(rows.Columns) %d")

	columns,err := rows.Columns()
	fields := make([]interface{},0,len(columns))
	fieldsIndex := reflectIndex(container)

	// 反射数据容器
	ref := reflect.ValueOf(container).Elem()
	// 遍历所有列
	for _,v := range columns {
		index,ok := fieldsIndex[v]
		if ok {
			fmt.Println("i am ok")
			fields = append(fields,ref.Field(index).Addr().Interface())
		}else{
			var tmp interface{}
			fields = append(fields,&tmp)
		}
	}

	//var fields []interface{}
	//fields = append(fields,&id)
	//fmt.Println(rows.Columns())
	//fields := make([]interface{})

	for rows.Next() {
		err := rows.Scan(fields...)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	fmt.Println(rows)
}

// 遍历columns，然后找出对应地址

//func reflect23(columns []string,container interface{}) (fields []interface{}) {
//
//	arr := make([]interface{},len(columns))
//	for k,v := range columns {
//
//	}
//
//	obj := reflect.ValueOf(container)
//	ref := obj.Elem()
//	typ := ref.Type()
//	for i:=0;i<ref.NumField();i++{
//		field := ref.Field(i)
//		// typ.field(i).Name 字段名称
//		// field.Type 字段类型
//		// field.Interface() 字段的值
//		fields = append(fields,field.Addr().Interface())
//	}
//	return
//}

func reflectIndex(v interface{}) map[string]int {
	index := make(map[string]int)
	obj := reflect.ValueOf(v)
	ref := obj.Elem()
	typ := ref.Type()
	for i:=0;i<ref.NumField();i++{
		tag := typ.Field(i).Tag.Get("mysql")
		index[tag] = i
	}
	return index
}

/*


select 支持数组,字符串
msql.select().from(model).where().query()


*/
