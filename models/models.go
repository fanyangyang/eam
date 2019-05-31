package models

import "github.com/astaxie/beego/orm"

var Ormer orm.Ormer

const (
	USER_STATUS_ON  = 1
	USER_STATUS_OFF = 2

	TODO_STATUS_NEW     = 1
	TODO_STATUS_ING     = 2
	TODO_STATUS_DONE    = 3
	TODO_STATUS_TIMEOUT = 4

	TODO_LEVEL_NORMAL    = 1
	TODO_LEVEL_IMPORTANT = 2

	ITYPE_STATUS_ON  = 1
	ITYPE_STATUS_OFF = 2

	ITEM_STATUS_STORAGE = 1
	ITEM_STATUS_USE     = 2
	ITEM_STATUS_BACK    = 3

	DEPARTMENT_STATUS_ON  = 1
	DEPARTMENT_STATUS_OFF = 2

	POSITION_STATUS_ON  = 1
	POSITION_STATUS_OFF = 2

	DEFAULT_PAGE_SIZE = 10
	DEFAULT_PAGE_NUM  = 1

	DEFAULT_ID = 0
)

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
	Message string `json:"message"`
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
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Charge *User   `orm:"rel(fk)" json:"charge"`
	Desc   string  `json:"desc"`
	Status int     `json:"status"` // 1: 在用 2：舍弃
	Users  []*User `orm:"reverse(many)" json:"users"`
}

type DepartmentRet struct {
	TotalNum    int64        `json:"total_num"`
	PageNum     int          `json:"page_num"`
	PageSize    int          `json:"page_size"`
	Departments []Department `json:"departments"`
}

type Position struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Desc   string  `json:"desc"`
	Users  []*User `orm:"reverse(many)" json:"users"`
	Status int     `json:"status"` //1：在用 2：舍弃
}

type PositionRet struct {
	TotalNum  int64      `json:"total_num"`
	PageNum   int        `json:"page_num"`
	PageSize  int        `json:"page_size"`
	Positions []Position `json:"positions"`
}

type User struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"` // TODO mail,personalPhoneNum,address
	Sex        int         `json:"sex"`
	Age        int         `json:"age"`
	Department *Department `orm:"rel(fk)" json:"department"`
	Position   *Position   `orm:"rel(fk)" json:"position"` // 添加TODO，说明多少人的职位未分配，部门未分配
	BoardDate  string      `json:"board_date"`                //入职时间
	Item       []*Item     `json:"item" orm:"reverse(many)"`
	Status     int         `json:"status"` // 1:在职、2:离职
}

type UserRet struct {
	TotalNum int64  `json:"total_num"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
	Users    []User `json:"users"`
}

type Item struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"` //手机号
	//Balance         float64 `json:"balance"` //话费余额
	User            *User  `orm:"rel(fk)" json:"user"`
	SerialCode      string `json:"serial_code"`
	ShoppingCode    string `json:"shopping_code"`
	SourcePlateForm string `json:"source_plate_form"`
	ShopDate        string `json:"shop_date"`
	IType           *IType `orm:"rel(fk)" json:"i_type"` //资产类别
	Desc            string `json:"desc"`                 // TODO 确定orm是否支持自动添加时间
	Status          int    `json:"status"`               //1：在库 2：已分配 3：回收
}

// TODO 话费充值记录

type ItemRet struct {
	TotalNum int64  `json:"total_num"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
	Items    []Item `json:"items"`
}

type IType struct {
	Id       int     `json:"id"`
	ParentId int     `json:"parent_id"`
	Name     string  `json:"name"`
	Desc     string  `json:"desc"`
	Status   int     `json:"status"` //1:on、2：off
	Items    []*Item `json:"items" orm:"reverse(many)"`
}

type ITypeRet struct {
	TotalNum int64   `json:"total_num"`
	PageNum  int     `json:"page_num"`
	PageSize int     `json:"page_size"`
	ITypes   []IType `json:"i_types"`
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
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Status   int    `json:"status"` //1:新建、2:进行中、3:已完成、4:已超时
	Desc     string `json:"desc"`
	DeadLine string `json:"dead_line"`
	Level    int    `json:"level"` //1:一般、 2:重要
}

type TODORet struct {
	TotalNum int64  `json:"total_num"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
	TODOs    []TODO `json:"todos"`
}
