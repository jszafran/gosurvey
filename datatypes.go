package datatypes

type QuestionResponses []int

func (qr *QuestionResponses) countAnswers(idx []int, q Question) map[int]int {
	answers := q.getEmptyAnswersMap()
	for _, v := range idx {
		answers[q.data[v]] += 1
	}
	return answers
}

type Question struct {
	data   []int
	text   string
	minVal int
	maxVal int
}

func (q *Question) getEmptyAnswersMap() map[int]int {
	m := make(map[int]int, 0)
	for i := q.minVal; i <= q.maxVal; i++ {
		m[i] = 0
	}
	return m
}
