I made a task tracker CLI that helps manage tasks. It has statuses for each task which can be modified by the user. Each task will also have a created and modified time which are stored in the task,JSON file which will also hold previously created tasks. The program will load previous tasks and have "memory" in between sessions.

To launch the program, download all program files and build the project or run this command in your terminal (windows): go run ./main.go ./functions.go ./utils.go

The list of functions are given on running the program, run them fully lowercase.

On running the program, a tasks.json file will be created if it does not already exist, if it does, then it will be read and the present tasks will be read into the saveData struct which houses all task information in the program. 

The functions that modify or add tasks will automatically be added to the tasks.json and the tasks.json will be read it when the project is started up. 

Project URL: https://roadmap.sh/projects/task-tracker
