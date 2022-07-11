package Training

import (
	"fmt"
	"math/rand"
)

type hero struct {
	hp      int
	atk     int
	def     int
	potions [5]int
	gold    int
	maxHP   int
}

type goblin struct {
	hp  int
	atk int
	def int
}

func generateHero(startingG int) hero {
	player := hero{}
	player.hp = rand.Intn(11) + 20
	player.atk = rand.Intn(2) + 1
	player.def = rand.Intn(4) + 1
	player.potions = [5]int{2, 2, 2, 2, 2}
	player.gold = startingG
	player.maxHP = player.hp

	return player
}

func generateGoblin() goblin {
	gob := goblin{}
	gob.hp = rand.Intn(6) + 5
	gob.atk = rand.Intn(2) + 2
	gob.def = rand.Intn(2) + 1

	return gob
}

func GoblinTower() {
	fmt.Println("hello")
	startingG := 0
	for true {
		player := generateHero(startingG)
		steps := 0
		lvl := 1
		slain := 0
		chance := 50
		roll := 0
		input := ""
		for player.hp > 0 {
			steps += 1
			roll = rand.Intn(100)
			if roll <= chance {
				gob := generateGoblin()
				for gob.hp > 0 && player.hp > 0 {
					fmt.Println("Your HP = ", player.hp)
					fmt.Println("Goblin HP = ", gob.hp)

					action := 0
					for action == 0 {
						fmt.Println("What will you do? 1 to attack. 2 to drink potion")
						fmt.Scanln(&action)
						switch action {
						case 1:
							action = 1
						case 2:
							action = 2
						default:
							fmt.Println("Invalid action")
							action = 0
						}
					}
					if action == 1 {
						fmt.Print("You attack! ")
						damage := player.atk - gob.def
						if damage <= 0 {
							damage = 1
						}
						gob.hp -= damage
						fmt.Println(damage, " damage!")
					} else if action == 2 {
						for ii := 0; ii < 5; ii++ {
							if player.potions[ii] > 0 {
								fmt.Println("You drink a potion! Recovered HP!")

								player.hp += player.potions[ii]
								player.potions[ii] -= 1
								if player.hp > player.maxHP {
									player.hp = player.maxHP
								}
								break
							}
						}

					}
					if gob.hp <= 0 {
						fmt.Println("You win! You get 2 gold")
						player.gold += 2
						slain += 1
						break
					}
					fmt.Print("Goblin attacks! ")
					damage := gob.atk - player.def
					if damage <= 0 {
						damage = 1
					}
					player.hp -= damage
					fmt.Println(damage, " damage!")
					if player.hp <= 0 {
						break
					}
				}
			}
			fmt.Println("You have cleared ", steps, " steps")
			if steps%10 == 0 {
				lvl += 1
				chance += 5
				fmt.Println("Level up!")
				if player.gold >= 4 {
					fmt.Println("Welcome to the potion shop! Want to buy a potion?")
					input = ""
					for input != "x" && player.gold >= 4 {
						fmt.Println("You have ", player.gold, " gold.")
						fmt.Println("Enter z to buy, x to leave")
						fmt.Scanln(&input)
						switch input {
						case "z":
							bought := false
							for ii := 0; ii < 5; ii++ {
								if player.potions[ii] < 2 {
									player.potions[ii] += 1
									player.gold -= 4
									bought = true
									break
								}
							}
							if !bought {
								fmt.Println("You can't carry any more potions!")
							}
						case "x":
							break
						default:
							fmt.Println("Sorry, I didn't understand you. Please try again")
						}
					}
					fmt.Println("Now leaving the potion shop. Come again!")
				}
			}
		}
		fmt.Println("You died!")
		input = ""
		fmt.Println("Do you want to keep playing? Enter z to stop, anything else to keep playing")
		fmt.Scanln(&input)
		if input == "z" {
			fmt.Println("You reached level", lvl)
			fmt.Println("Goblins slain: ", slain)
			break
		} else {
			startingG = player.gold
		}
	}
}
