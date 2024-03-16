package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getUser(t *testing.T) {
	t.Parallel()
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want *user
	}{
		{
			name: "success",
			args: args{id: 1},
			want: &user{id: 1, name: "user1"},
		},
		{
			name: "not found",
			args: args{id: 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, _ := getUser(tt.args.id)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_listPostsByUserIDs(t *testing.T) {
	t.Parallel()
	type args struct {
		userIDs []int
	}
	tests := []struct {
		name string
		args args
		want []*post
	}{
		{
			name: "success",
			args: args{userIDs: []int{1, 2}},
			want: []*post{
				{id: 1, userID: 1, content: "post1", createdAt: 1},
				{id: 2, userID: 2, content: "post2", createdAt: 2},
				{id: 3, userID: 1, content: "post3", createdAt: 3},
				{id: 4, userID: 1, content: "post4", createdAt: 4},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := listPostsByUserIDs(tt.args.userIDs)
			require.Equal(t, tt.want, got)
		})
	}
}
