# polling-app

Simple polling application using go.

The aim of this project is to learn more about go.

The idea is to showcase skills: 

- create APIs using go
- postgres for persistence
- redis for cache
- frontend using react

## Polling/Voting application

Design API for creating, managing and voting in polls.

Use go, react, redis, postgres, docker, nginx.

Use combination of client-ip and user-agent to identify unique participants.

### Voting systems

- first past the post - single vote, highest tally wins
- ranked, strict ordering of voting options, where each position is weighted appropriately
- ranked, strict ordering with run off candidates until a single candidate has majority
- rating, assign each option a rating, tallying total across all candidates

### Create Poll

- title
- optional time limit
- vote tallying strategy
- vote candidates
- unique voter or multiple voter

Validate fields.
Create unique identifier for poll.
Save poll entry in database.
Return link for poll.

### Viewing Poll

By direct link only.

If poll still active:

If client-ip and user-agent not seen before for this poll, offer ability to vote.
If user known and voted, show waiting for poll to end.
If user known and not voted, offer ability to vote.
If user is creator of poll, show option to end poll early.

If poll closed:

Show poll results.

### End Poll

Stop accepting new votes and tally the results.


## Implementation



add lines into nginx.conf: 
proxy_set_header Host $host;
proxy_set_header X-Real-IP $remote_addr;
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

## Running

### Docker

This assumes you can run docker and docker-compose as non-root user.
https://docs.docker.com/engine/install/linux-postinstall/
