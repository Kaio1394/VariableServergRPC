# VariableServer gRPC 
This a student project to know about api gRPC. In this same account in GitHub there is a repository 'VariableServer', so that build as API REST.

## About
This a simple api to save variables to be use by test automation or other Technologies to share information.
Example:
If I have a test automation where there's are two scenarios:
- Scenario to verify status of Module X;
- Scenario to open Module X;

In the end of status scenario, I can make my automation scenario create a variable by my api gRPC called 'STATUS_MODULE'. Let's imagine that the variable has the value 'O' (means Open).
When a call my second scenario, in my first step, I can read the variable 'STATUS_MODULE', and if I want to Open my Module, I need that the value of variable to be equal 'C' (means Close), if the variable is not equal 'O', I can't run my second scenario, because my Module it's already open.

## How to use

## Technologies used:
- Golang
- Api gRPC
- Sqlite
- Gorm
- Logrus(Log)