package service

// GetSyncMQByName  根据站点名称停止 es_sync_mq 服务
//func GetSyncMQByName(name string) (*config.EsSyncMQ, error) {
//	for _, v := range global.Cfg.EsSyncMq {
//		if v.Name == name {
//			return &v, nil
//		}
//	}
//	return nil, errors.New("sync-mq not found: " + name)
//}
//
//// StopSyncMQ  停止 es_sync_mq 进程
//func StopSyncMQ(mq *config.EsSyncMQ) error {
//	// 使用 binary + workdir 精确匹配，避免误杀
//	cmd := exec.Command(
//		"pkill",
//		"-f",
//		fmt.Sprintf("%s -config %s", mq.Binary, mq.Config),
//	)
//	cmd.Dir = mq.WorkDir
//	return cmd.Run()
//}
