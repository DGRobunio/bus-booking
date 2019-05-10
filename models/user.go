package models

import (
	"bus-booking/util"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"math/rand"
	"strconv"
)

type User struct {
	UserID   string
	Account  string
	Password string
	Salt     string
	Balance  float64
	IsAdmin  bool
}

func NowUser(u *User, session *string) error {
	*session = "session:" + *session
	result := util.Redis.HMGet(*session, "userId", "account", "balance", "isAdmin")
	if result.Err() != nil {
		return errors.New("user info: error")
	} else if result.Val()[0] == nil {
		return nil
	} else {
		u.UserID = result.Val()[0].(string)
		u.Account = result.Val()[1].(string)
		u.Balance, _ = strconv.ParseFloat(result.Val()[2].(string), 64)
		u.IsAdmin, _ = strconv.ParseBool(result.Val()[3].(string))
		return nil
	}
}

func Login(u *User) (string, error) {
	var session string
	c := make(chan bool, 2)
	go func(session *string, c chan bool) {
		const charSet = "abcdefghijklmnopqrstuvwxyz0123456789"
		b := make([]byte, 26)
		for i := range b {
			b[i] = charSet[rand.Intn(len(charSet))]
		}
		*session = string(b)
		c <- true
	}(&session, c)
	go checkUserPassword(u, c)
	for i := 0; i < cap(c); i++ {
		if !<-c {
			return "", errors.New("login: error")
		}
	}
	go updateUserCache(u, &session)
	return session, nil
}

func Logout(session *string) {
	go func() {
		util.Redis.Del("session:" + *session)
	}()
}

func SignUp(u *User) error {
	var salt string
	c := make(chan bool, 2)
	go func(salt *string, c chan bool) {
		const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJLKMNOPQRSTUVWXYZ0123456789"
		b := make([]byte, 8)
		for i := range b {
			b[i] = charSet[rand.Intn(len(charSet))]
		}
		*salt = string(b)
		c <- false
	}(&salt, c)
	go checkUserExists(u, c)
	for i := 0; i < cap(c); i++ {
		if <-c {
			return errors.New("sign up: error")
		}
	}
	go func() {
		stmt, err := util.DB.Prepare("INSERT INTO users (account, password, salt) VALUES (?, ?, ?)")
		util.Report(err)
		password := sha1.New()
		password.Write([]byte(u.Password + salt))
		_, err = stmt.Query(u.Account, hex.EncodeToString(password.Sum(nil)), salt)
		util.Report(err)
	}()
	return nil
}

func UpdateUser(u *User, session *string, newPassword *string) error {
	err := NowUser(u, session)
	if err != nil || u.Account == "" {
		return errors.New("update user: error")
	}
	c := make(chan bool, 2)
	go checkUserExists(u, c)
	go checkUserPassword(u, c)
	for i := 0; i < cap(c); i++ {
		if !<-c {
			return errors.New("update user: error")
		}
	}
	go func(u *User) {
		const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJLKMNOPQRSTUVWXYZ0123456789"
		b := make([]byte, 8)
		for i := range b {
			b[i] = charSet[rand.Intn(len(charSet))]
		}
		salt := string(b)
		stmt, err := util.DB.Prepare("UPDATE users SET password = ?, salt = ? WHERE account = ?")
		util.Report(err)
		password := sha1.New()
		password.Write([]byte(*newPassword + salt))
		_, err = stmt.Query(hex.EncodeToString(password.Sum(nil)), salt, u.Account)
		util.Report(err)
	}(u)
	return nil
}

func check(u *User) bool {
	if u.Account == "" {
		return false
	}
	stmt, err := util.DB.Prepare("SELECT * FROM users WHERE account = ?")
	util.Report(err)
	rows, err := stmt.Query(u.Account)
	util.Report(err)
	if  !rows.Next() {
		return false
	}
	var user User
	err = rows.Scan(&user.UserID, &user.Account, &user.Password, &user.Salt, &user.Balance,
		&user.IsAdmin)
	util.Report(err)
	*u = user
	return true
}

func checkUserExists(u *User, c chan bool) {
	c <- check(u)
}

func checkUserPassword(u *User, c chan bool) {
	if u.Account == "" || u.Password == "" {
		c <- false
	} else {
		stmt, err := util.DB.Prepare("SELECT * FROM users WHERE account = ?")
		util.Report(err)
		var user User
		rows, err := stmt.Query(u.Account)
		util.Report(err)
		for rows.Next() {
			err := rows.Scan(&user.UserID, &user.Account, &user.Password, &user.Salt, &user.Balance,
				&user.IsAdmin)
			util.Report(err)
		}
		password := sha1.New()
		password.Write([]byte(u.Password + user.Salt))
		if user.Password == hex.EncodeToString(password.Sum(nil)) {
			*u = user
			c <- true
		} else {
			c <- false
		}
	}
}

func updateUserCache(u *User, session *string) {
	_, err := util.Redis.HMSet("session:"+*session, map[string]interface{}{
		"userId":  u.UserID,
		"account": u.Account,
		"balance": u.Balance,
		"isAdmin": u.IsAdmin,
	}).Result()
	util.Report(err)
}
