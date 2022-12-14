package model

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	TotalFavorite int64  `json:"total_favorite,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
}
