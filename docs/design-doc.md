# Design Doc.



###### versions history 
| Version      | Date			| Description  	| Authors         |Status|
| :---------: | :-------------:| ---------| ----------------|:--:|
| 1            | 04/10/2014    | Creation 		|@First_developer |<span style="color:orange">draft</span>|

<!--| 2            | 04/10/2014    | Creation 		|@First_developer |<span style="color:orange">draft</span>|
-->


## Introduction

This document provides to documentation, as complete as possible, of the software that is built here. It tries to give a good understanding of what is to be built and how is to be built. 

As a result, people, software developers or contributors we be able to catch the main purpose of this software, understand the philosophy behind the design and the implementation of the software. 

Thus, everyone will be able to easily contribute to the software without the fear of departing from the goal behind it.


## Motivation
As a software developer, we have to make a bit of application management. We are accountable for the well-working of the application in production. 

To monitor those applications, we use many different scripts and command line tools to accomplish management  or monitoring tasks. Those scripts are often either internal or external which are often written in Bash, Python, Ruby or many other scripting languages. 

But, arrives the moment when you get full of scripts and noone knows the purpose of each scritps. Even more, everybody makes copies of existing scripts, customizes them or simply renames them and that's make the life of new commers harder.

Logan is a command line tool that aim to organize all thoses scripts in a more comprehensive and consistant way.

It introduces a simple sementic way of organising scripts. Everyone or every team who want to organize their scripts, must agree on the terms/vocabulay they'll used to describe tasks execute by those scripts. They should categorize scripts using Logan scpecific idioms.  

  


## Overview
Logan use a simple set of idioms to help organizing your scripts. It should be consider like a helper agent that execute what we want by giving him an order. 

The logan agent knows how to execute **actions**.

##Logan idioms
That's the list of idioms used in the logan world to help categorizing your tasks in a meaningful way.

####Actions
Every tasks executed by a logan agent is called a **logan action**. A logan action is composed of a **goal** with a given **requirements**.

```
<intent>:<target>:[<context>] [<requirement>...]
``` 

######Exemple : Show a command help by given its name as a requirement
```
show:help COMMAND="create_new_app" APP_NAME="sampleApp"
``` 
As a result, we got : 

>**intent:** show<br>
**target:** help<br>
**content:** default<br>

>**requirements:** <br>
>COMMAND = "create_new_app"<br> 
APP_NAME = "sampleApp"

####Configuration


####Use cases

#### Dependencies

####Performance



## Implementation


#### Architecture  


#### Modules


#### Data structures  


#### Processing  

## Extensibility

####Custom providers



