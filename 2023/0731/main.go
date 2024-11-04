package main

import (
	"database/sql"
	"github.com/rs/zerolog/log"
	_ "github.com/mattn/go-sqlite3"
)

/**
*
* Author: Marek
* Date: 2023-07-31 13:10
* Email: 364021318@qq.com
*
 */

var db *sql.DB

func InitDB() *sql.DB {
	if db != nil {
		return db
	}

	var err error

	db, err = sql.Open("sqlite3", "./hub.db")
	if err != nil {
		log.Fatal().Err(err).Msg("db connect failed")
	}

	// 测试连接是否成功
	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("ping db failed")
	}

	log.Info().Msg("database connected")

	// 创建数据表
	_, err = db.Exec(`
	CREATE TABLE devices (
	device_id TEXT PRIMARY KEY,
	gateway TEXT,
	serial TEXT,
	status INTEGER,
	dimming_level INTEGER,
	provider TEXT,
	temperature REAL,
	fan_speed TEXT
	)
`)
	if err != nil && err.Error() != "table devices already exists" {
		log.Error().Err(err).Msg("create table failed:\n")
	}

	return db
}

type Device struct {
	DeviceID     string
	Gateway      string
	Serial       string
	Provider     string
	FanSpeed     string
	Status       int
	DimmingLevel int
	Temperature  float32
}

func InsertDevice(device Device) error {
	_, err := db.Exec("INSERT INTO devices(device_id, gateway, serial, "+
		"status, dimming_level, provider, temperature, fan_speed) VALUES(?,?,?,?,?,?,?,?)",
		device.DeviceID, device.Gateway, device.Serial, device.Status, device.DimmingLevel,
		device.Provider, device.Temperature, device.FanSpeed)

	return err
}

func UpdateRow(device Device) error {
	_, err := db.Exec("UPDATE devices  set gateway = ?, serial = ?, status = ?, "+
		"dimming_level = ?, provider = ?, temperature = ?, "+
		"fan_speed = ? where device_id = ?",
		device.Gateway, device.Serial, device.Status, device.DimmingLevel,
		device.Provider, device.Temperature, device.FanSpeed, device.DeviceID)

	return err
}

func GetRowByDeviceID(deviceID string) (Device, error) {
	row, err := db.Query("SELECT device_id,gateway,serial,provider,fan_speed,status,dimming_level,temperature FROM devices where device_id = ? limit 1", deviceID)
	if err != nil {
		return Device{}, err
	}

	if row.Err() != nil {
		return Device{}, row.Err()
	}

	var device Device

	if row.Next() {
		err := row.Scan(&device.DeviceID, &device.Gateway, &device.Serial,
			&device.Provider, &device.FanSpeed, &device.Status, &device.DimmingLevel,
			&device.Temperature)
		if err != nil {
			return Device{}, err
		}
	}

	row.Close()

	log.Info().Int("status", device.Status).Str("deviceID", device.DeviceID).Msg("GetRowByDeviceID")

	return device, nil
}

func main() {
	_ = InitDB()
	log.Info().Msg("get one device info")
	device := Device{
		DeviceID:     "014a79e7-a915-414d-8821-b4601b7f4abb",
		Gateway:      "3C6A2CFFFED21FBB",
		Serial:       "D7B8EC07004B1201",
		Provider:     "OWON",
		FanSpeed:     "low",
		Status:       1,
		DimmingLevel: 10,
		Temperature:  0.9,
	}
	dev, err := GetRowByDeviceID(device.DeviceID)
	if dev.DeviceID == "" {
		log.Error().Err(err).Msg("query get device err")
		err = InsertDevice(device)
		if err != nil {
			log.Error().Err(err).Str("deviceID", device.DeviceID).Msg("insert device failed")
		} else {
			log.Info().Str("deviceID", device.DeviceID).Msg("insert device succes")
		}
	} else {
		err = UpdateRow(device)
		if err != nil {
			log.Error().Err(err).Str("deviceID", device.DeviceID).Msg("update device failed")
		} else {
			log.Info().Str("deviceID", device.DeviceID).Msg("update device succes")
		}
	}
//	db, err := sql.Open("sqlite3", "./test.db")
//	if err != nil {
//		log.Fatal().Msg("db connect failed")
//	}
//
//	defer db.Close()
//
//	// 测试连接是否成功
//	err = db.Ping()
//	if err != nil {
//		log.Fatal().Msg("ping db failed")
//	}
//
//	log.Info().Msg("database connected")
//
//	// 创建数据表
//	_, err = db.Exec(`
//	CREATE TABLE devices (
//	device_id TEXT PRIMARY KEY,
//	gateway TEXT,
//	serial TEXT,
//	status INTEGER,
//	dimming_level INTEGER,
//	provider TEXT,
//	temperature REAL,
//	fan_speed TEXT
//	)
//`)
//if err != nil {
//	log.Error().Err(err).Msg("create table failed:\n")
//	return
//}
//
//	// 插入数据
//	_, err = db.Exec("INSERT INTO devices(device_id, gateway, serial, status, dimming_level, provider, temperature, fan_speed) VALUES(?,?,?,?,?,?,?,?)", "014a79e7-a915-414d-8821-b4601b7f4abb", "3C6A2CFFFED21FBB","D7B8EC07004B1201",1,10,"OWON", 0.9, "low")
//	if err != nil {
//		log.Error().Err(err).Msg("insert data failed:")
//		return
//	}
//
//	//lastinsertId, _ := res.LastInsertId() // 获取自增ID
//	//log.Info().Int64("id:", lastinsertId).Msg("last insert id")
//
//	rows, err := db.Query("select device_id,temperature FROM devices")
//	if err != nil {
//		log.Error().Err(err).Msg("query data failed")
//		return
//	}
//
//	defer rows.Close()
//
//	for rows.Next() {
//		var device_id string
//		var temperature float32
//
//		err := rows.Scan(&device_id, &temperature)
//		if err != nil {
//			log.Error().Err(err).Msg("get data failed")
//			return
//		}
//
//		log.Info().Str("device_id", device_id).Float32("temperature", temperature).Msg("query res:")
//	}

	//_, err = db.Exec("UPDATE devices  SET gateway = ?, serial = ?, status = ?, dimming_level = ?, provider = ?, temperature = ?, fan_speed = ? where device_id = ?",
	//	"3C6A2CFFFED21FBB","D7B8EC07004B1201",1,10,"OWON", 0.99, "low", "014a79e7-a915-414d-8821-b4601b7f4abb")
	//log.Error().Err(err).Msg("333")
}