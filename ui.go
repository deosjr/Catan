package main

import (
    "fmt"
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func dieroll() int {
    return rand.Intn(6)+1 + rand.Intn(6)+1
}

func loop(g game) {
    for _, player := range g.players {
        // opportunity to play cards before rolling
        // (you may play 1 card any time during your turn

        // rolling dice, get resources/discard + place robber
        roll := dieroll()
        fmt.Printf("%s rolled %d!\n", player.color, roll)
        res := g.resourceProduction(roll)
        for _, p := range g.players {
            if res[p.color].isEmpty() {
                continue
            }
            fmt.Printf("%s receives %s\n", p.color, res[p.color])
            p.hand.add(res[p.color])
        }

        // trading

        // building
    }
}
