package linuxNetwork

import (
	"go-networkInfo/cmd"
	"fmt"
	"strings"
	"net"
	"regexp"
)
//  bridge 结构体
//  Id string         ID 	字符串
//	Name string       名称  	字符串
//	Driver string     类型  	字符串
//	Scope string      范围  	字符串
type Bridge struct {
	Id 			string
	Name 		string
	STP 		string
	Interface 	string
}

type InterInfo struct{
	Index  int
	Name   string
	HWAddr string
	Flags  string
	IP     string
	Ipv6   string
}

//Linux中 主机须有ipconfig
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


//Linux中获得主机网桥信息  主机须有brctl
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
//Linux中显示主机网桥信息
func ShowBridge(){

	ttt:=GetBridge()
	fmt.Printf("%-8s%-20s%-10s%-15s%-15s\n",
		"Number","Id","Name","STP","Interface")
	for i,value :=range ttt{

		fmt.Printf("%-8d%-20s%-20s%-15s%-15s\n",
			i,value.Id,value.Name,value.STP,value.Interface)
	}
}



//Linux/windows中获得本机的网络信息
func GetInterfaces() []InterInfo{
	var arraysIface []InterInfo
	interfaces, err := net.Interfaces()
	if err!=nil{
		panic(err)

	}

	for i :=range interfaces{
		var Iface InterInfo
		addrs, err := interfaces[i].Addrs()
		if err != nil {
			panic(err)
		}

		Iface.Name=interfaces[i].Name
		Iface.Index=interfaces[i].Index
		Iface.Flags=interfaces[i].Flags.String()
		Iface.HWAddr=interfaces[i].HardwareAddr.String()

		for _,v := range addrs{
			//Iface.IP+=v.String()+","
			reg := regexp.MustCompile(`:`)
			result:=reg.FindAllString(v.String(), -1)
			//fmt.Printf("%q\n", reg.FindAllString(v.String(), -1))
			if len(result)>0{
				Iface.Ipv6=v.String()
			}else {
				Iface.IP=v.String()
			}
		}


		//if len(addrs)>1{
		//	Iface.Ipv6=addrs[1].String()
		//}
		//

		arraysIface = append(arraysIface,Iface)
	}

    return arraysIface
}
//Linux/windows中获得本机的网络信息
func ShowInterface(){
	arraysIface:=GetInterfaces()

	for _,value:=range arraysIface{
		fmt.Printf("Index	: %v\n",value.Index)
		fmt.Printf("Name 	: %v\n",value.Name)
		fmt.Printf("IP   	: %v\n",value.IP)
		fmt.Printf("Ipv6	: %v\n",value.Ipv6)
		fmt.Printf("HWAddr	: %v\n",value.HWAddr)
		fmt.Printf("Flags	: %v\n",value.Flags)
		fmt.Println("**********************************************************")
	}

}