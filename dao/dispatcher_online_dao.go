package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lemon-robot-golang-commons/utils/lru_date"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
	"lemon-robot-server/sysinfo"
	"time"
)

type DispatcherOnlineDao struct{}

func NewDispatcherOnlineDao() *DispatcherOnlineDao {
	return &DispatcherOnlineDao{}
}

func (i *DispatcherOnlineDao) Delete(query interface{}, where ...interface{}) {
	db.Db().Where(query, where).Delete(&entity.DispatcherOnline{})
}

func (i *DispatcherOnlineDao) DeleteByServerNodeMachineSign(nodeMachineSign string) {
	i.Delete("bind_server_machine_sign = ?", nodeMachineSign)
}

func (i *DispatcherOnlineDao) Save(entity *entity.DispatcherOnline) *gorm.DB {
	return db.Db().Save(entity)
}

func (i *DispatcherOnlineDao) FirstByExample(example *entity.DispatcherOnline) entity.DispatcherOnline {
	result := entity.DispatcherOnline{}
	db.Db().First(&result, example)
	return result
}

func (i *DispatcherOnlineDao) CountByUserExample(example *entity.DispatcherOnline) int {
	var count int
	db.Db().Model(&entity.DispatcherOnline{}).Where(example).Count(&count)
	return count
}

func (i *DispatcherOnlineDao) ClearAllOffline() {
	db.Db().Exec(`DELETE LDO
		FROM lr_dispatcher_onlines LDO,
			 lr_server_nodes LSN
		WHERE LDO.bind_server_machine_sign = LSN.machine_sign
		AND LSN.active_time < ?`, lru_date.GetInstance().CalculateTimeByDurationStr(
		time.Now(), fmt.Sprintf("-%ds", sysinfo.LrServerConfig().ClusterNodeActiveInterval*2)))
}
