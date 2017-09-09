package mipush

import (
	"log"
	"testing"
	"time"
)

func TestMiPush(t *testing.T) {
	message := NewMessage().
		Description("notify description").
		Payload("this+is+xiaomi+push").
		Title("notify title").
		NotifyType(NotifyTypeSound).
		TimeToLive(time.Millisecond * 1).
		PassThrough(PassThroughNotify).
		NotifyID(0)

	AppID := ""
	AppKey := ""
	AppSecret := ""
	Init(AppID, AppKey, AppSecret)
	PackageName := ""
	xxx := ""
	log.Println(SendRegID(message, PackageName, xxx))
}
