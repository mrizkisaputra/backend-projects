# Expense Tracker
Sample solution for the [Github User Activity](https://roadmap.sh/projects/expense-tracker) challenge from [roadmap.sh](https://roadmap.sh/). Expense tracker is a project to manage expense your finances. The application run from the command line interface (CLI).
Some of the features in this application include:

- add, update, and delete expense
- view of all expenses
- view summary of all expenses and  
  summary of expenses for a specific month (of current year)
- allow user to export expense to a CSV file

## How to run this project application?
1. Install [Go language](https://go.dev), and setup system environment. Check the installation ``go version``
2. Download ZIP file or use ``git clone https://github.com/mrizkisaputra/backend-projects.git``
3. Open terminal and navigate to directory project **expense-tracker** ``cd /mrizkisaputra-backend-projects/expense-tracker``
4. Run ``go install``
5. List of available commands in example below


## Example
```shell
# to see the list of available commands
expense-tracker --help

#add expense: description & amount is mandatory, category is optional
expense-tracker add --description="Launch" --amount=20 --category="food"
#output: Expense added successfully (ID: 1)

#update expense: amount is mandatory, description is optional
expense-tracker update 1 --amount=100
#output: Expense updated successfully (ID: 1)

#delete expense
expense-tracker delete 1
#output: Expense deleted successfully (ID: 1)

#view all expenses
expense-tracker list

#view expenses by category
expense-tracker list --category="food"

#view summary expenses
expense-tracker summary

#view summary expenses for a specific month of current year
expense-tracker summary --month=09

#export expenses: only support export to a CSV file
expense-tracker export --ext="csv"
```