# page-statistics
page-statistics

CONTENTS OF THIS FILE
---------------------

 * Introductionn
 * Requirements
 * Installation
 * Maintainers

INTRODUCTION
------------

 * 
 * To start this application, go to root folder and run
 * go run main.go
 *  
 * This will expose the rest API on .
 * The rest api details are -
 * 
 * [GET] Rest API will return the message reated to health check else return error
 *  Run the command - curl http://localhost:8888/ping
 *  Outcome - "Health check is OK"
 * 
 *  
 * [POST] API to post the view counter of article id. Run the below command to post the request.
 *  curl -H 'Content-Type: application/json' -d '{"article_id":"abc"}' -X POST http://localhost:8888/counter/v1/statistics
 * 
 *  
 * [GET] API to get the details of views of a particular article id
 *  curl http://localhost:8888/counter/v1/statistics/article_id/{article_id}
 *  example curl http://localhost:8888/counter/v1/statistics/article_id/abc
 *  

REQUIREMENTS
-------------------

 * Golang must be installed properly.
 * Update the database information in .env file at location docker/.env

INSTALLATION
-----------

 * To start this application, go to root folder and run
 * go run main.go

MAINTAINERS
-----------

Current maintainers:

 * Rahul kumar (rahul2u88@gmail.com)

