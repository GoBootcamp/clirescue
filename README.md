# CLI Rescue exercise

## Setup

1. Start by forking this repo so you have your own copy.
2. clone the repo in your Go path:

```go
$ cd $GOPATH/src/github.com/<your username>
$ git clone git@github.com:<your username>/cli_rescue.git
```

Install the dependencies:

```go
$ go get github.com/codegangsta/cli
```

If everything is properly setup, you should be able to run your code:

```go
$ go run main.go
```

You can use the following demo Pivotal Tracker account:

* Username: gobootcamp
* Password: santamonica


## Start rescuing

This code was written by someone who just started learning Go.
Beginners often make the same mistakes so assume that this is your code
and you are now tasked to refactor it.
The end goal is to create a CLI to interface with the [Pivotal Tracker
API](https://www.pivotaltracker.com/help/api/rest/v5).

The first step is to return the user details and to store the user token
on the drive so next calls can automatically be made without asking for
the user's username and password.

### A few hints:

* Understanding Go path's system might be a bit tricky at first.
* Forking the repo might affect your import statements.
* The original author might not have made a good use of packages.
* You more thna likely have some OOP experience, you might want to
  leverage it in this exercise.
* Checkout the `init` function.
* Curious about how a 3rd party lib works? Try go doc.
* Simpler is better.
* Look into how to write tests.
* Bonus point for removing 3rd part dependencies.

As always, refer to the [Effective Go doc](http://golang.org/doc/effective_go.html) if you have any questions.
