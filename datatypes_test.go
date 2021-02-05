package datatypes

import "testing"
import "reflect"

func TestGetEmptyAnswersMap(t *testing.T) {
	qst := Question{
		data:   make([]int, 0),
		text:   "Whatever",
		minVal: 1,
		maxVal: 5,
	}
	t.Run("Test if empty answers counts are generated properly.", func(t *testing.T) {
		got := qst.getEmptyAnswersMap()
		want := map[int]int{
			1: 0,
			2: 0,
			3: 0,
			4: 0,
			5: 0,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v", got, want)
		}
	})
}

func TestCountAnswers(t *testing.T) {
	qst := Question{
		data:   []int{1, 3, 3, 1, 5, 1, 2, 5, 1},
		text:   "Whatever",
		minVal: 1,
		maxVal: 6,
	}
	t.Run("Test if empty index results in 0 counts.", func(t *testing.T) {
		idx := make([]int, 0)
		got := qst.countAnswers(idx)
		want := map[int]int{
			1: 0,
			2: 0,
			3: 0,
			4: 0,
			5: 0,
			6: 0,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v", got, want)
		}
	})
	t.Run("Test if answers are counted properly for full index.", func(t *testing.T) {
		idx := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		got := qst.countAnswers(idx)
		want := map[int]int{
			1: 4,
			2: 1,
			3: 2,
			4: 0,
			5: 2,
			6: 0,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v", got, want)
		}
	})
	t.Run("test if answer are counted properly for selective index.", func(t *testing.T) {
		idx := []int{2, 5, 8}
		got := qst.countAnswers(idx)
		want := map[int]int{
			1: 2,
			2: 0,
			3: 1,
			4: 0,
			5: 0,
			6: 0,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v", got, want)
		}
	})
}

func TestQuestionResults(t *testing.T) {
	t.Run("Test if respondents are summed properly.", func(t *testing.T) {
		qr := QuestionResults{1: 200, 2: 300, 3: 100, 4: 400}
		got := qr.getAllRespondentsCount()
		want := 1000

		if got != want {
			t.Errorf("Got %d, exptected %d, given input %v", got, want, qr)
		}
	})
}

func TestOrgNodes(t *testing.T) {
	t.Run("Test if rollup filtering works.", func(t *testing.T) {
		on := OrgNodes{"N01.", "N01.01.", "N01.01.01.", "N01.01.02.", "N02."}
		got := on.filterByOrgUnit("N01.01.", Rollup)
		want := []int{1, 2, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v, given input %v", got, want, on)
		}

		got = on.filterByOrgUnit("N01.03.", Rollup)
		want = make([]int, 0)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v, given input %v", got, want, on)
		}
	})
}