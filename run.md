input:
go build -o go-cli .
or
./go-cli --name=Hadi --age=25

output:
Hello, Hadi!
You are 25 years old.

after adding verbose as a bool:
input: 
./go-cli --name=John --age=30 --verbose

output:
Verbose mode enabled
Hello, Hadi!
You are 25 years old.
