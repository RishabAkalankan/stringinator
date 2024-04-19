## Decisions Taken
1. Introduced zerolog as it is really fast. Also written a wrapper layer around zerolog so that some custom initializations can be done in the new logger package. All the other packages can now make use of this wrapper logger to log the messages.
2. Provided capability to log a message with request-id. This will be useful incase the same request is being transferred to some other service for performing additional operations. Using this request-id, the different services that are invoked can be easily identified.
3. Added a new package `github.com/go-playground/validator` to carry out input request validations.
4. When the input is having some accented letters like `âç` etc, the length of the string becomes unpredictable. So to avoid confusions, rune length is taken instead of the string length.
5. Restructured the code into packages to support future extensibility.
6. Added Makefile for automating the process of building the apps, running the application, etc. 
7. Rate Limiter is implemented with 10 req/s and 5 burst reqs
8. Each API request is logged for making the debugging easier. Log is of the format `Received Request for: [METHOD] URI`

