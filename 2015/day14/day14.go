package day14

type Reindeer struct {
    Speed int
    Duration int
    Rest int
}

func ( r *Reindeer ) RunRunReindeer( time int ) (dist int) {
    for time > r.Duration {
        dist += r.Speed * r.Duration
        time -= r.Duration
        time -= r.Rest
    }
    
    if time > 0 {
        dist += r.Speed * time
    }
    
    return
}
