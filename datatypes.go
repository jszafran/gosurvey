package datatypes

import "strings"

type OrgFilterType int

const (
	Direct OrgFilterType = iota
	Rollup
)

type Question struct {
	data   []int
	text   string
	minVal int
	maxVal int
}

type QuestionResults map[int]int

type OrgNodes []string

func (qr *QuestionResults) getAllRespondentsCount() int {
	res := 0
	for _, v := range *qr {
		res += v
	}
	return res
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

func (on *OrgNodes) filterByOrgUnit(orgNode string, filterType OrgFilterType) []int {
	res := make([]int, 0)
	for ix, v := range *on {
		if filterType == Rollup {
			if strings.HasPrefix(v, orgNode) {
				res = append(res, ix)
			}
		} else if filterType == Direct {
			if v == orgNode {
				res = append(res, ix)
			}
		}
	}
	return res
}
