package inventory

import (
	"errors"
	"slices"
)

type Item struct {
	Name  string
	Price float64
	Qty   int
}

var ErrNotFound = errors.New("Item not found")

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
		return errors.New("Invalid value for quantity, cannot be less or equal to zero")
	}

	for j := range i.Items {
		if i.Items[j].Name == name {
			if i.Items[j].Qty < qty {
				return errors.New("Invalid value for quantity, cannot be more than original value")
			}

			if i.Items[j].Qty == qty {
				slices.Delete(i.Items, j, j+1)
				return nil
			}

			i.Items[j].Qty = i.Items[j].Qty - qty
			return nil
		}
	}

	return ErrNotFound
}

// Total can work on a copy, does not modify, works with value reciever
func (i Inventory) Total() float64 {
	var res float64

	for j := range i.Items {
		res += i.Items[j].Price * float64(i.Items[j].Qty)
	}

	return res
}

// SortedByPrice can work on a copy, does not modify, works with value reciever
func (i Inventory) SortedByPrice() []Item {
	slices.SortFunc(i.Items, func(a, b Item) int { return int(a.Price*100) - int(b.Price*100) })
	return i.Items
}

// SortedByTotalPrice can work on a copy, does not modify, works with value reciever
func (i Inventory) SortedByTotalPrice() []Item {
	slices.SortFunc(i.Items, func(a, b Item) int { return int(a.Price*100)*a.Qty - int(b.Price*100)*a.Qty })
	return i.Items
}
