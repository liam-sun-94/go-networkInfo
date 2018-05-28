package dockerNetwork

import (
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"fmt"
	"golang.org/x/net/context"
)

//显示docker网桥信息
func ShowDockerBridge(){
	netList:=GetDockerBridge()

	fmt.Println("--------------未辨别--------------------")
	for _,netListitem :=range netList{
		fmt.Println(netListitem)
	}
}
//获得docker网桥信息
func GetDockerBridge()[]types.NetworkResource{
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	netList,err:=cli.NetworkList(ctx,types.NetworkListOptions{})
	if err!=nil{
		panic(err)
	}

	return netList
}



//获取某一个链接的详细信息，通过ID查询
func GetDockerInspect(networkID string) types.NetworkResource {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	dockerInspect,err:=cli.NetworkInspect(ctx,networkID,types.NetworkInspectOptions{})
	if err!=nil{
		panic(err)
	}

	return dockerInspect
}
//显示某一个链接的详细信息，通过ID显示
func ShowDockerInspect(networkID string) {

	dockerInspect := GetDockerInspect(networkID)

	fmt.Println(dockerInspect)
}