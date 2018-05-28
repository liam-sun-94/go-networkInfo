package main

import "go-networkInfo/linuxNetwork"

func main(){
	//
	//fmt.Println("---------本机网桥信息-----------")
	//linuxNetwork.ShowBridge()
 	//fmt.Println(linuxNetwork.GetBridge())
	//
	//
	//fmt.Println("------------------------------")
	////linuxNetwork.ShowECI()
	//*        *
	//*        *
	//fmt.Println("----------显示Docker网桥信息---------")
	//dockerNetwork.ShowDockerBridge()
	//fmt.Println(dockerNetwork.GetDockerBridge())
	//
	//fmt.Println("----------显示Docker某一个的信息---------")
	//networkID:="191b8e4fd9152ef0901221b24c49a4081515a6d8585c253f232d28f2185f1fc9"
	//dockerNetwork.ShowDockerInspect(networkID)
	//fmt.Println(dockerNetwork.GetDockerInspect(networkID))

	linuxNetwork.GetECI()


}