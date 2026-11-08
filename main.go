package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	allTasksTxt string = "Tasks.txt"
	// finishedTasksTxt string = "Finish.txt"
	// inWorkTxt        string = "Working.txt"
	// unFinishedTxt    string = "Ubfinished.txt"
	flag string
)

func get_number() int {
	tasks := loadData()
	if len(tasks) < 1 {
		return 1
	}

	lastLine := tasks[len(tasks)-1]
	splitting := strings.SplitN(lastLine, "|", 2)
	if len(splitting) < 1 {
		return len(tasks) + 1
	}
	number, err := strconv.Atoi(strings.TrimSpace(splitting[0]))
	if err != nil {
		return len(tasks) + 1
	}
	return number + 1
}

func loadData() []string {
	file, err := os.OpenFile(allTasksTxt, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("ошибка при открытии файла: %s\n", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tasks []string
	for scanner.Scan() {
		text := scanner.Text()
		tasks = append(tasks, text)
	}
	if err = scanner.Err(); err != nil {
		fmt.Printf("ошибка при чтении файла: %s\n", err)
		return nil
	}
	return tasks

}

func inputData(prompt string) string {
	for {
		fmt.Println(prompt)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("ошибка: %s\nпопробуйте ещё раз", err)
			continue
		} else {
			return strings.TrimSpace(input)
		}
	}
}

func writeFile(filename string, data string, flag string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("ошибка при открытии файла: %s\n", err)
		return
	}
	defer file.Close()
	n := get_number()
	result := fmt.Sprintf("%d | %s | %s\n", n, data, flag)
	if _, err := file.WriteString(result); err != nil {
		fmt.Printf("ошибка записи: %s\n", err)
		return
	}
	fmt.Printf("Успешное добавление данных\n")
}

func addTask() {
	newTask := inputData("Введите название новой задачи")
	flag = "add"
	writeFile(allTasksTxt, newTask, flag)
}

func deleteTask() {
	task := inputData("Введите номер задачи")
	number, err := strconv.Atoi(task)
	if err != nil {
		fmt.Printf("Нужно ввести число")
		return
	}
	tasks := loadData()
	if number < 1 || number > len(tasks) {
		fmt.Printf("Такой задачи не существует\n")
		return
	}
	number--
	tasks = append(tasks[:number], tasks[number+1:]...)
	file, err := os.OpenFile(allTasksTxt, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("ошибка при открытии файла: %s\n", err)
		return
	}
	defer file.Close()

	for _, line := range tasks {
		file.Write([]byte(line + "\n"))
	}
}

func takeTask() {
	task := inputData("Введите номер задания, которе вы хотите выполнить: ")
	flag = "inProgress"
	number, err := strconv.Atoi(task)
	if err != nil {
		fmt.Printf("введите число\n")
		return
	}
	tasks := loadData()
	for i, line := range tasks {
		if i+1 == number {
			s := strings.Split(line, " | ")
			s[2] = flag
			result := fmt.Sprintf("%s | %s | %s", s[0], s[1], flag)
			tasks[i] = result
		}
	}
	file, err := os.OpenFile(allTasksTxt, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("ошибка при открытии файлы: %s", err)
		return
	}
	defer file.Close()

	for _, line := range tasks {
		file.Write([]byte(line + "\n"))
	}
}

func finishTask() {
	task := inputData("Введите номер задания, которе вы выполнили: ")
	flag = "done"
	number, err := strconv.Atoi(task)
	if err != nil {
		fmt.Printf("введите число\n")
		return
	}
	tasks := loadData()
	for i, line := range tasks {
		if i+1 == number {
			s := strings.Split(line, " | ")
			s[2] = flag
			result := fmt.Sprintf("%s | %s | %s", s[0], s[1], flag)
			tasks[i] = result
		}
	}
	file, err := os.OpenFile(allTasksTxt, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("ошибка при открытии файлы: %s", err)
		return
	}
	defer file.Close()

	for _, line := range tasks {
		file.Write([]byte(line + "\n"))
	}
}

func finishedTasks() {
	file, err := os.OpenFile(allTasksTxt, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("ошибка при открытии файла, попробуйте ещё раз: %s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "done") {
			fmt.Printf("%d. %s\n", i, text)
			i++
		}
	}
	if err = scanner.Err(); err != nil {
		fmt.Printf("ошибка при чтении файла: %s\n", err)
		return
	}
}

func takenTasks() {
	file, err := os.OpenFile(allTasksTxt, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("ошибка при открытии файла, попробуйте ещё раз: %s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "inProgress") {
			fmt.Printf("%d. %s\n", i, text)
			i++
		}
	}
	if err = scanner.Err(); err != nil {
		fmt.Printf("ошибка при чтении файла: %s\n", err)
		return
	}
}

func allTasks() {
	file, err := os.OpenFile(allTasksTxt, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("ошибка при открытии файла, попробуйте ещё раз: %s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("%d. %s\n", i, text)
		i++
	}
	if err = scanner.Err(); err != nil {
		fmt.Printf("ошибка при чтении файла: %s\n", err)
		return
	}
}

func commands() {
	fmt.Printf("Список команд:\n")
	fmt.Printf("1) Add - добавить задачу\n")
	fmt.Printf("2) Delete - удалить задачу\n")
	fmt.Printf("3) Take - взять задачу в работу\n")
	fmt.Printf("4) Finish - завершить задачу\n")
	fmt.Printf("5) Finished - показать завершённые задачи\n")
	fmt.Printf("6) Taken - показать задачи в работе\n")
	fmt.Printf("7) All - показать все задачи\n")
	fmt.Printf("8) Commands - список команд\n")
	fmt.Printf("9) Exit - выход\n")
}

func start() {
	var (
		request string
		//prompt  string
	)
	for {
		request = inputData("Введите команду")
		switch request {
		case "Add":
			addTask()
		case "Delete":
			deleteTask()
		case "Take":
			takeTask()
		case "Finish":
			finishTask()
		case "Finished":
			finishedTasks()
		case "Taken":
			takenTasks()
		case "All":
			allTasks()
		case "Commands":
			commands()
		case "Exit":
			return
		default:
			fmt.Println("Команда не найдена")
			fmt.Println("Вы можете воспотльзоваться командой \"Commands\", чтобы увидеть все команды")
		}
	}
}

func main() {
	start()
}
