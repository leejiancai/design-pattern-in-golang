package bridge

import "fmt"

// 演示一个发送根据事件级别发送通知的功能。本来整个逻辑可以合在一起实现，但是为了后续迭代各自演进，可以把发送通知和不同事件级别处理的类分开实现。
type MsgSender interface {
	send(msg string) string
}

type Notification interface {
	notify(msg string) string
}

type TeleMsgSender struct{}

func (t *TeleMsgSender) send(msg string) string {
	return fmt.Sprintf("TeleMsgSender:%s", msg)
}

type WechatSender struct{}

func (w *WechatSender) send(msg string) string {
	return fmt.Sprintf("WechatSender:%s", msg)
}

// 通过组合的方式是实现
type NormalNotification struct {
	sender MsgSender
}

func (n *NormalNotification) notify(msg string) string {
	return n.sender.send(msg)
}

type UrgencyNotification struct {
	sender MsgSender
}

func (n *UrgencyNotification) notify(msg string) string {
	return n.sender.send(msg)
}
