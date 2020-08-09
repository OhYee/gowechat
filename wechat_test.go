package wechat

import "testing"

var wc = Wechat{
	Token:  "ohyee_token",
	AESKey: "QTTdEebUiSH4FnPSH3OhY1ePjVmYv9UoAWnTZKvWg5Q",
}

func TestWechat_CheckSignature(t *testing.T) {
	type args struct {
		signature string
		timestamp string
		nonce     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test check permission",
			args: args{
				signature: "6c19d0b3232e40dd2507345529ea8403d4628529",
				timestamp: "1596978401",
				nonce:     "1614485621",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wc.CheckSignature(tt.args.signature, tt.args.timestamp, tt.args.nonce); got != tt.want {
				t.Errorf("Wechat.CheckSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}
