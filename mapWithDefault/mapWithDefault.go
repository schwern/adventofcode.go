package mapWithDefault

type Map struct {
    m map[string]interface{}
    d interface{}
}

func New( d interface{} ) *Map {
    self := Map{ m: make( map[string]interface{} ), d: d }
    return &self
}

func (self *Map) Get( key string ) interface{} {
    if val,ok := self.m[key]; ok {
        return val
    } else {
        return self.d
    }
}

func (self *Map) Set( key string, val interface{} ) {
    self.m[key] = val
}

func (self *Map) Delete( key string ) {
    delete(self.m, key)
}

func (self *Map) Exists( key string ) bool {
    _,ok := self.m[key]
    return ok
}
