package result

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestMapData_Add(t *testing.T) {
	t.Log(MapSuccess.Add("id", 2).Add("test", "3"))
	t.Log(MapSuccess.Add("id", 7).Add("test", "3eyw"))

	pager := GetInfoPager{
		GetInfo: &GetInfo{
			MapData: MapSuccess,
			Data:    "test",
		},
		Pager: Pager{
			ClientPage: 2,
			EveryPage:  3,
			TotalNum:   5,
		},
	}
	// pager
	t.Log(pager.Add("id", 1).Add("test", 2))
}

func TestMapData_AddStruct(t *testing.T) {
	type User struct {
		ID   uint64
		Name string `json:"name"`
	}
	var user = User{
		ID:   0,
		Name: "test",
	}
	t.Log(MapSuccess.Add("id", 2).AddStruct(user))
}

func httpServerDemo(w http.ResponseWriter, r *http.Request) {
	pager := GetInfoPager{
		GetInfo: &GetInfo{
			MapData: MapSuccess,
			Data:    "test",
		},
		Pager: Pager{
			ClientPage: 2,
			EveryPage:  3,
			TotalNum:   5,
		},
	}
	// pager
	log.Println(pager.Add("id", 1).Add("test", 2))
	fmt.Fprintf(w, pager.Add("id", 1).Add("test", 2).String())
}

// test http request
func TestRequest(t *testing.T) {
	//http.HandleFunc("/", httpServerDemo)
	//err := http.ListenAndServe(":9090", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
}
