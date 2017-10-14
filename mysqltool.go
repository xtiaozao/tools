package main;
 
import (
    "database/sql"
   _ "github.com/go-sql-driver/mysql"
    "fmt"
    "time"  
)
 
func main() {
    //打开数据库
    //DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
    ma := make(map[string]int)
    db, err := sql.Open("mysql", "cc:CC_DBPASS@tcp(127.0.0.1:3306)/cc");
    if err != nil {
        fmt.Println("123") 
        fmt.Println(err);
    }else{
        fmt.Println("connect success")
    }
       
   flag := false 
 for{

    time.Sleep(time.Millisecond*5000)
    showTables,err := db.Query("show tables")
    if err != nil{
         fmt.Println("show tables wrong") 
         return  
    }    
    defer showTables.Close()  
  
   

     for showTables.Next(){
        var name  string
        var num   int
        err = showTables.Scan(&name)
        if err!=nil{
          fmt.Println("get string error")
          return  
        }
       count,err :=  db.Query("select count(*) from "+name+";")
       count.Next()  
       if err!=nil{
            fmt.Println("count error")
       }        
       err = count.Scan(&num)
       count.Close()   
       if err != nil{
           fmt.Println("get count error")
           return         
       }
       if flag==false{
          ma[name] = num      
       }else{
          if ma[name]!=num{
             fmt.Println(name,"'s change =",num-ma[name])
             ma[name] = num  
         }     
      }    
    }
      flag=true      
  }

    db.Close()
}
