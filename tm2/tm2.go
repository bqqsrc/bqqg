package tm2
import (
	"time"
)

const (
	YEAR = 1    //字符串显示年
	MONTH = 2   //字符串显示月
	DAY = 4   //字符串显示日
	HOUR = 8   //字符串显示时
	MINUTE = 16   //字符串显示分
	SECOND = 32   //字符串显示秒
)

const (
	UTC = 0   //标准时区
	CST = 8   //东八区（北京时区）
)

// 将interface{}转换为时间字符串，格式为2006-01-02 15:04:05
// 
// mode表示字符串要输出的内容，需要什么就给mode添加什么位
//
// 例如，字符串需要年月日，mode就取值tm2.YEAR || tm2.MONTH || tm2.DAY，
// 字符串需要15:04:05，mode就取值tm2.HOUR || tm2.MINUTE || tm2.SECOND
//
// 支持将实际类型为int、int64、float32、float64、byte、[]byte、string、time.Time的interfac{}转换为日期字符串
// 
// zone指定时区，当为-1时，则表示当前时区
//
// int、int64、float32、float64、byte数字类型会先以数值转为时间戳(int64)，再转换为time.Time，再转换为字符串类型
// 
// string会先使用strconv相关接口转换为时间戳（int64)，再转换为time.Time，之后转换为字符串类型，
// 如果string无法转换为int64时间戳，则会按照2006-01-02 15:04:05的格式来获取日期
// 如果字符串不是2006-01-02 15:04:05的格式，将会报错
//
// []byte先转换为字符串，再以字符串的方式获取日期字符串
func I2DateTimeStr(data interface{}, mode, zone int) (string, error) {
	return "", nil
}

// 将interface{}转换为时间戳（int64）
//
// 支持将实际类型为int、int64、float32、float64、byte、[]byte、string、time.Time的interfac{}转换为时间戳
// 
// zone指定时区，当为-1时表示当前时区
//
// int、int64、float32、float64、byte数字类型会直接转换为时间戳（int64）返回
// 
// string会先使用strconv相关接口转换为（int64)作为时间戳返回，
// 如果string无法转换为int64时间戳，则会按照2006-01-02 15:04:05的格式转换为time.Time再获取时间戳
// 如果字符串不是2006-01-02 15:04:05的格式，将会报错
//
// []byte先转换为字符串，再以字符串的方式获取日期字符串
func I2TimeStamp(data interface{}, zone int) (int64, error) {
	return 0, nil
}

// 将interface{}转换为time.Time
//
// 支持将实际类型为int、int64、float32、float64、byte、[]byte、string、time.Time的interfac{}转换为time.Time
//
// zone指定时区，当为-1时表示当前时区
//
// int、int64、float32、float64、byte数字类型会先以数值作为时间戳(int64)，再转换为time.Time
// 
// string会先使用strconv相关接口转换为时间戳（int64)，再转换为time.Time
// 如果string无法转换为int64时间戳，则会按照2006-01-02 15:04:05的格式转换为time.Time
// 如果字符串不是2006-01-02 15:04:05的格式，将会报错
//
// []byte先转换为字符串，再以字符串的方式获取日期字符串
func I2Time(data interface{}, zone int) (time.Time, error) {
	return time.Now(), nil
}

// 将time.Time转为日期字符串
// 
// mode表示字符串要输出的内容，需要什么就给mode添加什么位
//
// 例如，字符串需要年月日，mode就取值tm2.YEAR && tm2.MONTH && tm2.DAY，
// 字符串需要15:04:05，mode就取值tm2.HOUR && tm2.MINUTE && tm2.SECOND
//
// zone指定时区，当为-1时表示当前时区
func Time2DateTimeStr(tm time.Time, mode, zone int) (string) {
	return ""
}

// 将time.Time转为int64时间戳
func Time2TimeStamp(tm time.Time) (int64) {
	return 0
}

// 将日期时间字符串转为int64时间戳
//
//日期格式必须为2006-01-02 15:04:05
//
// zone指定时区，当为-1时表示当前时区
func DateTimeStr2TimeStamp(dateStr string, zone int) (int64, error) {
	return 0, nil
}

// 将日期时间字符串转为time.Time
//
// 日期格式必须为2006-01-02 15:04:05
//
// zone指定时区，当为-1时表示当前时区
func DateTimeStr2Time(dataStr string, zone int) (time.Time, error) {
	return time.Now(), nil
}

// 将时间戳转为日期字符串
// 
// mode表示字符串要输出的内容，需要什么就给mode添加什么位
//
// 例如，字符串需要年月日，mode就取值tm2.YEAR && tm2.MONTH && tm2.DAY，
// 字符串需要15:04:05，mode就取值tm2.HOUR && tm2.MINUTE && tm2.SECOND
//
// zone指定时区，当为-1时表示当前时区
func TimeStamp2DataTimeStr(timeStamp int64, zone int) (string) {
	return ""
}

// 将时间戳转为time.Time
func TimeStamp2Time(timeStamp int64) (time.Time) {
	return time.Now()
}

// 获取当前时间字符串
// 
// mode表示字符串要输出的内容，需要什么就给mode添加什么位
//
// 例如，字符串需要年月日，mode就取值tm2.YEAR && tm2.MONTH && tm2.DAY，
// 字符串需要15:04:05，mode就取值tm2.HOUR && tm2.MINUTE && tm2.SECOND
//
// zone指定时区，当为-1时表示当前时区
func GetNowDateTimeStr(mode, zone int) string {
	return ""
}

// 获取当前时间戳（int64)，秒
func GetNowTimeStamp() int64 {
	return time.Now().Unix()
}

// 获取当前时间戳（int64），纳秒
func GetNowNanosecond() int64 {
	return time.Now().UnixNano()
}

func getLocalZone() {

}

func getLayout(mode int) string {
	return ""
}



// func I2DateTimeStr(data interface{}) (string, error) {
// 	var timeStamp int64
// 	var dateStr string
// 	strType := true
// 	switch data.(type) {
// 	case int64:
// 		timeStamp, _ = data.(int64)
// 		strType = false
// 		break
// 	case int:
// 		value, _ := data.(int)
// 		timeStamp = int64(value)
// 		strType = false
// 		break
// 	case string:
// 		dateStr = data.(string)
// 		break
// 	case []byte:
// 		value := data.([]byte)
// 		dateStr = string(value)
// 		break
// 	default:
// 		loger.Errorf("Error, I2DateTimeStr error, a %s value can't convert to TimeStamp", reflect.TypeOf(data))
// 		return "", false
// 	}
// 	var tm time.Time
// 	if strType {
// 		var err error
// 		tm, err = time.Parse("2006-01-02 15:04:05", dateStr)
// 		if err != nil {
// 			loger.Errorf("MustDateTimeStr error, time.Parse(\"2006-01-02 15:04:05\", %s) error, err: %s", dateStr, err)
// 			ok := false
// 			timeStamp, ok = I2Int64(dateStr)
// 			if ok {
// 				strType = false
// 			} else {
// 				return "", false
// 			}
// 		}
// 	} 
// 	if !strType {
// 		tm = time.Unix(timeStamp, 0)
// 	}
// 	return tm.Format("2006-01-02 15:04:05"), true
// }

