package dingding

import (
	"reflect"
	"testing"
)

func TestClient_Send(t *testing.T) {
	type fields struct {
		AccessToken string
		Secret      string
	}
	type args struct {
		message Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				AccessToken: tt.fields.AccessToken,
				Secret:      tt.fields.Secret,
			}
			got, err := c.Send(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Send() got = %v, want %v", got, tt.want)
			}
		})
	}
}
