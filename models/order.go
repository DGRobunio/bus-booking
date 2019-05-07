package models

import (
	"bus-booking/util"
	"errors"
)

type Order struct {
	OrderID string `json:"orderID"`
	Bus     Bus    `json:"bus"`
	OrderAt string `json:"orderAt"`
	Status  uint8  `json:"status"`
}

func AllOrders(session *string, o *[] Order) error {
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
			&order.Bus.BusID, &order.Bus.License, &order.Bus.TotalSeats, &order.Bus.EmptySeats, &order.Bus.Departure,
			&order.Bus.Destination, &order.Bus.BeginAt, &order.Bus.EndAt, &order.Bus.Price, &order.Bus.Info,
			&order.Bus.Weekly, &order.Bus.Status)
		util.Report(err)
		*o = append(*o, order)
	}
	return nil
}

func OneOrder(session *string, o *Order) error {
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
