package employee

import (
	"fmt"
)

type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

var employeeInstance *employee //pointer

func Init( //มองว่า new คือ constructor
	firstname string,
	lastname string,
	totalLeaves, leavesTaken int) *employee {

	if employeeInstance == nil { // nil คือยังไม่ถูกสร้าง **ใช้ได้กับ pointer เท่านั้น
		employeeInstance = &employee{
			FirstName:   firstname,
			LastName:    lastname,
			TotalLeaves: totalLeaves,
			LeavesTaken: leavesTaken,
        }
        
	}

	return employeeInstance

}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
