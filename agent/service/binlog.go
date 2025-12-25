package service

type BinlogPosition struct {
	File     string `json:"file"`
	Position uint64 `json:"position"`
}

//func GetBinlogPosition() (*BinlogPosition, error) {
//	cfg := global.Cfg.EsSyncMq
//
//	row := db.QueryRow("SHOW MASTER STATUS")
//
//	var file string
//	var position uint64
//	var ignore1, ignore2, ignore3, ignore4 sql.NullString
//
//	if err := row.Scan(&file, &position, &ignore1, &ignore2, &ignore3, &ignore4); err != nil {
//		return nil, err
//	}
//
//	return &BinlogPosition{
//		File:     file,
//		Position: position,
//	}, nil
//}
