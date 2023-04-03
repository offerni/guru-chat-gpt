# guru-chat-gpt
Endpoint receives a message as query param, queries all guru cards related to that user token, sanitizes, stringifies the content, split it into chunks of a customizeable size, feeds those chunks to chat GPT's as "system" role messages and calls the API with the provided dataset + the message from the request.
This will allow chatgpt to provide an answer based on the content of the dataset provided. 
There are currently some limitations as gpt-3.5-turbo API only allows `4097` tokens and there's no way around that.

The main functionality is done but it needs better handling to summarize the dataset before feeding it to chat gpt

Example:
```
http://localhost:9091/search?sessionID=92bb3bca-c548-439c-baf8-77c37fb0d5e7&message="Hello World"
```
* `message` is the promt the client wants to chat about
* `sessionID` is required and ensures uniqueness preventing the wrong client to listen to the event sent

This is a work in progress.
