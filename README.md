# Comcast DevX Coding Interview

Through this exercise we wish to assess a candidates ability to execute on some of the more common types of tasks we see on a daily basis.  This exercise includes web services, string/collection manipulation, with a lead into design decisions for conversation during the interview.

Your job is to get the app running, [implement some features](#features-for-you-to-implement), and send us your implementation.

## Getting Started

To get you off to the right start we've provided an initial implementation of a basic rest service.  It provides an initial implementation of the `/`, `/stringinate` and `/stats` endpoints.  This implementation is available in 3 different languages:

- [Python](python/)
- [Java](java/)
- [Go](go/)

Each contains some instructions for requirements and how to get started in that environment.

## Guidelines

We don't want to restrict your language choice here.  If your most productive language is NodeJs, Rust, etc. feel free to port the starting service over to the rest framework of your choice.  

Feel free to Google for things when you have questions.  We'd love some insight into your research process so please capture references (URLs) to sites you find helpful in solving these tasks.  Especially if you are tackling the tasks in a language with which you are not the most familiar.  References should be collected in the [CITATIONS.md](CITATIONS.md) file.

It's not expected that this exercise take more than an hour or two.  That being said, if you use
a language/framework combination other than those provided it could run over that time. We've provided a [NOTES.md](NOTES.md) file for you to document any decisions taken or things you'd like to communicate about you solutions.

This is a git project.  We prefer many small commits over one large commit.  It helps us understand your progression towards completion as well as giving us smaller increments to review.

## Submit

When you are done with the exercise, please send us:

* A zip of your project with preserved git files, which includes:
    1. Required features
    2. 1 Developer Choice feature
    3. Citations
    4. Decision document

Important Note: While compressing the files and folders for sharing your solution with us, please remove all the dependencies (say jar files under target directory, vendor modules etc.) downloaded while building the application. This is to avoid any increase the size of the zip file for hassle free transfer and/or possible security filters from blocking the zip file for sharing.

## The API

The application supports a small set of API endpoints that can be used to get information about and manipulate string values.  The application also tracks statistics about all the strings that have been sent to the server.

`/`

The root of the server, displays info about the other endpoints.  This is the only endpoint that does not return JSON.

`/stringinate`

Get all of the info you've ever wanted about a string. Accepts GET and POST requests.  For POSTs the endpoint takes JSON of the following form: 
    
```
    {"input":"your-string-goes-here"}
```
    
For GETs an input string is specified as `?input=<your-input>`.

`/stats`

Get statistics about all strings the server has seen, including the number of times each input has been received along with the longest and most popular strings etc.

## Features for You to Implement

Now that you have a running server it's time to add some features and make some changes to the application.

There are a few features that need to be added to the user-facing API:

* For the `/stringinate` endpoint, for a given input string we need to find the character that occurs most frequently and add that character, along with its number of occurrences to the API response JSON. You decide how to represent this in the JSON response.  Ignore white space and punctuation.
* For the `stats` endpoint, track which string input has been seen the most times. Return this value as the `most_popular` key in the response JSON.
* For the `stats` endpoint, track which string input is the longest string to be seen by the server and return as the `longest_input_received` key in the response JSON.

### Developer's Choice

It's your turn to add a feature or make other changes to improve the application.  Implement one new feature idea or change that makes the code better.

Some things you might consider:

* Could there be improvements to logging that might help with troubleshooting problems?
* Are there any input/output validations that might be helpful?
* Would the application benefit from unit or functional tests?
* Are there other interesting string manipulations or stats to implement? Longest palindrome? Anagram finder? Is a given string a dictionary word? The possibilities are virtually endless!
* Could any caching be added to reduce computation for a given input string?
* Could the code be refactored for better maintainability, understandability, extensibilty, etc?
* Could statistics be optionally stored outside of the app so that they can persist across server restarts?
* Would rate limiting to 5rps from a given IP be a useful tool for mitigating excess traffic to the server?

# The GoLang service implementation

This repo contains a small [Echo](https://echo.labstack.com/) application containing the initial API endpoints.

## Setup

Clone the repo, make sure you have golang => 1.16 installed, and then run `./install.sh`.

Once the dependency install has completed, execute `./run.sh` to launch the application.

## Testing Your Local Server

Using `curl` send the following test request to the Flask application.

```
curl -id '{"input":"this is the input"}' -H 'Content-Type: application/json' http://localhost:1323/stringinate
```

If the application is running you should see the following output in your terminal:

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Mon, 07 Jun 2021 14:58:57 GMT
Content-Length: 42

{"input":"this is the input","length":17}
```