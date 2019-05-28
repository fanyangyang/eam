package models

import "github.com/astaxie/beego/orm"

var Ormer orm.Ormer

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Department), new(Position), new(Item), new(IType), new(AskList), new(TODO))
}

func NewOrm() {
	Ormer = orm.NewOrm()
	Ormer.Using("default") // 默认使用 default，你可以指定为其他数据库
}

type RespSuccess struct {
	Success bool   `json:"success"`
	Desc    string `json:"desc"`
}

type Account struct {
	Id       int
	UserName string
	Password string
	Number   string
	Mail     string
}

type Department struct {
	Id         int
	ChargeId   int
	ChargeName string
	Count      int
	Desc       string
}

type Position struct {
	Id   int
	Desc string
}

type User struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Sex          int     `json:"sex"`
	Age          int     `json:"age"`
	DepartmentId int     `json:"department_id"`
	PositionId   int     `json:"position_id"`
	BoardDate    string  `json:"board_date"`
	Item         []*Item `json:"item" orm:"reverse(many)"`
	Status       int     `json:"status"` // 在职、离职
}

type UserRet struct {
	Count    int64  `json:"count"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
	Users    []User `json:"users"`
}

type Item struct {
	Id              int
	Name            string
	Number          string  //手机号
	Balance         float64 //话费余额
	User            *User `orm:"rel(fk)"`
	SerialCode      string
	ShoppingCode    string
	SourcePlateForm string
	ShopDate        string
	ITypeId         int    //资产类别
	Desc            string // TODO 确定orm是否支持自动添加时间
}

type IType struct {
	Id       int
	ParentId int
	Name     string
	Desc     string
}

type AskList struct {
	Id      int
	AType   int // 申领，申还，申换
	WhoId   int
	Who     string
	forWhat string
	Date    string
	Status  int //TODO 申请状态
	Desc    string
}

type TODO struct {
	Id     int
	Title  string
	Body   string
	Status int //新建、进行中、已完成
	Desc   string
}
