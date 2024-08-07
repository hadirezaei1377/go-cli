package project1

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Branch struct {
	ID          int
	Name        string
	Address     string
	PhoneNumber string
	JoinDate    string
	EmployeeNum int
	Region      string
}

var (
	command string
	region  string
)

func init() {
	flag.StringVar(&command, "command", "", "Command to execute (list, get, create, edit, status)")
	flag.StringVar(&region, "region", "", "Region of the branch")
}

func main() {
	flag.Parse()
	branches := readData()

	switch command {
	case "list":
		listBranches(branches, region)
	case "get":
		var id int
		fmt.Print("Enter Branch ID: ")
		fmt.Scanf("%d", &id)
		getBranch(branches, region, id)
	case "create":
		createBranch(branches, region)
	case "edit":
		var id int
		fmt.Print("Enter Branch ID: ")
		fmt.Scanf("%d", &id)
		editBranch(branches, region, id)
	case "status":
		statusBranches(branches, region)
	default:
		fmt.Println("Invalid command. Use one of the following: list, get, create, edit, status.")
	}
}

func readData() []Branch {
	file, err := os.Open("data/result.csv")
	if err != nil {
		if os.IsNotExist(err) {
			return []Branch{}
		}
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	var branches []Branch
	for _, line := range lines {
		id, _ := strconv.Atoi(line[0])
		employeeNum, _ := strconv.Atoi(line[5])
		branch := Branch{
			ID:          id,
			Name:        line[1],
			Address:     line[2],
			PhoneNumber: line[3],
			JoinDate:    line[4],
			EmployeeNum: employeeNum,
			Region:      line[6],
		}
		branches = append(branches, branch)
	}
	return branches
}

func writeData(branches []Branch) {
	file, err := os.Create("data/result.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, branch := range branches {
		line := []string{
			strconv.Itoa(branch.ID),
			branch.Name,
			branch.Address,
			branch.PhoneNumber,
			branch.JoinDate,
			strconv.Itoa(branch.EmployeeNum),
			branch.Region,
		}
		writer.Write(line)
	}
}

func listBranches(branches []Branch, region string) {
	for _, branch := range branches {
		if strings.EqualFold(branch.Region, region) {
			fmt.Printf("ID: %d, Name: %s, Address: %s, Phone: %s, Join Date: %s, Employees: %d\n",
				branch.ID, branch.Name, branch.Address, branch.PhoneNumber, branch.JoinDate, branch.EmployeeNum)
		}
	}
}

func getBranch(branches []Branch, region string, id int) {
	for _, branch := range branches {
		if branch.ID == id && strings.EqualFold(branch.Region, region) {
			fmt.Printf("ID: %d, Name: %s, Address: %s, Phone: %s, Join Date: %s, Employees: %d\n",
				branch.ID, branch.Name, branch.Address, branch.PhoneNumber, branch.JoinDate, branch.EmployeeNum)
			return
		}
	}
	fmt.Println("Branch not found.")
}

func createBranch(branches []Branch, region string) {
	var branch Branch
	branch.ID = len(branches) + 1
	branch.Region = region

	fmt.Print("Enter Name: ")
	fmt.Scanf("%s", &branch.Name)
	fmt.Print("Enter Address: ")
	fmt.Scanf("%s", &branch.Address)
	fmt.Print("Enter Phone Number: ")
	fmt.Scanf("%s", &branch.PhoneNumber)
	fmt.Print("Enter Join Date: ")
	fmt.Scanf("%s", &branch.JoinDate)
	fmt.Print("Enter Number of Employees: ")
	fmt.Scanf("%d", &branch.EmployeeNum)

	branches = append(branches, branch)
	writeData(branches)
}

func editBranch(branches []Branch, region string, id int) {
	for i, branch := range branches {
		if branch.ID == id && strings.EqualFold(branch.Region, region) {
			fmt.Print("Enter Name: ")
			fmt.Scanf("%s", &branch.Name)
			fmt.Print("Enter Address: ")
			fmt.Scanf("%s", &branch.Address)
			fmt.Print("Enter Phone Number: ")
			fmt.Scanf("%s", &branch.PhoneNumber)
			fmt.Print("Enter Join Date: ")
			fmt.Scanf("%s", &branch.JoinDate)
			fmt.Print("Enter Number of Employees: ")
			fmt.Scanf("%d", &branch.EmployeeNum)

			branches[i] = branch
			writeData(branches)
			return
		}
	}
	fmt.Println("Branch not found.")
}

func statusBranches(branches []Branch, region string) {
	totalBranches := 0
	totalEmployees := 0

	for _, branch := range branches {
		if strings.EqualFold(branch.Region, region) {
			totalBranches++
			totalEmployees += branch.EmployeeNum
		}
	}

	fmt.Printf("Total branches in %s: %d\n", region, totalBranches)
	fmt.Printf("Total employees in %s: %d\n", region, totalEmployees)
}
