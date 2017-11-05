package mariadb

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Mariadb struct {
	dbconn *sql.DB
}

type DB struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func getdbconf() DB {

	//默认读取南京仓库数据库
	db := DB{Host: "10.42.5.245", Port: "3306", Database: "portus_development", Username: "root", Password: "portus"}

	h := os.Getenv("DB_HOST")
	if h != "" {
		db.Host = h
	}

	po := os.Getenv("DB_PORT")
	if po != "" {
		db.Port = po
	}

	d := os.Getenv("DB_DATABASE")
	if d != "" {
		db.Database = d
	}

	u := os.Getenv("DB_USERNAME")
	if u != "" {
		db.Username = u
	}
	pa := os.Getenv("DB_PASSWORD")
	if pa != "" {
		db.Password = pa
	}

	return db
}

//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]

func NewMariadb(database string) (m *Mariadb, err error) {
	m = new(Mariadb)
	db := getdbconf()
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db.Username, db.Password, db.Host, db.Port, db.Database)
	m.dbconn, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println("Open database error: %s\n", err)
	}

	for {
		err = m.dbconn.Ping()
		if err != nil {
			time.Sleep(10 * time.Second)
			fmt.Println("Cannot connect mariadb now ! Waiting for mysql init process done , try again later!!")
			continue
		}
		break
	}

	return m, err
}
