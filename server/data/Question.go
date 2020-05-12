package data

type Question struct {
	Text     	string
	Answer   	string
	Asked    	int
	Answered 	int
	Category 	string
	Difficulty 	string
}

// NewQuestion returns a new Question
func NewQuestion(Text string, Answer string, Category string, Difficulty string) *Question {
	q := new(Question)
	q.Text = Text
	q.Category = Category
	q.Answer = Answer
	q.Asked = 0
	q.Answered = 0
	q.Difficulty = Difficulty
	return q
}

