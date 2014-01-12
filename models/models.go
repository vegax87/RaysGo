package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lunny/xorm"
	"time"
)

var (
	Engine *xorm.Engine
)
// models declaration
type User struct {
	Id         int64
	Rid        int64     `xorm:"index"`
	Name       string    `xorm:"not null unique"`
	Email      string    `xorm:"not null unique"`
	Password   string    `xorm:"not null"`
	CreateTime time.Time `xorm:"index"`
	LoginTime  time.Time
	Picture    string
	Signature  string
	Status     int `xorm:"not null default 0"`
}

type Role struct {
	Id          int64
	Name        string `xorm:"not null unique"`
	Description string `xorm:"text"`
}

type NodeType struct {
	Id          int64
	Name        string `xorm:"not null unique"`
	Description string `xorm:"text"`
}

type Node struct {
	Id          int64
	Uid         int64  `xorm:"index"`
	Tid         int64  `xorm:"index"`
	Title       string `xorm:"not null"`
	Content     string `xorm:"text"`
	ContentType string
	Summary     string    `xorm:"text"`
	CreateTime  time.Time `xorm:"index"`
	UpdateTime  time.Time
}

type Category struct {
	Id          int64
	Uid         int64  `xorm:"index"`
	Name        string `xorm:"not null"`
	Description string `xorm:"text"`
}

type CategoryTerm struct {
	Id     int64
	Uid    int64  `xorm:"index"`
	Cid    int64  `xorm:"index"`
	Pid    int64  `xorm:"index default 0"`
	Name   string `xorm:"not null"`
	Weight int64  `xorm:"default 0"`
}

type Comment struct {
	Id           int64
	Uid          int64 `xorm:"index default 0"`
	Pid          int64 `xorm:"index default 0"`
	Nid          int64 `xorm:"index"`
	Title        string
	Content      string `xorm:"text"`
	ContentType  string
	CreateTime   time.Time `xorm:"index"`
	UpdateTime   time.Time `xorm:"index"`
	UserHost     string
	UserName     string
	UserEmail    string
	UserHomePage string
}

type File struct {
	Id        int64
	Uid       int64  `xorm:"index"`
	Name      string `xorm:"not null"`
	Uri       string `xorm:"not null"`
	Size      int64  `xorm:"not null default 0"`
	Mimetype  string
	Status    int
	Timestamp time.Time
}

type UriAlias struct {
	Id     int64
	Source string `xorm:"not null"`
	Uri    string `xorm:"not null"`
}

type Variable struct {
	Id    int64
	Value string `xorm:"text"`
}

/* end of model declaration */

// database settings
var (
	dbtype     string
	dbname     string
	dbhost     string
	dbport     string
	dbuser     string
	dbpassword string
	dbcharset  string
)

func init() {
	loadDbConfig()

	_, err := InitEngine()
	if err != nil {
		fmt.Println(err)
	}
}

func loadDbConfig() {
	dbtype = beego.AppConfig.String("dbtype")
	dbhost = beego.AppConfig.String("dbhost")
	dbport = beego.AppConfig.String("dbport")
	dbuser = beego.AppConfig.String("dbuser")
	dbpassword = beego.AppConfig.String("dbpassword")
	dbname = beego.AppConfig.String("dbname")
	dbcharset = beego.AppConfig.String("dbcharset")
	if dbport == "" {
		dbport = "3306"
	}
}

func ConnectDb() (*xorm.Engine, error) {
	fmt.Println("database type: " + dbtype)
	switch {
	case dbtype == "sqlite":
		return xorm.NewEngine("sqlite3", dbname)

	case dbtype == "mysql":
		return xorm.NewEngine("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v",
			dbuser, dbpassword, dbhost, dbport, dbname, dbcharset))

	case dbtype == "pgsql":
		return xorm.NewEngine("postgres", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v",
			dbuser, dbpassword, dbhost, dbport, dbname, dbcharset))
	}
	return nil, errors.New("No database found!")
}

func InitEngine() (*xorm.Engine, error) {
	var err error
	Engine, err = ConnectDb()

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	Engine.SetDefaultCacher(cacher)

	return Engine, err
}

// Create database instance according to the models declaration
func CreateDb() {
	if err := Engine.Sync(
		new(User),
		new(Role),
		new(Node),
		new(NodeType),
		new(Comment),
		new(Category),
		new(CategoryTerm),
		new(Variable),
		new(File),
		new(UriAlias)); err != nil {
		fmt.Println("Database sync failed: ", err)
	} else {
		fmt.Println("Database sync successfully.")
	}
}
