package gomeme

//
//import (
//	"time"
//	"encoding/json"
//	"github.com/davecgh/go-spew/spew"
//)
//
//
//
//
//type AutoGenerated struct {
//	ComputedData struct {
//		Num0 []struct {
//			Date    time.Time `json:"date"`
//			Zone    int       `json:"zone"`
//			Focus   int       `json:"focus"`
//			Calm    int       `json:"calm"`
//			Posture int       `json:"posture"`
//			BkiSum  float64   `json:"bki_sum"`
//			BkiN    float64   `json:"bki_n"`
//			ScLbs   float64   `json:"sc_lbs"`
//		}
//	} `json:"computed_data"`
//	Cursor string `json:"cursor"`
//}
//
//func main(){
//	result := AutoGenerated{}
//	json.Unmarshal([]byte(jsonStr,&result),
//	spew.Dump(result)
//}
//
//var jsonStr = `
//{
//  "computed_data": {
//    "0": [
//      {
//        "date": "2016-06-01T00:59:45+09:00",
//        "zone": 50,
//        "focus": 50,
//        "calm": 50,
//        "posture": 80,
//        "bki_sum" : 8.76,
//        "bki_n" : 6.15,
//        "sc_lbs" : 2.37
//      },
//      {
//        "date": "2016-06-01T00:59:30+09:00",
//        "zone": 60,
//        "focus": 50,
//        "calm": 50,
//        "posture": 90,
//        "bki_sum" : 8.76,
//        "bki_n" : 6.15,
//        "sc_lbs" : 2.37
//      }
//    ],
//"1": [
//      {
//        "date": "2016-06-01T00:59:45+09:00",
//        "zone": 50,
//        "focus": 50,
//        "calm": 50,
//        "posture": 80,
//        "bki_sum" : 8.76,
//        "bki_n" : 6.15,
//        "sc_lbs" : 2.37
//      },
//      {
//        "date": "2016-06-01T00:59:30+09:00",
//        "zone": 60,
//        "focus": 50,
//        "calm": 50,
//        "posture": 90,
//        "bki_sum" : 8.76,
//        "bki_n" : 6.15,
//        "sc_lbs" : 2.37
//      }
//    ]
//  },
//  "cursor": "eyJhIjoxLCJiIjoyLCJjIjozfQ=="
//}`