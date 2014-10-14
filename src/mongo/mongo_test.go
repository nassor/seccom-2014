package mongo

import (
	"mongo/models"
	"testing"

	"gopkg.in/mgo.v2"
)

var collection *mgo.Collection

func init() {
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	database := mongoSession.DB("seccom")

	// clean database
	err = database.DropDatabase()
	if err != nil {
		panic(err)
	}

	collection = database.C("user")
}

func TestAddUser(t *testing.T) {
	nassor := models.User{"Nassor Paulino da Silva", 30}
	AddUser(collection, &nassor)

	user := new(models.User)
	err := collection.Find(nil).One(&user)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if user.Name != "Nassor Paulino da Silva" {
		t.Fatalf("Resultado esperado: %s / Resultado recebido:  %s", "Nassor Paulino da Silva", user.Name)
	}
}

func TestListUser(t *testing.T) {
	users, err := ListUsers(collection)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if users[0].Name != "Nassor Paulino da Silva" {
		t.Fatalf("Resultado esperado: %s / Resultado recebido:  %s", "Nassor Paulino da Silva", users[0].Name)
	}
}
