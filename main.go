package main

import (
  "./proto"
  //"encoding/json"
  "fmt"
  "log"
  "net"
  "time"
  )

func main() {

  url := "testnet43.hedera.com:80"

  conn, err := net.Dial("tcp", url)
  if err != nil { log.Fatalln(err) }
  defer conn.Close()

  fmt.Println(conn)

  nodeShard:=0
  nodeRealm:=0
  // nodeNum:=2
  // pubKey:=""
  // privateKey:=""

  // payAccShard:=0
  // payAccRealm:=0
  // payAccNum:=0

  timeNow := time.Now()
  timeStamp := proto.Timestamp{ timeNow.Unix(), int32(timeNow.UnixNano()) }

  key := proto.Key{"isKey_Key"}
  keyList := []proto.Key{key}

  cc := proto.FileCreateTransactionBody{&timeStamp, &keyList, []byte("EMR json content"), nodeShard, nodeRealm, "NewRealmAdminKey"}

  fmt.Println(cc)

  fmt.Println("done")

}// .main

