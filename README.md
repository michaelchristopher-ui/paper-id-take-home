# Paper.id take home task: Disbursement Endpoint

  

A web application that has one API disbursement, accessible through localhost:<PORT>/disbursement/disburse
which accepts a json with the key: account_id_from, account_id_to, amount and description. The API disburses the amount specified from the "from" account to the "to" account.
Sample accounts with the ID 1 and 2 are available with 3000 credits.

A postman collection is available for use which has an example request that can be referred to as well.

# Tech Stack

- Golang 1.21.1

- Echo

# Setup

## Installing Dependencies

  

- This is a Golang app. Download Golang here (https://go.dev/). Make sure you are downloading the version compatible with the tech stack.
- or if you are on ubuntu, use snap install go --channel=1.21/stable --classic


## Running the app

- make all


This will do three things as written in the Makefile:

- Setup, which basically calls go mod vendor.
- Build, which builds the application.
- Run, which runs the application with the specified config.


The app will be run locally on the port specified in the cmd/app/config.yaml file. The default port is 8008.
