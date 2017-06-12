package day21

type ShopShelf map[string]Item
type Shop struct {
    Inventory []ShopShelf
}

func NewShop() *Shop {
    self := Shop{ Inventory: []ShopShelf{} }
    
    return &self
}

func (self *Shop) AddItem( item Item ) {
    self.extendShelves( item.Type )
    
    self.Inventory[item.Type][item.Name] = item
}

func (self *Shop) BuyItem( itemType ItemType, name string ) Item {
    item := self.Inventory[itemType][name]
    delete(self.Inventory[itemType], name)
    return item
}

func (self *Shop) ListItems( itemType ItemType ) ShopShelf {
    self.extendShelves( itemType )
    
    return self.Inventory[itemType]
}

func (self *Shop) extendShelves( itemType ItemType ) {
    // Make sure we have a slot for this item type
    for len(self.Inventory) <= int(itemType) {
        self.Inventory = append( self.Inventory, make( ShopShelf) )
    }
}