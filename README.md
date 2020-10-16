# About 
As you know AWS Api Gateway doesn't allow you to set multiple entries to be returned in response as 
"ACCESS_CONTROL_ALLOW_ORIGIN" header. 

There are some questions on StackOverflow related to that issue:
- [Link 1](https://stackoverflow.com/questions/39628640/aws-api-gateway-cors-access-control-allow-origin-multiple-entries)
- [Link 2](https://stackoverflow.com/questions/40230648/aws-api-gateway-cors-access-control-allow-origin-regex-with-cloudformation)

So as it was suggested we can create a lambda function to handle request with OPTION headers and response 
with a correct domain name in the header.

This lambda function, written in `Go` might help you to tackle that problem.
