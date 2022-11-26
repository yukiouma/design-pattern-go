# State

> State is a behavioral design pattern that lets an object alter its behavior when its internal state changes. It appears as if the object changed its class.

## Assumption
Imagine you are building a news app, you are writing a feature to achieve new pulishing event.
If user want to pulish a news, it need to go througth 3 steps:
* user draft the news
* let admin user moderate the news
* if the news passed moderation, it can be published

so, a news contains 3 states and 3 operation:

state:
  * draft
  * moderation
  * published

operation:
  * update
    update content of news, only works in draft state
  * publish
    draft -> moderation -> publshed
  * revert
    published -> draft
    moderation -> draft