package reflect

import (
	"log"
	"testing"
)

// order
type Order struct {
	ID         int64 `json:"id"`
	UserID     int64 `json:"user_id"`     // user id
	ServiceID  int64 `json:"service_id"`  // service table id
	CreateTime int64 `json:"create_time"` // createtime
}

func TestGetDataID(t *testing.T){
	or := Order{}//new(Order)
	or.ID = 23
	id,_ := GetDataByFieldName(or, "ID")
	log.Println("id value is ",id)
}
