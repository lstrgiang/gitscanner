

# Candidate Note



## Prerequisite
- Golang 1.16
- Make 
- Docker/docker-compose

## Tools I used 

- sqlboiler: this library generate repositories that handle database operations, with much faster speed than casual ORM 
- swaggers: this is used for creating swagger documents, can also be used to
generate rest codes (currently I do not apply to avoid redundancy)
- migrate: this is used to run migration scripts, written in sql queries, to manage database schema easily

## My Design

My code base was designed following clean architecture with mainly 3 layers:
- UseCase: this layer is used to handle business logic
- Infrastructure: this layer contains packages related to infrastructure and does not involve into business logic, including repositories and database, midlewares, etc.
- Controller: this is the handler layer that expose API endpoint  


### Directory Structure
- cmd: implemented necessary cmd, including main server application, may contains other kind of utilities if neccessary
- internal: wrapped and encapsulated source of code that will not be exposed
- db: includes sqlboiler configuration and migrations scripts
- docker: includes dockerfile and docker-compose file as well as neccessary configuration files
- bin: mostly executable plugins and generated/built applications
- scripts: bash scripts to handle  operations quickly
- docs: swagger api document for api endpoints, request body schemas and responses  

## What I have done and have not done

### From Problem Statements
- [x] A user can CRUD repositories. A repository contains a name and a link to the repo.
- [x] A user can trigger a scan against a repository.
- [x] A user can view the Security Scan Result ("Result") List

- [x] Project Structure: Clear organization and structure of folders, code and functionality.
- [x] Clean Code: Code Consistency, use of linters, formatting, error handling, and anything else that shows your skills. Simple is better than complex.
- [x] Stack Knowledge: Proper use of Golang and selected frameworks/libraries.
- [x] Implementation: The implementation has to work according to the specs.
- [x] Proper Documentation: 
    - A High-Level Design for the components/infrastructure if any.
    - Describe how you came up with the solution and what makes it a good one for the use-case.
    - Describe what the project is doing, what has been used, how to configure it, how to start it, test it etc.

- [  ] Unit Tests: Covering the core functionality with unit tests (positive and negative test-cases) 
(not much time to do, but I implemented some unit tests to demonstrate how I do unit testing)

**Bonus points for:**
- [x] SQL schema
- [x] API documentation
- [x] Containerized app
- [x] Use of appropriate design patterns
- [  ] Microservice Architecture
- [  ] Any extra feature (just write it in your documentation)



### Additional
- [x] Unit testing for some services to demonstrate how I do unit testing
- [ ] Unit tests cover all core functions (not much time to do)
- [ ] Unit testing for handlers/controllers
- [ ] Generate mock objects
- [ ] Integration testing



# Backend Engineer Coding Challenge



> This repository contains the coding challenge for backend engineers.

**Note:** Please don't fork this repository, create a pull request against it, or use GuardRails in the repo name. Otherwise other candidates may take inspiration from it. Once the coding challenge is completed, you can submit it via this [Google form](https://forms.gle/i5nZWZKoUnTWj3td9).

## Description

Build a simple code scanning application that detects sensitive keywords in public git repos.
The application must fulfil the following requirements:
- A user can CRUD repositories. A repository contains a name and a link to the repo.
- A user can trigger a scan against a repository.
- A user can view the Security Scan Result ("Result") List

How to do a scan:
- Just keep it simple by iterating the words on the codebase to detect SECRET_KEY findings.
- SECRET_KEY start with prefix public_key || private_key.

The Result entity should have the following properties and be stored in a database of your choice:
- Id: any type of unique id
- Status: "Queued" | "In Progress" | "Success" | "Failure"
- RepositoryName: string
- RepositoryUrl: string
- Findings: JSONB, see [example](example-findings.json)
- QueuedAt: timestamp
- ScanningAt: timestamp
- FinishedAt: timestamp

Wherever you'd have to add something that requires product subscriptions or significant extra time, just mention it in your documentation.

**What we want to see:**
- Project Structure: Clear organization and structure of folders, code and functionality.
- Clean Code: Code Consistency, use of linters, formatting, error handling, and anything else that shows your skills. Simple is better than complex.
- Stack Knowledge: Proper use of Golang and selected frameworks/libraries.
- Implementation: The implementation has to work according to the specs.
- Unit Tests: Covering the core functionality with unit tests (positive and negative test-cases).
- Proper Documentation: 
    - A High-Level Design for the components/infrastructure if any.
    - Describe how you came up with the solution and what makes it a good one for the use-case.
    - Describe what the project is doing, what has been used, how to configure it, how to start it, test it etc.

**Bonus points for:**
- SQL schema
- API documentation
- Containerized app
- Use of appropriate design patterns
- Microservice Architecture
- Any extra feature (just write it in your documentation)

**Things you don't have to worry about:**

- Authentication/Authorization
- CI configuration / Deployment
- APM
- Authentication / Authorization / Auditing
