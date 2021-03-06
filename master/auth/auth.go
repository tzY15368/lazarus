package auth

import (
	"errors"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tzY15368/lazarus/gen-go/RPCService"
	"github.com/tzY15368/lazarus/master/models"
	"gorm.io/gorm"
)

var mapDidChange bool
var mapMutex sync.Mutex

func ChangeAuthMap() {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	mapDidChange = true
}

func AuthMapDidChange() bool {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	if mapDidChange {
		mapDidChange = false
		return true
	} else {
		mapDidChange = false
		return false
	}
}

func TokenIsValid(token string) bool {
	var user models.User
	models.DB.Where("token = ?", token).Take(&user)
	return time.Now().Before(user.ExpireAt)
}

func GetUserMap() RPCService.UserData {
	var entries []models.User
	userMap := make(map[string]struct{})
	userdata := RPCService.UserData{}
	err := models.DB.Where("expire_at >= ?", time.Now()).Find(&entries).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Error("auth error:", err.Error())
	} else {
		logrus.Infof("auth: got %d entries", len(entries))
	}
	for _, user := range entries {
		userMap[user.Token] = struct{}{}
		userdata = append(userdata, user.Token)
	}
	return userdata
}
