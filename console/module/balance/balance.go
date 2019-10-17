package balance

import (
	"errors"
	"fmt"
	"time"

	dao_balance "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/dao-balance"

	"github.com/eolinker/goku-api-gateway/console/module/service"
	driver2 "github.com/eolinker/goku-api-gateway/server/driver"
	entity "github.com/eolinker/goku-api-gateway/server/entity/balance-entity-service"
)

//Add 新增负载
func Add(info *Param) (string, error) {
	serviceInfo, err := service.Get(info.ServiceName)
	if err != nil {
		return fmt.Sprintf("serviceName:%s", err.Error()), err
	}
	switch serviceInfo.Type {
	case driver2.Static:
		{
			if info.Static == "" && info.StaticCluster == "" {
				return "param:static 和 staticCluster 不能同时为空", errors.New("param:static 和 staticCluster 不能同时为空")
			}
			now := time.Now().Format("2006-01-02 15:04:05")
			result, err := dao_balance.AddStatic(info.Name, info.ServiceName, info.Static, info.StaticCluster, info.Desc, now)

			return result, err
		}
	case driver2.Discovery:
		{
			if info.AppName == "" {
				return "param:appName 不能为空", errors.New("param:appName 不能为空")
			}
			now := time.Now().Format("2006-01-02 15:04:05")
			result, err := dao_balance.AddDiscovery(info.Name, info.ServiceName, info.AppName, info.Desc, now)

			return result, err
		}

	}

	return "无效serviceName", errors.New("invalid serviceName")
}

//Save 保存服务发现
func Save(info *Param) (string, error) {
	serviceInfo, err := service.Get(info.ServiceName)
	if err != nil {
		return fmt.Sprintf("serviceName:%s", err.Error()), err
	}
	switch serviceInfo.Type {
	case driver2.Static:
		{
			if info.Static == "" && info.StaticCluster == "" {
				return "param:static 和 staticCluster 不能同时为空", errors.New("param:static 和 staticCluster 不能同时为空")
			}
			now := time.Now().Format("2006-01-02 15:04:05")
			result, err := dao_balance.SaveStatic(info.Name, info.ServiceName, info.Static, info.StaticCluster, info.Desc, now)

			return result, err
		}
	case driver2.Discovery:
		{
			if info.AppName == "" {
				return "param:appName 不能为空", errors.New("param:appName 不能为空")
			}
			now := time.Now().Format("2006-01-02 15:04:05")
			result, err := dao_balance.SaveDiscover(info.Name, info.ServiceName, info.AppName, info.Desc, now)

			return result, err
		}

	}

	return "无效serviceName", errors.New("invalid serviceName")
}

//Get 通过负载名称获取负载信息
func Get(name string) (*Info, error) {
	b, e := dao_balance.Get(name)
	if e != nil {
		return nil, e
	}

	return ReadInfo(b), nil
}

//Search 关键字获取负载列表
func Search(keyworkd string) ([]*Info, error) {
	var entities []*entity.Balance
	if keyworkd == "" {
		es, e := dao_balance.GetAll()
		if e != nil {
			return nil, e
		}
		entities = es
	} else {
		es, e := dao_balance.Search(keyworkd)
		if e != nil {
			return nil, e
		}
		entities = es
	}

	infos := make([]*Info, 0, len(entities))

	for _, ent := range entities {
		infos = append(infos, ReadInfo(ent))
	}
	return infos, nil
}

//GetAll 获取所有负载列表
func GetAll() ([]*Info, error) {

	entities, e := dao_balance.GetAll()
	if e != nil {
		return nil, e
	}
	infos := make([]*Info, 0, len(entities))

	for _, ent := range entities {
		infos = append(infos, ReadInfo(ent))
	}
	return infos, nil
}

//Delete 删除负载
func Delete(name string) (string, error) {
	result, err := dao_balance.Delete(name)

	return result, err
}

//GetBalancNames 获取负载名称列表
func GetBalancNames() (bool, []string, error) {
	return dao_balance.GetBalanceNames()
}

//BatchDeleteBalance 批量删除负载
func BatchDeleteBalance(balanceNames []string) (string, error) {
	result, err := dao_balance.BatchDelete(balanceNames)

	return result, err
}
