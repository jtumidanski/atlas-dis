package main

import (
	"atlas-dis/database"
	"atlas-dis/drop"
	json2 "atlas-dis/json"
	"atlas-dis/logger"
	"atlas-dis/rest"
	"atlas-dis/tracing"
	"bufio"
	"bytes"
	"context"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
)
import "gorm.io/gorm"

const serviceName = "atlas-dis"

func main() {
	l := logger.CreateLogger(serviceName)
	l.Infoln("Starting main service.")

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	tc, err := tracing.InitTracer(l)(serviceName)
	if err != nil {
		l.WithError(err).Fatal("Unable to initialize tracer.")
	}
	defer func(tc io.Closer) {
		err := tc.Close()
		if err != nil {
			l.WithError(err).Errorf("Unable to close tracer.")
		}
	}(tc)

	db := database.Connect(l, database.SetMigrations(drop.Migration))

	rest.CreateService(l, db, ctx, wg, "/ms/dis", drop.InitResource)

	initializeDrops(l, db)

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Infof("Initiating shutdown with signal %s.", sig)
	cancel()
	wg.Wait()
	l.Infoln("Service shutdown.")
}

func initializeDrops(l logrus.FieldLogger, db *gorm.DB) {
	s, err := drop.GetAll(l, db)
	if err != nil {
		l.Fatalf(err.Error())
	}
	if len(s) > 0 {
		return
	}

	jsonFile, err := os.Open("./drop_data.json")
	if err != nil {
		l.WithError(err).Errorf("Unable to open drop data file.")
		return
	}
	defer jsonFile.Close()

	input := struct {
		MonsterId       uint32 `json:"monsterId"`
		ItemId          uint32 `json:"itemId"`
		MinimumQuantity uint32 `json:"minimumQuantity"`
		MaximumQuantity uint32 `json:"maximumQuantity"`
		Chance          uint32 `json:"chance"`
	}{}

	var monsterDrops []drop.Model
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

		md := drop.NewMonsterDropBuilder(0).
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

	err = drop.BulkCreateMonsterDrop(db, monsterDrops)
	if err != nil {
		l.Fatalf(err.Error())
	}
}
