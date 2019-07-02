# MetaKey使用
## 节点之间信息交互的方法


### 操作码opcode：
  + 标识这段信息处理哪一种操作，放在dht pmes中随KV对一起发送。对方节点接到信息之后，根据操作码来调用不同的回调函数完成操作。
  + `User_Init_Req` `Uer_Init_Res` `New_User_Notif` `TRestart`
  + `putto` `getfrom` `putlocal` `getlocal` 
  + `Proof` `Challenge` `Repair` `Delete_Block`
  + `Sync`


### 数据格式
  + `注意！` 这里分隔符用‘/’表示，在实际应用中，可能不使用'/'
  + key有固定的格式： `/mainID/keytype/arg1/arg2......`
    + `mainID`是本次操作主要节点的ID号，可以是UKP任何一个角色
    + `keytype`是这段信息的类型，回调函数根据不同的keytype做不同的操作
    + `arg`是操作数，具体信息不同，数量内容不一样
  + value记录本次数据传输中的主要数据，比如keeper的ID 挑战中的结果和签名等


### 具体操作以及信息汇总
  +  初始化流程
     +  keytype:`"init"`
     +  user：`User_Init_Req` key:`/UserID/"init"/keepercount/providercount` value:无
     +  keeper：`User_Init_Res` key:`/UserID/"init"/keepercount/providercount` value:`kid1kid2.../pid1pid2....`
     +  user：`New_User_Notif` key:`/UserID/"init"/keepercount/providercount` value:`kid1kid2.../pid1pid2....`
  + 元数据同步`Sync`
    + 涉及多种keytype
    + block位置信息 key:`/blockID/"block"` value:`pid/offset`
    + 挑战结果信息: key:`/uid/"chalres"/pid/kid/time` value: `length/result/proof`
    + 挑战汇总信息: key:`/uid/"chalpay"/pid/begin_time/end_time` value:`spacetime/signature/proof`
    + tendermint信息：key:`/pid/"tendermintinfo"` value:`id/ip/pubkey/p2pport/rpcport`
  + tendermint重新启动 `TRestart`
    + 与元数据同步中tendermint信息相同
  + 挑战流程
    + keeper：`Challenge` key:`/uid/"challenge"/h/chaltime` value:`userconfigByte`
    + provider:`Proof` key:`/uid/"proof"/FaultBlock/chaltime` value:`proof`
  + 修复流程`Repair`
    + key：block位置信息 value:`rpid`
  + 数据块删除`Delete_Block`
    + key：block位置信息 value:无
  + 其他KV数据操作`putto` `getfrom` `putlocal` `getlocal`，部分数据类型在之前写过
    + block位置信息 key:`/blockID/"block"` value:`pid/offset`
    + config信息 key:`PeerID/"userconfig"` value:具体数据
    + 角色信息 key:`PeerID/"roleinfo"` value:`user`/`keeper`/`provider`
    + legerinfo信息 key:`PeerID/legerinfo` value:具体数据
    + 与角色相连的其他节点信息 
      + user key:`PeerID/'uids'` value:`uid1uid2uid3....`
      + keeper key:`PeerID/'kids'` value:`kid1kid2kid3....`
      + provider key:`PeerID/'pids'` value:`pid1pid2pid3....`
    + 节点的tendermint信息 key:`PeerID/"bft"` value:`"sample"`或`bft/IP:p2pport/IP:rpcport`


###使用逻辑
  + 数据结构
    ```go
    type KeyMeta struct {
        mid     string 
        keytype string
        options []interface{}
    }
    ```
  + 收到pmes，根据opcode不同，进入不同的回调函数，key整理成KeyMeta结构，进行之后的操作