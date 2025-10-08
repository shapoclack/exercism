package annalyn

func CanFastAttack(knightIsAwake bool) bool {
    return !knightIsAwake
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
    return knightIsAwake || archerIsAwake || prisonerIsAwake
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
    return !archerIsAwake && prisonerIsAwake
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
    // Заключенный не спит, рыцарь и лучник спят
    if !knightIsAwake && !archerIsAwake && prisonerIsAwake {
        return true
    }
    // Собака присутствует и лучник спит
    if petDogIsPresent && !archerIsAwake {
        return true
    }
    return false
}