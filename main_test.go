package main

import (
	"encoding/json"
	"os"
	"reflect"
	"sync"
	"testing"
)

func Test_getEnv(t *testing.T) {
	type args struct {
		key      string
		fallback string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test", args{"TONGO_LOOP", "1"}, "1"},
		{"test", args{"TONGO_VALUE", "1"}, "test"},
	}
	os.Setenv("TONGO_VALUE", "test")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEnv(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("getEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_vote(t *testing.T) {
	type args struct {
		loop  int
		value int
		url   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vote(tt.args.loop, tt.args.value, tt.args.url)
		})
	}
}

func Test_hackTheVote(t *testing.T) {
	type args struct {
		presenterID string
		url         string
		votes       Votes
		wg          *sync.WaitGroup
		value       int
		id          int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

		// {"test", args{"", "", Votes{}, nil, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hackTheVote(tt.args.presenterID, tt.args.url, tt.args.votes, tt.args.wg, tt.args.value, tt.args.id)
		})
	}
}

func Test_getIdentifier(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIdentifier(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIdentifier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPresenterIdAndVotes(t *testing.T) {
	type args struct {
		text  string
		value int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   Votes
		wantErr bool
	}{
		// TODO: Add test cases.
		// {"vote 1",args{"1",5}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getPresenterIDAndVotes(tt.args.text, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPresenterIdAndVotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getPresenterIdAndVotes() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getPresenterIdAndVotes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getChoices(t *testing.T) {
	jsonState := `{"questions":[{"choices": [{"id":1}]}]}`
	var state State
	json.Unmarshal([]byte(jsonState), &state)
	votesExpected := Votes{
		"1": {2, 1},
	}
	type args struct {
		props State
		value int
	}
	tests := []struct {
		name string
		args args
		want Votes
	}{
		{"test", args{state, 2}, votesExpected},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getChoices(tt.args.props, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getChoices() = %v, want %v", got, tt.want)
			}
		})
	}
}
