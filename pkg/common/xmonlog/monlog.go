package xmonlog

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"lark/pkg/common/xmongo"
	"time"
)

type MonlogMgr struct {
	db             *mongo.Database
	LogChan        chan *MonLog
	AutoCommitChan chan *LogBatch
	BatchSize      int
	CommitTimeout  int //毫秒
}

type LogBatch struct {
	LogType string
	Logs    []interface{}
}

type MonLog struct {
	LogType string
	Log     interface{}
}

var logMgr *MonlogMgr

func NewMonlog(batchSize int, commitTimeout int) *MonlogMgr {
	var (
		err error
	)
	logMgr = &MonlogMgr{
		LogChan:        make(chan *MonLog, 1000),
		AutoCommitChan: make(chan *LogBatch, 1000),
		BatchSize:      batchSize,
		CommitTimeout:  commitTimeout}
	logMgr.db = xmongo.GetDB()
	if err == nil {
		go logMgr.polling()
	}
	return logMgr
}

/*
var instance *MongoLog
var once sync.Once
func Shared() *MongoLog {
	once.Do(func() {
		instance = &MongoLog{LogChan: make(chan *MonLog, 1000),
			AutoCommitChan: make(chan *LogBatch, 1000),
			BatchSize:      100,
			CommitTimeout:  1000}
	})
	return instance
}
*/

func (s *MonlogMgr) SetDB(db *mongo.Database) {
	if db == nil {
		return
	}
	if s.db != nil {
		return
	}
	s.db = db
	go s.polling()
}

func (s *MonlogMgr) polling() {
	var (
		log          *MonLog
		batchMaps    = map[string]*LogBatch{}
		timerMaps    = map[string]*time.Timer{}
		timeoutBatch *LogBatch
	)

	for {
		select {
		case log = <-s.LogChan:
			if _, ok := batchMaps[log.LogType]; ok == false {
				batchMaps[log.LogType] = &LogBatch{LogType: log.LogType, Logs: make([]interface{}, 0)}
				timerMaps[log.LogType] = time.AfterFunc(time.Duration(s.CommitTimeout)*time.Millisecond,
					func(batch *LogBatch) func() {
						return func() {
							s.AutoCommitChan <- batch
						}
					}(batchMaps[log.LogType]))
			}

			batchMaps[log.LogType].Logs = append(batchMaps[log.LogType].Logs, log.Log)
			if len(batchMaps[log.LogType].Logs) >= s.BatchSize {
				s.InsertMany(log.LogType, batchMaps[log.LogType].Logs)
				timerMaps[log.LogType].Stop()
				delete(batchMaps, log.LogType)
			}
		case timeoutBatch = <-s.AutoCommitChan:
			if _, ok := batchMaps[timeoutBatch.LogType]; ok == false {
				continue
			}
			if batchMaps[timeoutBatch.LogType] != timeoutBatch {
				continue
			}
			s.InsertMany(log.LogType, timeoutBatch.Logs)
			delete(batchMaps, timeoutBatch.LogType)
		default:
		}
	}
}

func (s *MonlogMgr) InsertMany(table string, documents []interface{}) {
	if s.db == nil {
		return
	}
	s.db.Collection(table).InsertMany(context.TODO(), documents)
}

func Insert(table string, obj interface{}) {
	logMgr.LogChan <- &MonLog{LogType: table, Log: obj}
}
