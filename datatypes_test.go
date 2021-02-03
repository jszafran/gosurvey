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
	t.Run("Test if empty answers counts are generated properly", func(t *testing.T) {
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
