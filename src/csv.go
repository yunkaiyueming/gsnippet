package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Order struct {
	Id, OrderNum, Address, Cretime string
}

func main() {
	f, err := os.Create("order2.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)

	data := make([]Order, 0, 20)
	for i := 1; i < 20; i++ {
		tmp := Order{strconv.Itoa(i), "AA00" + strconv.Itoa(i), "北京市", "2015-12-12"}
		data = append(data, tmp)
	}

	fmt.Println(data)

	w.Write([]string{"Id", "订单号", "订单地址", "订单时间"})
	for _, order_info := range data {
		w.Write([]string{order_info.Id, order_info.OrderNum, order_info.Address, order_info.Cretime})
	}
	w.Flush()
}
