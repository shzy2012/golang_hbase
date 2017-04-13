package main

import (
	"context"
	"fmt"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

func main() {
	//10.0.12.207:22
	client := gohbase.NewClient("10.0.12.206:2181")
	if client != nil {

		fmt.Println("client")
		// Values maps a ColumnFamily -> Qualifiers -> Values.
		//values := map[string]map[string][]byte{"cf": map[string][]byte{"a": []byte{0}}}

		// pFilter := filter.NewPrefixFilter([]byte("7"))
		// scanRequest, err := hrpc.NewScanStr(context.Background(), "scores")
		// scanRsp, err := client.Scan(scanRequest)
		// if err != nil {
		// 	fmt.Println(scanRsp)
		// }

		getRequest, err := hrpc.NewGetStr(context.Background(), "scores", "Tom")
		getRsp, err := client.Get(getRequest)
		if err == nil {
			for index, ele := range getRsp.Cells {
				fmt.Println(index)
				fmt.Println(string(ele.Family))
				fmt.Println(string(ele.Qualifier))
				fmt.Println(string(ele.Row))
				fmt.Println(string(ele.Value))
			}
		}
	}
	fmt.Println("hello")
}
