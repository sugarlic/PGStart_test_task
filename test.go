package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	solders, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b_hp, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b_solders_prod, _ := strconv.Atoi(scanner.Text())
	b_solders := 0
	res := 0
	active_solders := 0
	s_damage := 0

	for {
		active_solders = solders
		// fmt.Println("-", solders-b_solders, b_solders-(solders-b_hp))
		if solders <= b_solders ||
			(solders > b_hp && float32(2*solders-b_solders-b_hp) >= float32(float32(b_solders-(solders-b_hp))/1.595)) {
			s_damage = b_hp
			b_hp -= active_solders
			active_solders -= s_damage
			if b_hp < 0 {
				b_hp = 0
			}
			if active_solders < 0 {
				active_solders = 0
			}

			b_solders -= active_solders
		} else {
			s_damage = b_solders
			b_solders -= active_solders
			active_solders -= s_damage
			if b_solders < 0 {
				b_solders = 0
			}
			if active_solders < 0 {
				active_solders = 0
			}

			b_hp -= active_solders
		}

		solders -= b_solders
		if b_hp > 0 {
			b_solders += b_solders_prod
		}
		res++
		if (b_hp <= 0 && b_solders <= 0) || solders <= 0 {
			break
		}
	}

	if b_hp <= 0 && b_solders <= 0 {
		fmt.Println(res)
	} else {
		fmt.Println(-1)
	}

}
