package main

import (
	"fmt"
	"testing"

	"github.com/google/wire"
)

/*
   @File: wire_test.go
   @Author: khaosles
   @Time: 2023/6/12 12:04
   @Desc:
*/

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{Name: "kitty"}
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

func InitMission(name string) Mission {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}
}

func TestWire(t *testing.T) {
	m := InitMission("aaa")
	m.Start()
}
