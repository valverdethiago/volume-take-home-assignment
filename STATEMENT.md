# Senior Software Engineer Take-Home Programming Assignment for GoLang

## Story
There are over 100,000 flights a day, with millions of people and cargo being transferred around the world. 
With so many people and different carrier/agency groups, it can be hard to track where a person might be. In order 
to determine the flight path of a person, we must sort through all of their flight records.

## Goal
To create a simple microservice API that can help us understand and track how a particular person's flight 
path may be queried. The API should accept a request that includes a list of flights, which are defined by a source 
and destination airport code. These flights may not be listed in order and will need to be sorted to find the total 
flight paths starting and ending airports.

## Required JSON structure
- [["SFO", "EWR"]]                                                 => ["SFO", "EWR"]
- [["ATL", "EWR"], ["SFO", "ATL"]]                                 => ["SFO", "EWR"]
- [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]

## Specifications 
- Your miscroservice must listen on port 8080 and expose the flight path tracker under the /calculate endpoint.

- Create a private GitHub repo and add https://github.com/taariq, https://github.com/measure-fi, and 
https://github.com/verabehr as collaborators to the project. Please only add the collaborators when you are sure you 
are finished.

- Define and document the format of the API endpoint in the README.

- Use Golang and/or any tools that you think will help you best accomplish the task at hand.

- When you are done with the assignment, follow up and reply-all to the email that directed 
you to this document. Include your private github link and an estimate of how long you spent on the task 
and any interesting ideas you wish to share.
