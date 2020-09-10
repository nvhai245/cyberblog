package commentModel

type Comment struct {
	ID        int32  `db:"id"`
	PostId    int32  `db:"post_id"`
	AuthorId  int32  `db:"author_id"`
	Content   string `db:"content"`
	UpVote    int32  `db:"up_vote"`
	CreatedAt int32  `db:"created_at"`
	UpdatedAt int32  `db:"updated_at"`
}
