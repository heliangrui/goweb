package gis

import (
	"fmt"
	"go-web/model/vo"
	"math"
	"strconv"
)

// Point is a utility class for GPS calcuLngions.
// 小写方法是私有方法，大写方法是公有方法 可根据需要调整

const (
	WGS84 = 0 //大地坐标体系
	GCJ02 = 1 //火星坐标体系
	BD09  = 2 //百度坐标体系
)

const (
	pi   = 3.1415926535897932384626             // 圆周率
	x_pi = 3.14159265358979324 * 3000.0 / 180.0 // 圆周率对应的经纬度偏移
	a    = 6378245.0                            // 长半轴
	ee   = 0.00669342162296594323               // 扁率
)

func outOfChina(lon, lat float64) bool {
	if lon < 72.004 || lon > 137.8347 {
		return true
	} else if lat < 0.8293 || lat > 55.8271 {
		return true
	}
	return false
}
func transformLat(lon, lat float64) float64 {

	ret := -100.0 + 2.0*lon + 3.0*lat + 0.2*lat*lat + 0.1*lon*lat + 0.2*math.Sqrt(math.Abs(lon))
	ret += (20.0*math.Sin(6.0*lon*pi) + 20.0*math.Sin(2.0*lon*pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(lat*pi) + 40.0*math.Sin(lat/3.0*pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(lat/12.0*pi) + 320*math.Sin(lat*pi/30.0)) * 2.0 / 3.0
	return ret
}

func transformLng(lon, lat float64) float64 {
	ret := 300.0 + lon + 2.0*lat + 0.1*lon*lon + 0.1*lon*lat + 0.1*math.Sqrt(math.Abs(lon))
	ret += (20.0*math.Sin(6.0*lon*pi) + 20.0*math.Sin(2.0*lon*pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(lon*pi) + 40.0*math.Sin(lon/3.0*pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(lon/12.0*pi) + 300.0*math.Sin(lon/30.0*pi)) * 2.0 / 3.0
	return ret
}

func AutoExchange(point vo.Point, tag int8) vo.Point {

	result := vo.Point{}

	switch point.Type {
	case BD09:
		switch tag {
		case WGS84:
			result = Bd09_To_WGS84(point.Longitude, point.Latitude)
		case GCJ02:
			result = Bd09_To_Gcj02(point.Longitude, point.Latitude)
		case BD09:
			result = point
		}
	case GCJ02:
		switch tag {
		case WGS84:
			result = GCJ02_To_WGS84(point.Longitude, point.Latitude)
		case GCJ02:
			result = point
		case BD09:
			result = gcj02_To_Bd09(point.Longitude, point.Latitude)
		}
	case WGS84:
		switch tag {
		case WGS84:
			result = point
		case GCJ02:
			result = WGS84_To_Gcj02(point.Longitude, point.Latitude)
		case BD09:
			result = WGS84_To_Bd09(point.Longitude, point.Latitude)
		}

	}
	return result
}

// WGS84_To_Gcj02 84 to 火星坐标系 (GCJ-02) World Geodetic System ==> Mars Geodetic System
func WGS84_To_Gcj02(lng, lat float64) vo.Point {
	point := vo.Point{}
	point.Type = GCJ02

	if outOfChina(lng, lat) {
		point.Longitude = lng
		point.Latitude = lat
		return point
	}

	dLng := transformLng(lng-105.0, lat-35.0)
	dLat := transformLat(lng-105.0, lat-35.0)
	radlat := lat / 180.0 * pi
	magic := math.Sin(radlat)
	magic = 1 - ee*magic*magic
	SqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * SqrtMagic) * pi)
	dLng = (dLng * 180.0) / (a / SqrtMagic * math.Cos(radlat) * pi)
	point.Longitude = lng + dLng
	point.Latitude = lat + dLat
	return point
}

// GCJ02_To_WGS84
// 火星坐标系 (GCJ-02) to WGS84
func GCJ02_To_WGS84(lng, lat float64) vo.Point {
	point := vo.Point{}
	point.Type = WGS84
	fmt.Println("AAAAA")
	fmt.Println(lng)
	fmt.Println(lat)
	if outOfChina(lng, lat) {
		point.Longitude = lng
		point.Latitude = lat
		return point
	}
	fmt.Println("BBBBB")
	dLat := transformLat(lng-105.0, lat-35.0)
	dLng := transformLng(lng-105.0, lat-35.0)
	radLat := lat / 180.0 * pi
	magic := math.Sin(radLat)
	magic = 1 - ee*magic*magic
	SqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * SqrtMagic) * pi)
	dLng = (dLng * 180.0) / (a / SqrtMagic * math.Cos(radLat) * pi)
	mgLat := lat + dLat
	mgLng := lng + dLng

	point.Longitude = lng*2 - mgLng
	point.Latitude = lat*2 - mgLat

	fmt.Println(point.Longitude)
	fmt.Println(point.Latitude)

	return point
}

/**
 * 火星坐标系 (GCJ-02) 与百度坐标系 (BD-09) 的转换算法 将 GCJ-02 坐标转换成 BD-09 坐标
 */
func gcj02_To_Bd09(lng, lat float64) vo.Point {
	point := vo.Point{}
	point.Type = BD09

	x := lng
	y := lat
	z := math.Sqrt(x*x+y*y) + 0.00002*math.Sin(y*x_pi)
	theta := math.Atan2(y, x) + 0.000003*math.Cos(x*x_pi)
	point.Longitude = z*math.Cos(theta) + 0.0065
	point.Latitude = z*math.Sin(theta) + 0.006

	return point
}

/**
 * * 火星坐标系 (GCJ-02) 与百度坐标系 (BD-09) 的转换算法 * * 将 BD-09 坐标转换成GCJ-02 坐标 * * @param
 * bd_Lng * @param bd_Lat * @return
 */
func Bd09_To_Gcj02(lng, lat float64) vo.Point {
	point := vo.Point{}
	point.Type = GCJ02
	x := lng - 0.0065
	y := lat - 0.006
	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*x_pi)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*x_pi)
	point.Longitude = z * math.Cos(theta)
	point.Latitude = z * math.Sin(theta)
	return point
}

/**将WGS84转为bd09
 */
func WGS84_To_Bd09(lng, lat float64) vo.Point {
	gcj02 := WGS84_To_Gcj02(lng, lat)
	bd09 := gcj02_To_Bd09(gcj02.Longitude, gcj02.Latitude)
	return bd09
}

func Bd09_To_WGS84(lng, lat float64) vo.Point {

	gcj02 := Bd09_To_Gcj02(lng, lat)
	wgs84 := GCJ02_To_WGS84(gcj02.Longitude, gcj02.Latitude)

	return wgs84
}

/**保留小数点后六位
 */
func retain6(num float64) float64 {
	value, _ := strconv.ParseFloat(strconv.FormatFloat(num, 'f', 6, 64), 64)
	return value
}
