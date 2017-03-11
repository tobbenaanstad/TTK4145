package main

import (
	"./src/DriveElevator"
	"./src/Master"
	"./src/Network"
	"./src/Slave"
	"fmt"
	"time"
)

var MasterIP = "129.241.187.141"

var local_ip = ""

const BCAST_PORT = ":40018"
const PORT = ":30018"

func main() {
	fmt.Println("Start main...")

	chan_change_to_master := make(chan bool, 1)
	chan_change_to_slave := make(chan bool, 1)

	chan_kill := make(chan bool, 1)
	chan_kill2 := make(chan bool, 1)

	fmt.Println("In main...")
	state := 0

	//go DriveElevator.Run_elevator(chan_state, chan_dir, chan_floor, chan_order_executed, chan_new_hw_order, chan_new_master_order, chan_set_lights)
	local_ip, _ := Network.Udp_get_local_ip()
	chan_local_ip <- local_ip
	for {
		switch state {
		case 0:
			time.Sleep(5 * time.Second)
			Slave.Slave_init(BCAST_PORT, chan_change_to_master, chan_kill, chan_kill2)
			fmt.Println("Slave init done...")

		SlaveLoop:
			for {
				select {
				case change_state := <-chan_change_to_master:
					if change_state == true {
						time.Sleep(10 * time.Millisecond)

						//Må først gå i en sikker state. Dvs utfør alle interne ordre først. Ikke at i mot noen nye.
						state = 1
						fmt.Println("Killing slave")
						chan_kill <- true
						chan_kill2 <- true

						break SlaveLoop
					}
				}
			}
			time.Sleep(200 * time.Millisecond)
		case 1:
			Master.Master_init(chan_change_to_slave, PORT, chan_master)
			for {
				select {
				case <-time.After(60 * time.Second):
					state = 0
					// case change_state := <-chan_change_to_slave:
					// 	if change_state == true {
					// 		state = 1
					// 	}
				}
			}
			state = 2

		case 2:

		}
	}
}
