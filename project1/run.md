گرفتن اطلاعات نمایندگی
برای گرفتن اطلاعات کامل یک نمایندگی با استفاده از شناسه (ID) آن، دستور زیر را اجرا کنید:
go run main.go -command=get -region=tehran


ایجاد نمایندگی جدید
برای ایجاد یک نمایندگی جدید در یک استان خاص، دستور زیر را اجرا کنید:
go run main.go -command=create -region=tehran


ویرایش اطلاعات نمایندگی
برای ویرایش اطلاعات یک نمایندگی با استفاده از شناسه (ID) آن، دستور زیر را اجرا کنید:
go run main.go -command=edit -region=tehran


وضعیت نمایندگی‌ها
برای مشاهده تعداد نمایندگی‌ها و تعداد کل کارکنان در یک استان خاص، دستور زیر را اجرا کنید:
go run main.go -command=status -region=tehran
$ go run main.go -command=list -region=tehran
ID: 1, Name: Branch1, Address: Address1, Phone: 1234567890, Join Date: 2024-01-01, Employees: 10
ID: 2, Name: Branch2, Address: Address2, Phone: 0987654321, Join Date: 2024-01-01, Employees: 15

گرفتن اطلاعات نمایندگی
$ go run main.go -command=get -region=tehran
Enter Branch ID: 1
ID: 1, Name: Branch1, Address: Address1, Phone: 1234567890, Join Date: 2024-01-01, Employees: 10


ایجاد نمایندگی جدید
$ go run main.go -command=create -region=tehran
Enter Name: Branch3
Enter Address: Address3
Enter Phone Number: 1122334455
Enter Join Date: 2024-01-01
Enter Number of Employees: 20


ویرایش اطلاعات نمایندگی
$ go run main.go -command=edit -region=tehran
Enter Branch ID: 1
Enter Name: NewBranch1
Enter Address: NewAddress1
Enter Phone Number: 1111111111
Enter Join Date: 2024-02-02
Enter Number of Employees: 12


وضعیت نمایندگی‌ها
$ go run main.go -command=status -region=tehran
Total branches in tehran: 3
Total employees in tehran: 45



