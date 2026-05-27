package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwitter(t *testing.T) {
	tests := []struct {
		name       string
		assertFunc func(t *testing.T)
	}{
		{
			name: "base-example",
			assertFunc: func(t *testing.T) {
				t.Helper()
				twitter := Constructor()
				twitter.PostTweet(1, 5)

				tweets := twitter.GetNewsFeed(1)
				tweetsWant := []int{5}
				if !assert.Equal(t, tweetsWant, tweets) {
					t.Errorf("got wrong tweets: %+v - %+v", tweetsWant, tweets)
				}
				twitter.Follow(1, 2)
				twitter.PostTweet(2, 6)
				tweets = twitter.GetNewsFeed(1)
				tweetsWant = []int{6, 5}
				if !assert.Equal(t, tweetsWant, tweets) {
					t.Errorf("got wrong tweets after subscription: %+v - %+v", tweetsWant, tweets)
				}
				twitter.Unfollow(1, 2)          // User 1 unfollows user 2.
				tweets = twitter.GetNewsFeed(1) // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.
				tweetsWant = []int{5}
				if !assert.Equal(t, tweetsWant, tweets) {
					t.Errorf("got wrong tweets after unsubscribe: %+v - %+v", tweetsWant, tweets)
				}
			},
		},
		{
			name: "test-1",
			assertFunc: func(t *testing.T) {
				t.Helper()
				twitter := Constructor()
				twitter.PostTweet(1, 1)

				tweets := twitter.GetNewsFeed(1)
				tweetsWant := []int{1}
				if !assert.Equal(t, tweetsWant, tweets) {
					t.Errorf("got wrong tweets: %+v - %+v", tweetsWant, tweets)
				}
				twitter.Follow(2, 1)
				tweets = twitter.GetNewsFeed(2)
				tweetsWant = []int{1}
				if !assert.Equal(t, tweetsWant, tweets) {
					t.Errorf("got wrong tweets after subscription: %+v - %+v", tweetsWant, tweets)
				}
				twitter.Unfollow(2, 1)
				tweets = twitter.GetNewsFeed(2)
				tweetsWant = []int{}
				if !assert.Equal(t, tweetsWant, tweets) {
					t.Errorf("got wrong tweets after unsubscribe: %+v - %+v", tweetsWant, tweets)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertFunc(t)
		})
	}
}
