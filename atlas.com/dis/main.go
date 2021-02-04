package main

import (
	"atlas-dis/database/monster_drop"
	"atlas-dis/domain"
	"atlas-dis/rest"
	json2 "atlas-dis/rest/json"
	"atlas-dis/retry"
	"bufio"
	"bytes"
	"gorm.io/driver/mysql"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)
import "gorm.io/gorm"

func connectToDatabase(attempt int) (bool, interface{}, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:the@tcp(atlas-db:3306)/atlas-dis?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		return true, nil, err
	}
	return false, db, err
}

func main() {
	l := log.New(os.Stdout, "dis ", log.LstdFlags|log.Lmicroseconds)

	r, err := retry.RetryResponse(connectToDatabase, 10)
	if err != nil {
		panic("failed to connect database")
	}
	db := r.(*gorm.DB)

	// Migrate the schema
	monster_drop.Migration(db)

	createRestService(l, db)

	initializeDrops(l, db)

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Println("[INFO] shutting down via signal:", sig)
}

func initializeDrops(l *log.Logger, db *gorm.DB) {
	s, err := monster_drop.GetAllMonsterDrops(db)
	if err != nil {
		l.Fatalf(err.Error())
	}
	if len(s) > 0 {
		return
	}

	jsonFile, err := os.Open("./drop_data.json")
	if err != nil {
		l.Fatalf(err.Error())
	}
	defer jsonFile.Close()

	input := struct {
		MonsterId       uint32 `json:"monsterId"`
		ItemId          uint32 `json:"itemId"`
		MinimumQuantity uint32 `json:"minimumQuantity"`
		MaximumQuantity uint32 `json:"maximumQuantity"`
		Chance          uint32 `json:"chance"`
	}{}

	var monsterDrops []domain.MonsterDrop
	reader := bufio.NewReader(jsonFile)
	for {
		var buffer bytes.Buffer
		var lin []byte
		var isPrefix bool

		for {
			lin, isPrefix, err = reader.ReadLine()
			buffer.Write(lin)
			if !isPrefix {
				break
			}
			if err != nil {
				if err != io.EOF {
					l.Fatalf(err.Error())
				}
				break
			}

		}
		lineObject := &input
		json2.FromJSON(lineObject, &buffer)

		md := domain.NewMonsterDropBuilder(0).
			SetMonsterId(lineObject.MonsterId).
			SetItemId(lineObject.ItemId).
			SetMinimumQuantity(lineObject.MinimumQuantity).
			SetMaximumQuantity(lineObject.MaximumQuantity).
			SetChance(lineObject.Chance).
			Build()
		monsterDrops = append(monsterDrops, md)

		if err == io.EOF {
			break
		}
	}

	err = monster_drop.BulkCreateMonsterDrop(db, monsterDrops)
	if err != nil {
		l.Fatalf(err.Error())
	}
}

func createRestService(l *log.Logger, db *gorm.DB) {
	rs := rest.NewServer(l, db)
	go rs.Run()
}
