package datatypes

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

func (q *Question) countAnswers(idx []int) map[int]int {
	answers := q.getEmptyAnswersMap()
	for _, v := range idx {
		answers[q.data[v]] += 1
	}
	return answers
}