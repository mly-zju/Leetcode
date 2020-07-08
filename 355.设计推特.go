package main

/*
 * @lc app=leetcode.cn id=355 lang=golang
 *
 * [355] 设计推特
 *
 * https://leetcode-cn.com/problems/design-twitter/description/
 *
 * algorithms
 * Medium (34.16%)
 * Likes:    36
 * Dislikes: 0
 * Total Accepted:    1.7K
 * Total Submissions: 4.7K
 * Testcase Example:  '["Twitter","postTweet","getNewsFeed","follow","postTweet","getNewsFeed","unfollow","getNewsFeed"]\n' +
  '[[],[1,5],[1],[1,2],[2,6],[1],[1,2],[1]]'
 *
 *
 * 设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：
 *
 *
 * postTweet(userId, tweetId): 创建一条新的推文
 * getNewsFeed(userId):
 * 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
 * follow(followerId, followeeId): 关注一个用户
 * unfollow(followerId, followeeId): 取消关注一个用户
 *
 *
 * 示例:
 *
 *
 * Twitter twitter = new Twitter();
 *
 * // 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
 * twitter.postTweet(1, 5);
 *
 * // 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
 * twitter.getNewsFeed(1);
 *
 * // 用户1关注了用户2.
 * twitter.follow(1, 2);
 *
 * // 用户2发送了一个新推文 (推文id = 6).
 * twitter.postTweet(2, 6);
 *
 * // 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
 * // 推文id6应当在推文id5之前，因为它是在5之后发送的.
 * twitter.getNewsFeed(1);
 *
 * // 用户1取消关注了用户2.
 * twitter.unfollow(1, 2);
 *
 * // 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
 * // 因为用户1已经不再关注用户2.
 * twitter.getNewsFeed(1);
 *
 *
*/

// @lc code=start
type People struct {
	id          int
	followList  map[int]bool
	articleList []Article
}

type Article struct {
	peopleId   int
	id         int
	createTime int
}

type Mystack struct {
	list []Article
}

func (this *Mystack) ShiftUp(index int) {
	topIndex := (index - 1) / 2
	if topIndex >= 0 && this.list[index].createTime > this.list[topIndex].createTime {
		this.list[index], this.list[topIndex] = this.list[topIndex], this.list[index]
		this.ShiftUp(topIndex)
	}
}

func (this *Mystack) ShiftDown(index int) {
	leftIndex, rightIndex := index*2+1, index*2+2
	largeIndex := index
	length := len(this.list)
	if leftIndex < length && this.list[largeIndex].createTime < this.list[leftIndex].createTime {
		largeIndex = leftIndex
	}
	if rightIndex < length && this.list[largeIndex].createTime < this.list[rightIndex].createTime {
		largeIndex = rightIndex
	}
	if largeIndex != index {
		this.list[index], this.list[largeIndex] = this.list[largeIndex], this.list[index]
		this.ShiftDown(largeIndex)
	}
}

func (this *Mystack) queue(ele Article) {
	this.list = append(this.list, ele)
	newLen := len(this.list)
	this.ShiftUp(newLen - 1)
}

func (this *Mystack) dequeue() Article {
	length := len(this.list)
	res := this.list[0]
	this.list[0] = this.list[length-1]
	this.list = this.list[0 : length-1]
	this.ShiftDown(0)
	return res
}

func (this *Mystack) isEmpty() bool {
	return len(this.list) == 0
}

type Twitter struct {
	timeVal int
	people  map[int]*People
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		timeVal: 0,
		people:  map[int]*People{},
	}
}

func (this *Twitter) InitUser(userId int) {
	if this.people[userId] == nil {
		// 如果用户不存在，则创建
		this.people[userId] = &People{
			id:          userId,
			followList:  map[int]bool{},
			articleList: []Article{},
		}
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.InitUser(userId)
	// 向用户文章列表增加文章
	this.people[userId].articleList = append(this.people[userId].articleList, Article{
		peopleId:   userId,
		id:         tweetId,
		createTime: this.timeVal,
	})
	// 时间id自增
	this.timeVal++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	res := []int{}

	// 如果id不存在，直接返回空
	if this.people[userId] == nil {
		return res
	}

	// 获取关注列表, 把自己也加入
	peopleList := this.people[userId].followList
	peopleList[userId] = true

	// 根据关注列表获取每个人的文章游标, 并初始化优先队列
	cursorList := map[int]int{}
	stack := Mystack{}
	for peopleId, _ := range peopleList {
		people := this.people[peopleId]
		cursorList[peopleId] = len(people.articleList) - 1
		if cursorList[peopleId] >= 0 {
			stack.queue(people.articleList[cursorList[peopleId]])
		}
	}

	// 每一轮pop一个元素，然后相应被pop的用户游标减1然后queue一条新的，循环10轮
	for i := 0; i < 10; i++ {
		if stack.isEmpty() {
			// 如果队列空，直接跳出
			break
		}
		// 先pop
		newArticle := stack.dequeue()
		res = append(res, newArticle.id)
		// 更新对应用户游标
		cursorList[newArticle.peopleId]--
		if cursorList[newArticle.peopleId] >= 0 {
			// 如果依然大于0，新的一条推入
			stack.queue(this.people[newArticle.peopleId].articleList[cursorList[newArticle.peopleId]])
		}
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	this.InitUser(followerId)
	this.InitUser(followeeId)

	// 执行关注
	this.people[followerId].followList[followeeId] = true
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.people[followerId] == nil {
		return
	}
	if this.people[followerId].followList[followeeId] == true {
		delete(this.people[followerId].followList, followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// func main() {
// 	obj := Constructor()
// 	obj.PostTweet(1, 5)
// 	fmt.Println(obj.GetNewsFeed(1))
// 	obj.Follow(1, 2)
// 	obj.PostTweet(2, 6)
// 	fmt.Println(obj.GetNewsFeed(1))
// 	obj.Unfollow(1, 2)
// 	fmt.Println(obj.GetNewsFeed(1))
// }

// @lc code=end
