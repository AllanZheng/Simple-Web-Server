package util
import(
	/*
 "log"
 "sync"
 "time"
 	*/
 "gopkg.in/mgo.v2"
 "github/SimpleWeb/entity"
 "fmt"
 "gopkg.in/mgo.v2/bson"
)
/*
var (
	session   *mgo.Session
)

func Initialize(){
	db := NewDB()
	defer db.Session.close()

}

func NewDB() *mgo.Database {
	session := NewSession()
	return session.DB(CfgDBName())
}
*/
func NewDB() *mgo.Database{
	session,err := mgo.Dial("localhost:27017")
	db := session.DB("test")
	if err!=nil{
       fmt.Println("Database initization error!")
	}
	return db
}
func InitDB(){
	db := NewDB()
	defer db.Session.Close()
	var Indices = []mgo.Index{
		{Key:[]string{"Name"}},
	}
	for _,i := range Indices{
		if err := db.C("test").EnsureIndex(i);err!=nil{
			fmt.Println("db create index failed!")
		}
	}
}
func CreateForm(data *entity.Info)(err error){
	db := NewDB()
	defer db.Session.Close()
	if err = db.C("test").Insert(data); err!=nil{
		fmt.Println("Create Failed!")
	}
   return 
}

func UpdateForm(data *entity.Info,toup *entity.Info)(err error){
	db := NewDB()
	query := bson.M{
		"name": data.Name,
	}
	defer db.Session.Close()
	if err = db.C("test").Update(query,toup);err!=nil{
		fmt.Println("Update Failed! ")
	}
	return 
}

func FindForm(data *entity.Info)(err error,res entity.Info){
	db := NewDB()
	query := bson.M{
		"name": data.Name,
	}
	defer db.Session.Close()
	
	if err = db.C("test").Find(query).One(&res); err != nil {
		fmt.Println("error!",err)
	}
	return
}

func DeleteForm(data *entity.Info)(err error){
	db := NewDB()
	defer db.Session.Close()
	
	query := bson.M{
		"name": data.Name,
	}
	
	if err = db.C("test").Remove(query);err!=nil{
		fmt.Println("Delete Failed!")
	}
	return
}