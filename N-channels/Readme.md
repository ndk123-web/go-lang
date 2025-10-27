# Channels

1. Listeners default value 0 if not sending any routine
2. Senders Directly Gives Deadlock issue if not anyone listening

| Concept         | Unbuffered Channel                     | Buffered Channel                 |
| --------------- | -------------------------------------- | -------------------------------- |
| Communication   | Direct handoff between sender/receiver | Through a queue of capacity `N`  |
| Blocking Rule   | Send blocks until recv ready           | Send blocks only if buffer full  |
| Receive Rule    | Recv blocks until send ready           | Recv blocks only if buffer empty |
| Data stored?    | No                                     | Yes (temporary)                  |
| Synchronization | Strong (strict timing)                 | Loose (decoupled timing)         |
| Use Case        | Real-time signaling                    | Producer-consumer queues         |
