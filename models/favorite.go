package models

import (
	"bus-booking/util"
	"errors"
	"log"
)

func AllFavorites(session *string, b *[]Bus) error {
	var user User
	err := NowUser(&user, session)
	util.Report(err)
	stmt, err := util.DB.Prepare(`SELECT buses.* FROM favorites, buses
		WHERE favorites.bus_id = buses.id AND user_id = ?`)
	util.Report(err)
	var bus Bus
	rows, err := stmt.Query(user.UserID)
	util.Report(err)
	for rows.Next() {
		err := rows.Scan(&bus.BusID, &bus.License, &bus.TotalSeats, &bus.EmptySeats, &bus.Departure, &bus.Destination,
			&bus.BeginAt, &bus.EndAt, &bus.Price, &bus.Info, &bus.Weekly, &bus.Status)
		util.Report(err)
		*b = append(*b, bus)
	}
	return nil
}

func Favorite(session *string, busID *string) error {
	var user User
	err := NowUser(&user, session)
	util.Report(err)
	if Favorited(&user.UserID, busID) || user.Account == "" {
		return errors.New("favorite: already favorite")
	}
	go func() {
		stmt, err := util.DB.Prepare("INSERT INTO favorites (user_id, bus_id) VALUES (?, ?)")
		util.Report(err)
		_, err = stmt.Query(user.UserID, busID)
		util.Report(err)
	}()
	return nil
}

func Unfavorite(session *string, busID *string) error {
	var user User
	err := NowUser(&user, session)
	util.Report(err)
	log.Print(*busID)
	if !Favorited(&user.UserID, busID) || user.Account == "" {
		return errors.New("favorite: no favorite")
	}
	go func() {
		stmt, err := util.DB.Prepare("DELETE FROM favorites WHERE user_id = ? AND bus_id = ?")
		util.Report(err)
		_, err = stmt.Query(user.UserID, busID)
		util.Report(err)
	}()
	return nil
}

func Favorited(userID *string, busID *string) bool {
	log.Print(*userID)
	stmt, err := util.DB.Prepare(`SELECT * FROM favorites WHERE user_id = ? AND bus_id = ?`)
	util.Report(err)
	rows, err := stmt.Query(*userID, *busID)
	util.Report(err)
	return rows.Next()
}
