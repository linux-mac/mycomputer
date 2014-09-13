package models

import (
       "github.com/astaxie/beego"
       "github.com/astaxie/beego/orm"
       _ "github.com/go-sql-driver/mysql"
)

type User struct {
     Name        string `orm:"pk"` 
     Email         string
     Description       string
//    Posts []*Post `orm:"reverse(many)"` 
}

type Item struct {
     Number                    int64 `orm:"pk"` 
     Image                 string
     Description	   string
     Username		   string
//     User *User `orm:"rel(fk)"`
}

type Comment struct {
     Id	     int64 `orm:"pk;auto"`
     Content string
}

func (*User) TableEngine() string {
     return engine()
}

func (*Item) TableEngine() string {
     return engine()
}

func (*Comment) TableEngine() string {
     return engine()
}

func engine() string {
     return "INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci"
}

func init() {
//     orm.RegisterModelWithPrefix("mycomputer_", new(User), new(Item))
     orm.RegisterModel(new(User), new(Item))
     orm.RegisterDriver("mysql", orm.DR_MySQL)

db_user := beego.AppConfig.String("db_user")
db_password := beego.AppConfig.String("db_password")
db_host := beego.AppConfig.String("db_host")
db_port := beego.AppConfig.String("db_port")
db_database := beego.AppConfig.String("db_database")
     
    orm.RegisterDataBase("default", "mysql", db_user + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + db_database + "?charset=utf8")

//    orm.RegisterDataBase("default", "mysql", "root:561801src@tcp(104.131.61.252:3306)/mycomputer", 30)
}
