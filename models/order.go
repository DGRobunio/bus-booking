package models

import (
	"bus-booking/util"
	"errors"
	"strings"
)

type Order struct {
	OrderID string `json:"orderID"`
	Bus     Bus    `json:"bus"`
	OrderAt string `json:"orderAt"`
	Status  uint8  `json:"status"`
}

func AllOrders(o *[] Order, session *string) error {
	var user User
	err := NowUser(&user, session)
	util.Report(err)
	stmt, err := util.DB.Prepare(`SELECT orders.id, orders.order_at, orders.status, buses.* FROM orders
		JOIN buses on orders.bus_id = buses.id WHERE user_id = ?`)
	util.Report(err)
	var order Order
	rows, err := stmt.Query(user.UserID)
	util.Report(err)
	for rows.Next() {
		err := rows.Scan(&order.OrderID, &order.OrderAt, &order.Status,
			&order.Bus.BusID, &order.Bus.License, &order.Bus.TotalSeats, &order.Bus.EmptySeats,
			&order.Bus.Departure, &order.Bus.Destination, &order.Bus.BeginAt, &order.Bus.EndAt,
			&order.Bus.Price, &order.Bus.Info, &order.Bus.Weekly, &order.Bus.Status)
		util.Report(err)
		*o = append(*o, order)
	}
	return nil
}

func OneOrder(o *Order, session *string) error {
	var user User
	err := NowUser(&user, session)
	util.Report(err)
	stmt, err := util.DB.Prepare(`SELECT orders.id, orders.order_at, orders.status, buses.* FROM orders
		JOIN buses on orders.bus_id = buses.id WHERE user_id = ? AND orders.id = ?`)
	util.Report(err)
	rows, err := stmt.Query(user.UserID, o.OrderID)
	util.Report(err)
	if !rows.Next() {
		return errors.New("bus: error")
	}
	err = rows.Scan(&o.OrderID, &o.OrderAt, &o.Status, &o.Bus.BusID, &o.Bus.License, &o.Bus.TotalSeats,
		&o.Bus.EmptySeats, &o.Bus.Departure, &o.Bus.Destination, &o.Bus.BeginAt, &o.Bus.EndAt, &o.Bus.Price,
		&o.Bus.Info, &o.Bus.Weekly, &o.Bus.Status)
	util.Report(err)
	return nil
}

func Book(session *string, busID *string) error {
	var user User
	bus := Bus{BusID: *busID}
	c := make(chan bool, 2)
	go func(c chan bool) {
		err := NowUser(&user, session)
		c <- check(&user)
		util.Report(err)
	}(c)
	go func(c chan bool) {
		err := OneBus(&bus)
		util.Report(err)
		c <- true
	}(c)
	for i := 0; i < cap(c); i++ {
		<-c
	}
	if user.Balance < bus.Price {
		return errors.New("book: not enough money")
	} else if bus.EmptySeats < 1 {
		return errors.New("book: not enough seats")
	}
	go func(order chan bool) {
		tx, err := util.DB.Begin()
		util.Report(err)
		insert, _ := tx.Prepare("INSERT INTO orders(status, user_id, bus_id) VALUES (?, ?, ?)")
		updateUser, _ := tx.Prepare(`UPDATE users SET users.balance = ? WHERE users.id = ?`)
		updateBus, _ := tx.Prepare(`UPDATE buses SET buses.empty_seats = ? WHERE buses.id = ?`)
		_, _ = insert.Query(0, user.UserID, bus.BusID)
		_, _ = updateUser.Query(user.Balance-bus.Price, user.UserID)
		_, _ = updateBus.Query(bus.EmptySeats-1, bus.BusID)
		err = tx.Commit()
		_ = insert.Close()
		_ = updateUser.Close()
		_ = updateBus.Close()
		if err != nil {
			util.Report(err)
			order <- false
			_ = tx.Rollback()
		} else {
			order <- true
		}
	}(util.Order)
	if <-util.Order {
		if check(&user) {
			str := strings.Replace(*session, "session:", "", -1)
			updateUserCache(&user, &str)
			return nil
		} else {
			return errors.New("book: failed")
		}
	} else {
		return errors.New("book: failed")
	}
}
