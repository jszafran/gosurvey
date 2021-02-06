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
	on := OrgNodes{"N01.", "N01.01.", "N01.01.01.", "N01.01.02.", "N02."}
	t.Run("Test if rollup filtering works for existing unit.", func(t *testing.T) {
		node := "N01.01."
		got := on.filterByOrgUnit(node, Rollup)
		want := []int{1, 2, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v, given input %v and node %s", got, want, on, node)
		}
	})
	t.Run("Test if rollup filtering works for nonexisting unit.", func(t *testing.T) {
		node := "N01.03."
		got := on.filterByOrgUnit(node, Rollup)
		want := make([]int, 0)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v, given input %v", got, want, on)
		}
	})
	t.Run("Test if direct filtering works for existing unit.", func(t *testing.T) {
		node := "N01.01.01."
		got := on.filterByOrgUnit("N01.01.01.", Direct)
		want := []int{2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v, given input %v and org unit %s", got, want, on, node)
		}
	})
	t.Run("Test if direct filtering works for non-existing unit.", func(t *testing.T) {
		node := "N05.01.02."
		got := on.filterByOrgUnit(node, Direct)
		want := make([]int, 0)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v, given input %v and org unit %s", got, want, on, node)
		}
	})
}

// tests for possible alternative implemenation of org unit filtering with 0/1
func TestOrgNodesZO(t *testing.T) {
	on := OrgNodes{"N01.", "N01.01.", "N01.01.01.", "N01.01.02.", "N02."}
	t.Run("Test if rollup filtering works for existing unit.", func(t *testing.T) {
		node := "N01.01."
		got := on.filterByOrgUnitZO(node, Rollup)
		want := []int8{0, 1, 1, 1, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v, given input %v and node %s", got, want, on, node)
		}
	})
	t.Run("Test if rollup filtering works for nonexisting unit.", func(t *testing.T) {
		node := "N01.03."
		got := on.filterByOrgUnitZO(node, Rollup)
		want := []int8{0, 0, 0, 0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v, given input %v", got, want, on)
		}
	})
	t.Run("Test if direct filtering works for existing unit.", func(t *testing.T) {
		node := "N01.01.01."
		got := on.filterByOrgUnitZO("N01.01.01.", Direct)
		want := []int8{0, 0, 1, 0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v, given input %v and org unit %s", got, want, on, node)
		}
	})
	t.Run("Test if direct filtering works for non-existing unit.", func(t *testing.T) {
		node := "N05.01.02."
		got := on.filterByOrgUnitZO(node, Direct)
		want := []int8{0, 0, 0, 0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v, given input %v and org unit %s", got, want, on, node)
		}
	})
}
