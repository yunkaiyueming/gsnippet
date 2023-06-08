package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	// "github.com/tengattack/gluasql"

	mysql "github.com/tengattack/gluasql/mysql"
	luajson "github.com/layeh/gopher-json"
	
	// redslib "github.com/go-redis/redis"

	redis "gsnippet/gopherredis"
	
	tcp "github.com/vadv/gopher-lua-libs/tcp"
)

func main(){
	// execLuaStr()
	execLuaFile()
}

//go给lua里调用的函数
func Double(L *lua.LState) int {
    lv := L.ToInt(1)             /* get argument */
    L.Push(lua.LNumber(lv * 2)) /* push result */
    return 1                     /* number of results */
}

func execLuaStr(){
	L := lua.NewState() // 创建一个lua解释器实例
	defer L.Close()
        // 执行字符串语句
	if err := L.DoString(`print("hello")`); err != nil {
		panic(err)
	}
}

func execLuaFile(){
	L := lua.NewState()
	L.SetGlobal("double", L.NewFunction(Double)) 
	L.PreloadModule("mysql", mysql.Loader)
	L.PreloadModule("redis", redis.Loader)
	L.PreloadModule("JSON", luajson.Loader)
	L.PreloadModule("tcp", tcp.Loader)

	defer L.Close()
	// 加载fib.lua
	if err := L.DoFile("dispatch.lua"); err != nil {
		panic(err)
	}

	requestStr := "cmd:user_login&uid:10001&zid:1"
	
	// 调用fib(n)
	err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("dispatch"), // 获取fib函数引用
		NRet:    1,                  // 指定返回值数量
		Protect: true,               // 如果出现异常，是panic还是返回err
	}, lua.LString(requestStr),lua.LNumber(1) ) // 传递输入参数n=10
	if err != nil {
		panic(err)
	}

	// 获取返回结果
	ret := L.Get(-1)
	// 从堆栈中扔掉返回结果
	L.Pop(1)

	// 打印结果
	res, ok := ret.(lua.LString)
	if ok {
		fmt.Println("go获取结果：",res)
	} else {
		fmt.Println("unexpected result")
	}
}