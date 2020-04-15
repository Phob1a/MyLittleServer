package Server

import (
	"database/sql"
	"fmt"
)

func ConnectToDB() *sql.DB {
   user:="root"
   psw:="123456"
   host:="localhost"
   port:=3306
   dbname:="MylittleDB"
   charset:="utf8"
   config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", user, psw, host, port, dbname, charset)
   db, err:=sql.Open("mysql",config)
   if err!=nil{
   	fmt.Println("Open database error",err)
   }
   return db
}

func Insert(db *sql.DB,id int,username string,password string,nickname string) error{
	sen:="INSERT INTO user(username,password,nickname)VALUES(?,?,?)"
	stmt, err:=db.Prepare(sen)
	if err!=nil{
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password, nickname)
	if err!=nil{
		return err
	}
}

func SelectByUsername(db *sql.DB,username string) (sql.Result,error){
    sen:="SELECT * FROM user WHERE username=?"
	stmt, err:=db.Prepare(sen)
	if err!=nil{
		return nil,err
	}
	defer stmt.Close()
	result,err:=stmt.Exec(username)
	if err!=nil{
		return nil,err
	}
	return result,nil
}

func UpdateNickname(db *sql.DB, id int, nickname string) error{
	sen :="Update user SET nickname=? WHERE id=?"
	stmt,err:=db.Prepare(sen)
	if err!=nil{
		return err
	}
	defer stmt.Close()
	result, err:=stmt.Exec(nickname,id)
	if err!=nil{
		return err
	}
	rowsAffected, err:=result.RowsAffected()
	if err!=nil{
		return err
	}
	fmt.Printf("%d records has been updated.\n",rowsAffected)
	return nil
}



