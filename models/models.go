package models

import (
	"github.com/astaxie/beego/orm"
)

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
	Id           int
	Name         string
	Sex          int
	Age          int
	DepartmentId int
	PositionId   int
	BoardDate    string
	Item         []*Item `orm:"reverse(many)"`
	Status       int     // 在职、离职
}

type Item struct {
	Id              int
	Name            string
	Number          string  //手机号
	Balance         float64 //话费余额
	User            *User   `orm:"rel(fk)"`
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

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Department), new(Position), new(Item), new(IType), new(AskList), new(TODO))
}
