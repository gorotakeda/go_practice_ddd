package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name: "正常なユーザー",
			user: User{
				Name:  "テストユーザー",
				Email: "test@example.com",
			},
			wantErr: false,
		},
		{
			name: "名前が空",
			user: User{
				Name:  "",
				Email: "test@example.com",
			},
			wantErr: true,
		},
		{
			name: "メールアドレスが不正",
			user: User{
				Name:  "テストユーザー",
				Email: "invalid-email",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
} 