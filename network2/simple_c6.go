func SendSocketMsg(uid, zid int, cmdName string, para map[string]interface{}) (error, map[string]interface{}) {
	var err error
	var serverIp string
	var port string
	var zidStr = strconv.Itoa(zid)
	serverIp, err = GetMySectionConfig("db", "z"+zidStr, "server")
	if err != nil {
		return err, nil
	}

	port, err = GetMySectionConfig("db", "z"+zidStr, "port")
	if err != nil {
		return err, nil
	}

	conn, err := net.Dial("tcp", serverIp+":"+port)
	checkError(err)

	secret := beego.AppConfig.String("game-server" + ":" + "secret")
	request := map[string]interface{}{
		"uid":    uid,
		"zoneid": zid,
		"cmd":    cmdName,
		"params": para,
		"secret": secret,
	}

	requestByte, _ := json.Marshal(request)
	fmt.Println(serverIp + ":" + port + "===>" + string(requestByte))

	_, err = conn.Write([]byte("1 " + string(requestByte) + "\r\n"))
	checkError(err)
	defer conn.Close()

	const RECV_BUF_LEN = 5
	initheader := make([]byte, RECV_BUF_LEN)
	n, _ := conn.Read(initheader)
	fmt.Println(n)
	header := initheader[1:4]
	fmt.Println("header:" + string(header))

	lenth := binary.LittleEndian.Uint16(header)
	fmt.Println(lenth)

	headerLen := uint16(n)
	fmt.Println(headerLen, lenth)

	lenth = lenth - headerLen
	var allMsg string
	var total uint16
	for {
		if lenth > 0 && lenth > total {
			nextLength := lenth - total
			if nextLength > uint16(1024) {
				nextLength = uint16(1024)
			}
			if nextLength > 0 {
				nextReadRet := make([]byte, nextLength)
				m, _ := conn.Read(nextReadRet)
				allMsg += string(nextReadRet[0:m])
				total += uint16(m)

				fmt.Println(m)
			}
		} else {
			break
		}
	}

	fmt.Println(allMsg)
	return nil, HandleMsg(allMsg)
}