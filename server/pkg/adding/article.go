package adding

type article struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	TopicId     uint64 `json:"topicId"`
	Description string `json:"description"`
}