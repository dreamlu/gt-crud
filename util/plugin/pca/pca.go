package pca

// 城市列表接口, 如有需求, 请打开注释
// 并将接口写入你的文档
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/dreamlu/gt"
//	"demo/util/result"
//	"github.com/gin-gonic/gin"
//	"github.com/mozillazg/go-pinyin"
//	"io/ioutil"
//	"net"
//	"net/http"
//	"reflect"
//	"sort"
//	"strings"
//)
//
//type Pca struct {
//	Code  string `json:"code"`
//	Name  string `json:"name"`
//	Alias string `json:"-"`
//}
//
//type Link struct {
//	Pca
//	Children []Link `json:"children"`
//}
//
//type CityItem struct {
//	Key   string `json:"key"`
//	Title string `json:"title"`
//	Items []Pca  `json:"items"`
//}
//
//var mem map[string]string
//
////城市列表
//func List(u *gin.Context) {
//	links, err := readPca()
//	if err != nil {
//		gt.Logger().Error(err)
//		u.JSON(http.StatusBadGateway, result.GetMapData(result.CodeError, "读取文件错误"))
//	}
//	initMap()
//	_ = replace(links)
//	u.JSON(http.StatusOK, result.GetSuccess(links))
//}
//
//func City(u *gin.Context) {
//	links, err := readPca()
//	if err != nil {
//		u.JSON(http.StatusBadGateway, result.GetMapData(result.CodeError, "读取文件错误"))
//		gt.Logger().Error(err)
//	}
//	initMap()
//	pcas := getCity(links)
//	SortBodyByName(pcas)
//	u.JSON(http.StatusOK, result.GetSuccess(Classify(pcas)))
//}
// ip定位城市
//func CityByIP(u *gin.Context) {
//	ip := u.Query("ip")
//	url := "https://apis.map.qq.com/ws/location/v1/ip?ip=" + ip + "&key=3DBBZ-GXRCF-EJKJ7-JHOGR-FYAZV-Y6F42"
//	response, _ := http.Get(url)
//	body, _ := ioutil.ReadAll(response.Body)
//	res := string(body)
//	fmt.Println(res)
//	m := make(map[string]interface{})
//	_ = json.Unmarshal(body, &m)
//	u.JSON(http.StatusOK, result.GetSuccess(m))
//}
//
//func readPca() ([]Link, error) {
//	pcaPath := conf.GetString("app.staticpath") + "pca/pca-code.json"
//	by, err := ioutil.ReadFile(pcaPath)
//	if err != nil {
//		return nil, err
//	}
//	var pca []Link
//	err = json.Unmarshal(by, &pca)
//	return pca, err
//}
//
//func replace(links []Link) error {
//	length := len(links)
//	for i := 0; i < length; i = i + 1 {
//		pca := links[i]
//		if pca.Children != nil && len(pca.Children) != 0 {
//			_ = replace(pca.Children)
//		}
//		if pca.Name == "市辖区" {
//			pca.Name = mem[pca.Code]
//		}
//		links[i] = pca
//	}
//	return nil
//}
//
//func getCity(links []Link) []Pca {
//	var pca Pca
//	var pcas []Pca
//	length := len(links)
//	for i := 0; i < length; i = i + 1 {
//		link := links[i]
//		pca = link.Pca
//		pca.Alias = strings.ToUpper(PinyinConvert(pca.Name))
//		if link.Children != nil && len(link.Children) != 0 {
//			pcas = append(pcas, getCity(link.Children)...)
//		}
//		if len(link.Code) == 4 {
//			if pca.Name == "市辖区" {
//				pca.Name = mem[pca.Code]
//				pca.Alias = strings.ToUpper(PinyinConvert(pca.Name))
//			}
//			pcas = append(pcas, pca)
//		}
//	}
//	return pcas
//}
//
//func PinyinConvert(str string) string {
//	str = pinyin.Paragraph(str)
//	return str
//}
//
//func Classify(p []Pca) []CityItem {
//	var re []CityItem
//	var s, e, length int
//	length = len(p)
//	for i := 'A'; i <= 'Z'; i = i + 1 {
//		k := string(i)
//		s = 0
//		e = 0
//		for j := 0; j < len(p); j = j + 1 {
//			if strings.HasPrefix(p[j].Alias, k) {
//				e = e + 1
//			} else {
//				break
//			}
//		}
//		r := CityItem{Key: k, Title: k, Items: p[s:e]}
//		re = append(re, r)
//		if e < length {
//			p = p[e:]
//		}
//	}
//	return re
//}
//
//func initMap() {
//	mem = map[string]string{}
//	mem["1101"] = "北京市"
//	mem["5001"] = "重庆市"
//	mem["3101"] = "上海市"
//	mem["1201"] = "天津市"
//}
//
//type CityWrapper struct {
//	Bodies []Pca
//	sort   func(p, q *Pca) bool
//}
//
//type SortBodyBy func(p, q *Pca) bool //定义一个函数类型
//
////数组长度Len()
//func (acw CityWrapper) Len() int {
//	return len(acw.Bodies)
//}
//
////元素交换
//func (acw CityWrapper) Swap(i, j int) {
//	acw.Bodies[i], acw.Bodies[j] = acw.Bodies[j], acw.Bodies[i]
//}
//
////比较函数，使用外部传入的by比较函数
//func (acw CityWrapper) Less(i, j int) bool {
//	return acw.sort(&acw.Bodies[i], &acw.Bodies[j])
//}
//
//func SortBodyByName(bodies []Pca) {
//	sort.Sort(CityWrapper{bodies, func(p, q *Pca) bool {
//		v := reflect.ValueOf(*p)
//		i := v.FieldByName("Alias")
//		v = reflect.ValueOf(*q)
//		j := v.FieldByName("Alias")
//		return i.String() < j.String()
//	}})
//}
//
//func RemoteIp(req *http.Request) string {
//	remoteAddr := req.RemoteAddr
//	if ip := req.Header.Get("Remote_addr"); ip != "" {
//		remoteAddr = ip
//	} else {
//		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
//	}
//
//	if remoteAddr == "::1" {
//		remoteAddr = "127.0.0.1"
//	}
//	return remoteAddr
//}
