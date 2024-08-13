package main

import (
	"fmt"
	"sort"
)

type Action interface {
	AddTask()
	SeeList()
	FinishTask()
}

var TaskId int

type Task struct {
	TaskName string
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}

var TaskList = make(map[string]bool)
var TaskManager = make(map[int]string)
var TaskOrder []int

func (task *Task) checkTask() bool {
	if _, exists := TaskList[task.TaskName]; exists {
		return true
	}
	return false
}

func (task *Task) AddTask(name string) {
	task.TaskName = name
	if task.checkTask() {
		fmt.Println("Task sudah ada")
	} else {
		TaskId++
		TaskList[task.TaskName] = true
		TaskManager[TaskId] = task.TaskName
		TaskOrder = append(TaskOrder, TaskId)
		fmt.Println("Task ditambahkan")
	}
}

func (task *Task) SeeList() {
	if len(TaskList) == 0 {
		fmt.Println("Tidak ada tugas")
	} else {
		fmt.Println("Tugas:")
		sort.Ints(TaskOrder)
		for _, id := range TaskOrder {
			fmt.Printf("%d. %s\n", id, TaskManager[id])
		}
	}
}

func (task *Task) FinishTask(taskId int) {
	if taskName, exists := TaskManager[taskId]; exists {
		delete(TaskList, taskName)
		delete(TaskManager, taskId)
		for i, id := range TaskOrder {
			if id == taskId {
				TaskOrder = append(TaskOrder[:i], TaskOrder[i+1:]...)
				break
			}
		}
		fmt.Println("Task selesai")
	} else {
		fmt.Println("Task tidak ditemukan")
	}
}

func main() {
	fmt.Println("Selamat datang di to-do list")
	for {
		var option int
		fmt.Printf("1. Lihat tugas\n2. Tambah tugas\n3. Selesaikan Tugas\n4. Keluar\n")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&option)

		var task Task
		switch option {
		case 1:
			clearTerminal()
			task.SeeList()
		case 2:
			var name string
			clearTerminal()
			fmt.Print("Masukkan tugas: ")
			fmt.Scanln(&name)
			task.AddTask(name)
		case 3:
			var taskId int
			clearTerminal()
			fmt.Print("Masukkan ID tugas: ")
			fmt.Scan(&taskId)
			task.FinishTask(taskId)
		case 4:
			clearTerminal()
			fmt.Println("Terima kasih")
			return
		}
	}
}