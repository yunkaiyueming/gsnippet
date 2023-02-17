
// Golang实现贪心算法
package main 
 
import (
  "fmt"
)
  
// 定义结构体 
type Data struct {
    index int // 索引
    value int // 优先级
}
 
func main() {
    // 定义一个Data类型的切片, 用于存储优先级
    dataSlice := []Data{
    {
        index: 0, 
        value: 2,
    },
    {
        index: 1, 
        value: 4,
    },
    {
        index: 2, 
        value: 3,
    },
    {
        index: 3, 
        value: 1,
    },
  }
  
  // 存储最终结果的切片 
  var result []int
  
  // 贪心算法实现
  // 每次取出最大优先级，放入结果，再从原切片中删除
  // 循环结束条件: 数据为空时结束循环
  for len(dataSlice) > 0 {
    // 每次取出优先级最大的
    maxValue := 0
    maxIndex := 0
    for i := 0; i < len(dataSlice); i++ {
        if dataSlice[i].value > maxValue {
            maxValue = dataSlice[i].value
            maxIndex = dataSlice[i].index
        }
    }
    // 将最大优先级放入结果切片
    result = append(result, maxIndex)
    // 从原切片中删除
    dataSlice = append(dataSlice[:maxIndex], dataSlice[maxIndex+1:]...)
  }
  
  fmt.Println(result) // 输出结果
} 