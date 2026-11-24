package inventory

import (
	"slices"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := map[string]struct {
		initial Inventory
		added   Item
		want    Inventory
	}{
		"empty initial": {
			initial: Inventory{},
			added:   Item{"apple", 1.0, 5},
			want:    Inventory{[]Item{{"apple", 1.0, 5}}},
		},
		"adding same element": {
			initial: Inventory{[]Item{{"apple", 1.0, 5}}},
			added:   Item{"apple", 1.0, 5},
			want:    Inventory{[]Item{{"apple", 1.0, 10}}},
		},
		"adding same but different price": {
			initial: Inventory{[]Item{{"apple", 1.0, 5}}},
			added:   Item{"apple", 1.5, 5},
			want:    Inventory{[]Item{{"apple", 1.0, 5}, {"apple", 1.5, 5}}},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			tc.initial.Add(tc.added)
			if !slices.Equal(tc.initial.Items, tc.want.Items) {
				t.Errorf("Add(%v) resulted in %v; want %v", tc.added, tc.initial, tc.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	cases := map[string]struct {
		initial Inventory
		removed Item
		wantErr string
		want    Inventory
	}{
		"empty initial": {
			initial: Inventory{},
			removed: Item{"apple", 1.0, 5},
			wantErr: "item not found",
			want:    Inventory{},
		},
		"removing same element": {
			initial: Inventory{[]Item{{"apple", 1.0, 5}}},
			removed: Item{"apple", 1.0, 5},
			want:    Inventory{[]Item{}},
		},
		"removing non existent item": {
			initial: Inventory{[]Item{{"apple", 1.0, 5}}},
			removed: Item{"banana", 1.7, 6},
			wantErr: "item not found",
			want:    Inventory{[]Item{{"apple", 1.0, 5}}},
		},
		"removing more than exists": {
			initial: Inventory{[]Item{{"apple", 1.0, 5}}},
			removed: Item{"apple", 1.0, 6},
			wantErr: "Invalid value for quantity, cannot be more than original value",
			want:    Inventory{[]Item{{"apple", 1.0, 5}}},
		},
		"removing less than zero": {
			initial: Inventory{[]Item{{"apple", 1.0, 5}}},
			removed: Item{"apple", 1.0, -5},
			wantErr: "Invalid value for quantity, cannot be less or equal to zero",
			want:    Inventory{[]Item{{"apple", 1.0, 5}}},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := tc.initial.Remove(tc.removed.Name, tc.removed.Qty)
			if err != nil && err.Error() != tc.wantErr {
				t.Errorf("Remove(%v) resulted in %v; want %v", err.Error(), tc.initial, tc.wantErr)
			}
			if !slices.Equal(tc.initial.Items, tc.want.Items) {
				t.Errorf("Remove(%v) resulted in %v; want %v", tc.removed, tc.initial, tc.want)
			}
		})
	}
}

func TestTotal(t *testing.T) {
	cases := map[string]struct {
		initial Inventory
		want    float64
	}{
		"empty initial": {
			initial: Inventory{},
			want:    0,
		},
		"normal": {
			initial: Inventory{[]Item{{"apple", 1.0, 5}, {"banana", 1.7, 6}}},
			want:    15.2,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := tc.initial.Total()
			if got != tc.want {
				t.Errorf("Total(%v) resulted in %v; want %v", tc.initial, got, tc.want)
			}
		})
	}
}

func TestSortedByPrice(t *testing.T) {
	cases := map[string]struct {
		initial Inventory
		want    []Item
	}{
		"empty": {
			initial: Inventory{},
			want:    []Item{},
		},
		"one element": {
			initial: Inventory{[]Item{{"banana", 1.5, 4}}},
			want:    []Item{{"banana", 1.5, 4}},
		},
		"normal": {
			initial: Inventory{[]Item{{"banana", 1.5, 4}, {"apple", 1.0, 5}}},
			want:    []Item{{"apple", 1.0, 5}, {"banana", 1.5, 4}},
		},
		"same price": {
			initial: Inventory{[]Item{{"apple", 1.0, 5}, {"banana", 1.0, 5}}},
			want:    []Item{{"apple", 1.0, 5}, {"banana", 1.0, 5}},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := tc.initial.SortedByPrice()
			if !slices.Equal(got, tc.want) {
				t.Errorf("Sort by price(%v) resulted in %v; want %v", tc.initial, got, tc.want)
			}
		})
	}
}

func TestSortedByTotalPrice(t *testing.T) {
	cases := map[string]struct {
		initial Inventory
		want    []Item
	}{
		"empty": {
			initial: Inventory{},
			want:    []Item{},
		},
		"one element": {
			initial: Inventory{[]Item{{"banana", 1.5, 4}}},
			want:    []Item{{"banana", 1.5, 4}},
		},
		"normal": {
			initial: Inventory{[]Item{{"banana", 1.0, 7}, {"apple", 1.0, 5}}},
			want:    []Item{{"apple", 1.0, 5}, {"banana", 1.0, 7}},
		},
		"same price": {
			initial: Inventory{[]Item{{"apple", 1.0, 5}, {"banana", 1.0, 5}}},
			want:    []Item{{"apple", 1.0, 5}, {"banana", 1.0, 5}},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := tc.initial.SortedByTotalPrice()
			if !slices.Equal(got, tc.want) {
				t.Errorf("Sort by price(%v) resulted in %v; want %v", tc.initial, got, tc.want)
			}
		})
	}
}
