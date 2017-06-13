package day21

type ShopShelf []Item
type Shop struct {
    Inventory []ShopShelf
}

func NewShop() *Shop {
    self := Shop{ Inventory: []ShopShelf{} }
    
    return &self
}

func (self *Shop) AddItem( item Item ) {
    self.extendShelves( item.Type )
    
    self.Inventory[item.Type] = append(
        self.Inventory[item.Type], item,
    )
}

func (self *Shop) GetShelf( itemType ItemType ) ShopShelf {
    self.extendShelves( itemType )
    
    return self.Inventory[itemType]
}

func (self *Shop) extendShelves( itemType ItemType ) {
    // Make sure we have a slot for this item type
    for len(self.Inventory) <= int(itemType) {
        self.Inventory = append( self.Inventory, ShopShelf{} )
    }
}
