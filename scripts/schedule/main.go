package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gocarina/gocsv"
	"gopkg.in/yaml.v2"
)

const (
	goConferenceEventDate      = "2022-04-23"
	goConferenceEventStartTime = "10:00"
)

type Schedule struct {
	Day   string
	Start string
	Rooms []*Room
}

type Room struct {
	Room  string
	Order int `yaml:"-"`
	Slots []*Slot
}

type Slot struct {
	Slot string
	Talk string
}

type ScheduleCSVRow struct {
	Room      string `csv:"room"`
	RoomOrder int    `csv:"room_order"`
	Slot      string `csv:"slot"`
	Talk      string `csv:"talk"`
}

const fileName = "schedule"

func generate(c []*Schedule) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirPath := filepath.Join(wd, "data")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(filepath.Join(dirPath, fmt.Sprintf("%s.yml", fileName)), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	out, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}

	rd := strings.NewReader(string(out))

	_, err = io.Copy(f, rd)
	if err != nil {
		panic(err)
	}
}

func convertScheduleCSVRowsToSchedules(rows []*ScheduleCSVRow) []*Schedule {
	roomIDRowsMap := make(map[string][]*ScheduleCSVRow)
	for _, row := range rows {
		roomIDRowsMap[row.Room] = append(roomIDRowsMap[row.Room], row)
	}
	rooms := make([]*Room, 0, len(roomIDRowsMap))
	for roomID, rows := range roomIDRowsMap {
		slots := make([]*Slot, len(rows))
		for i, row := range rows {
			slots[i] = &Slot{
				Slot: row.Slot,
				Talk: row.Talk,
			}
		}
		room := &Room{
			Room:  roomID,
			Order: rows[0].RoomOrder,
			Slots: slots,
		}
		rooms = append(rooms, room)
	}
	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i].Order < rooms[j].Order
	})
	// 1日開催なので slice の要素は1つ
	return []*Schedule{{
		Day:   goConferenceEventDate,
		Start: goConferenceEventStartTime,
		Rooms: rooms,
	}}
}

func main() {
	f, err := os.Open(fmt.Sprintf("./raw_data/%s.csv", fileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var rows []*ScheduleCSVRow
	if err := gocsv.UnmarshalFile(f, &rows); err != nil {
		panic(err)
	}
	generate(convertScheduleCSVRowsToSchedules(rows))
}
