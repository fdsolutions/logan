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
Logan use a simple set of idioms to help organizing your scripts. It should be consider as a helper agent that executes any giving actions. 

The logan agent knows how to execute **actions**. Every action's definitions o define all definitions of your actions in one or many configuration files. Te 





##Configuration


##Logan idioms
That's the list of idioms used in the logan world to help categorizing your tasks in a meaningful way.


####Action
Every tasks executed by a logan agent is called a **logan action**. A logan action is composed of a **goal** with a given **parameters**.

```
<intent>:<target>:[<context>] [<parameter>...]
``` 

**NOTE** 
>The part `<intent>:<target>:[<context>]` is called the **goal**. 

######Exemple : Show a command help by given its name as a requirement
```
show:help COMMAND="create_new_app" APP_NAME="sampleApp"
``` 
As a result, we got : 

>**intent:** show<br>
**target:** help<br>
**context:** default<br>

>**parameter:** <br>
>COMMAND = "create_new_app"<br> 
APP_NAME = "sampleApp"

**NOTE** 

>It's important to note that the action `show:help`, has no context set. When no context is provide in the action comamnd, the context is set to `'default'`.
A context set to default means that the action belongs to the default category of actions that's is to say the **logan category**.

####Intent
It's a verb describing the action to you need the logan agent to execute. The intent must stick to the following conventions: 

* It must be an infinitive verb without the preposition `'to'` like `'sleep'`, ` 'dance'`, `'show'`
* It must be contains no space or special characters other than the underscore '_'. In other words, It must be underscored. Eg: 'find-out' (NO), 'find_out' (YES)
* It should be lowercased.

####Target
It's the object with want to operate on or the item requested by the intent action.
In our previous example `show:help`, `help` is the target, It is requested by the intent action `show`.

####Context
Suppose that you want to install an application and you have two different way depending on the operating system your are on Ubuntu OS or CentOs Redhat OS. It this example, those two OS represent the context of the aplication.
To execute those tasks, we need to launch the following commands : 

* Ubuntu : `sudo apt-get install <APP_NAME>`<br>
* CentOS : `sudo yum instal <APP_NAME>`

If we try to use **logan** to manage those two installations, we should use a meaning full action name that best describe the action you want to perform : `install:app`. However, we'll have a conflict when launching that action because, we have no idea on which command should be used  



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



