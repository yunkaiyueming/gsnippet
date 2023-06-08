package gopherredis

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/go-redis/redis"
	lua "github.com/yuin/gopher-lua"
)

// Loader Loader is the module loader
func Loader(L *lua.LState) int {
	mt := L.NewTypeMetatable("redis")

	mtindex := L.NewTable()
	L.SetField(mtindex, "docmd", L.NewFunction(doCmd))
	L.SetField(mtindex, "close", L.NewFunction(closeClient))

	L.SetField(mt, "__index", mtindex)

	t := L.NewTable()
	// L.SetFuncs(t, api)
	L.SetField(t, "new", L.NewFunction(newRedis))
	L.Push(t)
	return 1
}

func newRedis(L *lua.LState) int {
	opt := L.CheckTable(1)
	host := "127.0.0.1"
	port := "6379"
	passwd := ""
	db := 0

	opt.ForEach(func(k, v lua.LValue) {
		k1, ok := k.(lua.LString)
		if !ok {
			L.ArgError(1, "only string allowed in table index")
			return
		}

		switch string(k1) {
		case "host":
			v1, ok := v.(lua.LString)
			if !ok {
				L.ArgError(1, "string required for host")
				return
			}
			host = string(v1)
		case "port":
			switch v1 := v.(type) {
			case lua.LString:
				port = string(v1)
			case lua.LNumber:
				port = fmt.Sprintf("%d", int64(v1))
			default:
				L.ArgError(1, "string or number required for port")
				return
			}
		case "passwd", "password", "pass":
			v1, ok := v.(lua.LString)
			if !ok {
				L.ArgError(1, "string required for passwd")
				return
			}
			passwd = string(v1)
		case "index", "db":
			switch v1 := v.(type) {
			case lua.LString:
				db1, err := strconv.Atoi(string(v1))
				if err == nil {
					db = db1
				}
			case lua.LNumber:
				db = int(v1)
			default:
				L.ArgError(1, "string or number required for index")
				return
			}
		}
	})

	fmt.Println(host, port, passwd, db)

	r := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: passwd,
		DB:       db,
	})

	ud := L.NewUserData()
	ud.Value = r
	L.SetMetatable(ud, L.GetTypeMetatable("redis"))
	L.Push(ud)
	return 1
}

func closeClient(L *lua.LState) int {
	ud := L.CheckUserData(1)
	r := ud.Value.(*redis.Client)
	r.Close()
	return 0
}

func doCmd(L *lua.LState) int {
	ud := L.CheckUserData(1)
	r := ud.Value.(*redis.Client)
	args := []interface{}{}
	for i := 2; i <= L.GetTop(); i++ {
		a := L.Get(i)
		switch a1 := a.(type) {
		case lua.LString:
			args = append(args, string(a1))
		case lua.LNumber:
			args = append(args, int64(a1))
		case *lua.LTable:
			a1.ForEach(func(k, v lua.LValue) {
				switch k1 := k.(type) {
				case lua.LString:
					args = append(args, string(k1))
				case lua.LNumber:
				default:
					L.ArgError(i, "only string or number index allowed in table")
					return
				}

				switch v1 := v.(type) {
				case lua.LString:
					args = append(args, string(v1))
				case lua.LNumber:
					args = append(args, int64(v1))
				default:
					L.ArgError(i, "only string or number value allowed in table")
					return
				}
			})
		default:
			L.ArgError(i, "need string, number or table")
			return 0
		}
	}
	fmt.Println(args...)
	res, err := r.Do(args...).Result()
	if err != nil {
		if err == redis.Nil {
			L.Push(lua.LNil)
			return 1
		}
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	v := toLValue(reflect.ValueOf(res), L)
	L.Push(v)
	return 1
}

func toLValue(a reflect.Value, L *lua.LState) lua.LValue {
	switch a.Kind() {
	case reflect.String:
		return lua.LString(a.String())
	case reflect.Slice:
		t := L.NewTable()
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i)
			v1 := toLValue(reflect.ValueOf(v.Interface()), L)
			t.RawSetInt(i+1, v1)
		}
		return t
	case reflect.Int, reflect.Int64:
		return lua.LNumber(a.Int())
	default:
		fmt.Println("unhandled type %v", a.Kind())
	}
	return lua.LNil
}