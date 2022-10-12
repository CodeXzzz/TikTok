package model

import "time"

// TableVideo 数据库中存入的视频数据
type TableVideo struct {
	Id          int64     `json:"id"`
	AuthorId    int64     `json:"author_id"`
	PlayUrl     string    `json:"play_url"`
	CoverUrl    string    `json:"cover_url"`
	PublishTime time.Time `json:"publish_time"`
	Title       string    `json:"title"`
}

// Video 服务层中，给video添加点赞评论数等项
type Video struct {
	TableVideo
	Author        User  `json:"author"`
	FavoriteCount int64 `json:"favorite_count"`
	CommentCount  int64 `json:"comment_count"`
	IsFavorite    bool  `json:"is_favorite"`
}

// FeedResponse 视频feed流
type FeedResponse struct {
	VideoList []Video `json:"video_list"`
	//本次返回的视频中，发布最早的视频的时间，作为下次请求的latest_time
	NextTime int64 `json:"next_time"`
}

// VideoListResponse 发布的视频列表
type VideoListResponse struct {
	VideoList []Video `json:"video_list"`
}
