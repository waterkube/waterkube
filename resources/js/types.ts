export interface Player {
    money: number
    boat_count: number
    diver_count: number
    submarine_count: number
}

export interface Grid {
    name: string
    type: string
    status: string
}

export interface Artifact {
    name: string
    price: number
    quantity: number
    type: string
}
