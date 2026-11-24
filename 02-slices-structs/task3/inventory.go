package inventory

import (
	"cmp"
	"errors"
	"slices"
)

type Item struct {
	Name  string
	Price float64
	Qty   int
}

var (
	ErrNotFound     = errors.New("item not found")
	ErrInvalidInput = errors.New("invalid input value")
)

type Inventory struct {
	Items []Item
}

// Add modifies original Inventory, needs pointer reciever
func (i *Inventory) Add(item Item) {
	for j := range i.Items {
		if i.Items[j].Name == item.Name && i.Items[j].Price == item.Price {
			i.Items[j].Qty = i.Items[j].Qty + item.Qty
			return
		}
	}

	i.Items = append(i.Items, item)
}

// Remove modifies original Inventory, needs pointer reciever
func (i *Inventory) Remove(name string, qty int) error {
	if qty <= 0 {
		return ErrInvalidInput
	}

	for j := range i.Items {
		if i.Items[j].Name == name {
			if i.Items[j].Qty < qty {
				return ErrInvalidInput
			}

			i.Items[j].Qty = i.Items[j].Qty - qty

			if i.Items[j].Qty == 0 {
				i.Items = slices.Delete(i.Items, j, j+1)
			}

			return nil
		}
	}

	return ErrNotFound
}

// Total works on pointer reciever
func (i *Inventory) Total() float64 {
	var res float64

	for j := range i.Items {
		res += i.Items[j].Price * float64(i.Items[j].Qty)
	}

	return res
}

// SortedByPrice works on pointer reciever, does not modify underlying slice
func (i *Inventory) SortedByPrice() []Item {
	items := slices.Clone(i.Items)
	slices.SortFunc(items, func(a, b Item) int { return cmp.Compare(a.Price, b.Price) })
	return items
}

// SortedByTotalPrice works on pointer reciever, does not modify underlying slice
func (i *Inventory) SortedByTotalPrice() []Item {
	items := slices.Clone(i.Items)
	slices.SortFunc(items, func(a, b Item) int { return cmp.Compare(a.Price*float64(a.Qty), b.Price*float64(b.Qty)) })
	return items
}
