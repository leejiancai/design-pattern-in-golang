package bridge

import "testing"

func TestNormalNotification_notify(t *testing.T) {
	type fields struct {
		sender MsgSender
	}
	type args struct {
		msg string
	}
	tel := TeleMsgSender{}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{name: "first", fields: fields{sender: &tel}, args: args{"Hello"}, want: "TeleMsgSender:Hello"},
		{name: "second", fields: fields{sender: &tel}, args: args{"World"}, want: "TeleMsgSender:World"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NormalNotification{
				sender: tt.fields.sender,
			}
			if got := n.notify(tt.args.msg); got != tt.want {
				t.Errorf("notify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrgencyNotification_notify(t *testing.T) {
	type fields struct {
		sender MsgSender
	}
	type args struct {
		msg string
	}

	wc := WechatSender{}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{name: "first", fields: fields{sender: &wc}, args: args{"Hello"}, want: "WechatSender:Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &UrgencyNotification{
				sender: tt.fields.sender,
			}
			if got := n.notify(tt.args.msg); got != tt.want {
				t.Errorf("notify() = %v, want %v", got, tt.want)
			}
		})
	}
}
