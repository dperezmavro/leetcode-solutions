package solution

import (
	"container/heap"
	"container/list"
	"sync/atomic"
)

type Twitter struct {
	Users        map[int]*User
	TweetCounter atomic.Int32
}

type Tweet struct {
	UserID    int
	TweetID   int
	TweetUUID int32
	Index     int
}

type User struct {
	ID               int
	FollowedAccounts map[int]*User
	PersonalTweets   *list.List
}

func (u *User) AddFollowedAccount(f *User) {
	u.FollowedAccounts[f.ID] = f
}

func (u *User) AddTweet(p *Tweet) {
	u.PersonalTweets.PushFront(p)

	// append([]*Tweet{p}, u.PersonalTweets...)
}

func (u *User) RemoveFollower(f *User) {
	delete(u.FollowedAccounts, f.ID)
}

func Constructor() Twitter {
	users := make(map[int]*User)
	return Twitter{
		Users:        users,
		TweetCounter: atomic.Int32{},
	}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.TweetCounter.Add(1)
	u := t.getOrCreateUser(userId)
	p := &Tweet{
		UserID:    userId,
		TweetID:   tweetId,
		TweetUUID: t.TweetCounter.Load(),
	}
	u.AddTweet(p)
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	u := t.getOrCreateUser(userId)
	res := make([]int, 0, 10)

	pq := &PriorityQueue{}
	heap.Init(pq)

	for e := u.PersonalTweets.Front(); e != nil; e = e.Next() {
		heap.Push(pq, e.Value.(*Tweet))
	}

	for _, f := range u.FollowedAccounts {
		for e := f.PersonalTweets.Front(); e != nil; e = e.Next() {
			heap.Push(pq, e.Value.(*Tweet))
		}
	}

	for i := 0; i < 10 && pq.Len() > 0; i++ {
		t := heap.Pop(pq).(*Tweet)
		res = append(res, t.TweetID)
	}

	return res
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	u := t.getOrCreateUser(followeeId)
	u_2 := t.getOrCreateUser(followerId)
	u_2.AddFollowedAccount(u)
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	u := t.getOrCreateUser(followeeId)
	u_2 := t.getOrCreateUser(followerId)
	u_2.RemoveFollower(u)
}

func (t *Twitter) getOrCreateUser(userId int) *User {
	if u, ok := t.Users[userId]; ok {
		return u
	}
	u := &User{
		ID:               userId,
		FollowedAccounts: make(map[int]*User),
		PersonalTweets:   &list.List{},
	}

	t.Users[userId] = u
	return u
}

// PQ to find the freshest posts
type PriorityQueue []*Tweet

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].TweetUUID > pq[j].TweetUUID }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Tweet)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}
