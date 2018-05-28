package linuxNetwork

import (
	"go-networkInfo/cmd"
	"fmt"
	"strings"
)
//  bridge 结构体
//  Id string         ID 	字符串
//	Name string       名称  	字符串
//	Driver string     类型  	字符串
//	Scope string      范围  	字符串
type Bridge struct {
	Id string
	Name string
	STP string
	Interface string
}



func GetECI()[]string {
	ECI:=cmd.Lcmd("ifconfig",[]string{ })
	//arrayECI:=strings.Fields(ECI)

	//fmt.Println(ECI)

	temp:=strings.Split(ECI,"\n")
	var arrayECI []string
	flag:=0
	for i,value :=range temp{
		fmt.Println(i,value)

		if value==""{
			var t string
			for j:=flag;j<i;j++{
				t=t+temp[j]

			}
			arrayECI=append(arrayECI,t)
			flag=i+1
		}
	}
	n:=len(arrayECI)-1
	arrayECI=arrayECI[0:n]
	for i,value :=range arrayECI{
		fmt.Println(i,value)

	}



	tttttt:=strings.Split(arrayECI[0],"        ")
	for i,value :=range tttttt{
		fmt.Println(i,value)

	}



	return arrayECI
}
func ShowECI (){
	ECI:=GetECI()
	fmt.Println(ECI)
}


//获得主机网桥信息
func GetBridge()[]Bridge{

	//调用获取命令行返回信息
	bridge:=cmd.Lcmd("brctl",[]string{"show"})

	//对信息进行数组化s
	arrayBridge:=strings.Split(bridge,"\n")

	//去掉标题
	n:=len(arrayBridge)-1
	temp:=arrayBridge[1:n]

	//定义分割后的临时存储数组，定义返回值变量
	var t []string
	var bridgeStruct []Bridge

	//将数据赋给返回变量
	for _,value :=range temp{
		//n格空格分割，避免一处添加一个空
		t=strings.Fields(value)
		t=append(t,"")
		//赋值
		bt:=Bridge{t[0],t[1],t[2],t[3]}
		//切片追加
		bridgeStruct=append(bridgeStruct,bt)

	}
	return bridgeStruct
}
//显示主机网桥信息
func ShowBridge(){

	ttt:=GetBridge()
	fmt.Printf("%-8s%-20s%-10s%-15s%-15s\n",
		"Number","Id","Name","STP","Interface")
	for i,value :=range ttt{

		fmt.Printf("%-8d%-20s%-20s%-15s%-15s\n",
			i,value.Id,value.Name,value.STP,value.Interface)
	}
}

//var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}
//func Fields(s string) []string {
//	// First count the fields.
//	// This is an exact count if s is ASCII, otherwise it is an approximation.
//	n := 0
//	wasSpace := 1
//	// setBits is used to track which bits are set in the bytes of s.
//	setBits := uint8(0)
//	for i := 0; i < len(s); i++ {
//		r := s[i]
//		setBits |= r
//		isSpace := int(asciiSpace[r])
//		n += wasSpace & ^isSpace
//		wasSpace = isSpace
//	}
//
//	if setBits < utf8.RuneSelf { // ASCII fast path
//		a := make([]string, n)
//		na := 0
//		fieldStart := 0
//		i := 12
//		// Skip spaces in the front of the input.
//		for i < len(s) && asciiSpace[s[i]] != 0 {
//			i++
//		}
//		fieldStart = i
//		for i < len(s) {
//			if asciiSpace[s[i]] == 0 {
//				i++
//				continue
//			}
//			a[na] = s[fieldStart:i]
//			na++
//			i++
//			// Skip spaces in between fields.
//			for i < len(s) && asciiSpace[s[i]] != 0 {
//				i++
//			}
//			fieldStart = i
//		}
//		if fieldStart < len(s) { // Last field might end at EOF.
//			a[na] = s[fieldStart:]
//		}
//		return a
//	}
//
//	// Some runes in the input string are not ASCII.
//	return FieldsFunc(s, unicode.IsSpace)
//}
//
//func FieldsFunc(s string, f func(rune) bool) []string {
//	// A span is used to record a slice of s of the form s[start:end].
//	// The start index is inclusive and the end index is exclusive.
//	type span struct {
//		start int
//		end   int
//	}
//	spans := make([]span, 0, 32)
//
//	// Find the field start and end indices.
//	wasField := false
//	fromIndex := 0
//	for i, rune := range s {
//		if f(rune) {
//			if wasField {
//				spans = append(spans, span{start: fromIndex, end: i})
//				wasField = false
//			}
//		} else {
//			if !wasField {
//				fromIndex = i
//				wasField = true
//			}
//		}
//	}
//
//	// Last field might end at EOF.
//	if wasField {
//		spans = append(spans, span{fromIndex, len(s)})
//	}
//
//	// Create strings from recorded field indices.
//	a := make([]string, len(spans))
//	for i, span := range spans {
//		a[i] = s[span.start:span.end]
//	}
//
//	return a
//}
