package main

import "fmt"

type gasEngine struct{
	gallons uint8
	mpg uint8
}

type electricEngine struct{
	kwh uint8
	mph uint8
}

func (e gasEngine) milesLeft() uint8{
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mph
}

type engine interface{
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8){
	if miles <= e.milesLeft(){
		fmt.Println("You can make it there!")
	}else{
		fmt.Println("Need to fuel up first!")
	}
}
