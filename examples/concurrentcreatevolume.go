package main

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/docker/docker/pkg/devicemapper"
)

type mutexLock struct {
	mutex_bronze *sync.Mutex
	mutex_silver *sync.Mutex
	mutex_gold   *sync.Mutex
}

const (
	Gold_Device = "/dev/mapper/gold_vg-convoy_Linear_gold_data"
	Silver_Device = "/dev/mapper/silver_vg-convoy_Linear_silver_data"
	Iron_Device = "/dev/mapper/iron_vg-convoy_Stripe_iron_data"
)


func main() {
	var mutex = mutexLock{
		mutex_bronze: &sync.Mutex{},
		mutex_silver: &sync.Mutex{},
		mutex_gold:   &sync.Mutex{},
	}

	var wait = make(chan int, 60)

	for i := 1; i <= 20; i++ {
		go mutex.createBronzeVolume(fmt.Sprintf("%d", i), wait)
		go mutex.createSilverVolume(fmt.Sprintf("%d", i), wait)
		go mutex.createGoldVolume(fmt.Sprintf("%d", i), wait)
	}

	for i := 1; i <= 60; i++ {
		if 1 == <-wait {
			fmt.Println(i)
		}
	}

	for i := 1; i <= 20; i++ {
		go mutex.deleteBronzeVolume(i, wait)
		go mutex.deleteSilverVolume(i, wait)
		go mutex.deleteGoldVolume(i, wait)
	}

	for i := 1; i <= 60; i++ {
		if 1 == <-wait {
			fmt.Println(i)
		}
	}
}

func (mutex *mutexLock) createBronzeVolume(id string, stop chan int) {
	mutex.mutex_bronze.Lock()
	defer mutex.mutex_bronze.Unlock()
	err := exec.Command("sh", "-c", "dmsetup message /dev/mapper/iron_vg-convoy_Stripe_iron_data 0 "+fmt.Sprintf("'create_thin %s'", id)).Run()
	if err != nil {
		fmt.Println(err)
		stop <- 0
		return
	}
	err = exec.Command("sh", "-c", "dmsetup create bronze_volume_"+id+" --table '0 131072 thin /dev/mapper/iron_vg-convoy_Stripe_iron_data "+id+"'").Run()
	if err != nil {
		fmt.Println(err)
		stop <- 0
		return
	}
	fmt.Println("complete create volume bronze_volume_", id)
	stop <- 1
}

func (mutex *mutexLock) createSilverVolume(id string, stop chan int) {
	mutex.mutex_silver.Lock()
	defer mutex.mutex_silver.Unlock()
	err := exec.Command("sh", "-c", "dmsetup message /dev/mapper/silver_vg-convoy_Linear_silver_data 0 "+fmt.Sprintf("'create_thin %s'", id)).Run()
	if err != nil {
		fmt.Println(err)
		stop <- 0
		return
	}
	err = exec.Command("sh", "-c", "dmsetup create silver_volume_"+id+" --table '0 131072 thin /dev/mapper/silver_vg-convoy_Linear_silver_data "+id+"'").Run()
	if err != nil {
		fmt.Println(err)
		stop <- 0
		return
	}
	fmt.Println("complete create volume silver_volume_", id)
	stop <- 1
}

func (mutex *mutexLock) createGoldVolume(id string, stop chan int) {
	mutex.mutex_gold.Lock()
	defer mutex.mutex_gold.Unlock()
	err := exec.Command("sh", "-c", "dmsetup message /dev/mapper/gold_vg-convoy_Linear_gold_data 0 "+fmt.Sprintf("'create_thin %s'", id)).Run()
	if err != nil {
		fmt.Println(err)
		stop <- 0
		return
	}
	err = exec.Command("sh", "-c", "dmsetup create gold_volume_"+id+" --table '0 131072 thin /dev/mapper/gold_vg-convoy_Linear_gold_data "+id+"'").Run()
	if err != nil {
		fmt.Println(err)
		stop <- 0
		return
	}
	fmt.Println("complete create volume gold_volume_", id)
	stop <- 1
}

func (mutex *mutexLock) deleteBronzeVolume(id int, stop chan int) {
	mutex.mutex_bronze.Lock()
	defer mutex.mutex_bronze.Unlock()

	err := devicemapper.RemoveDevice("bronze_volume_" + fmt.Sprintf("%d", id))
	if err != nil {
		fmt.Println(err)
		stop <- 0
		return
	}
	err = devicemapper.DeleteDevice("/dev/mapper/iron_vg-convoy_Stripe_iron_data", id)
	if err != nil {
		fmt.Println(err)
		stop <- 0
		return
	}
	fmt.Println("complete delete volume bronze_volume_", fmt.Sprintf("%d", id))
	stop <- 1
}
func (mutex *mutexLock) deleteSilverVolume(id int, stop chan int) {
	mutex.mutex_silver.Lock()
	defer mutex.mutex_silver.Unlock()

	err := devicemapper.RemoveDevice("silver_volume_" + fmt.Sprintf("%d", id))
	if err != nil {
		fmt.Println(err)
		stop <- 0
		return
	}
	err = devicemapper.DeleteDevice("/dev/mapper/silver_vg-convoy_Linear_silver_data", id)
	if err != nil {
		fmt.Println(err)
		stop <- 0
		return
	}
	fmt.Println("complete delete volume silver_volume_", fmt.Sprintf("%d", id))
	stop <- 1
}

func (mutex *mutexLock) deleteGoldVolume(id int, stop chan int) {
	mutex.mutex_gold.Lock()
	defer mutex.mutex_gold.Unlock()

	err := devicemapper.RemoveDevice("gold_volume_" + fmt.Sprintf("%d", id))
	if err != nil {
		fmt.Println(err)
		stop <- 0
		return
	}
	err = devicemapper.DeleteDevice("/dev/mapper/gold_vg-convoy_Linear_gold_data", id)
	if err != nil {
		fmt.Println(err)
		stop <- 0
		return
	}
	fmt.Println("complete delete volume gold_volume_", fmt.Sprintf("%d", id))
	stop <- 1
}
