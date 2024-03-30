# real-time-forum-typing-in-progress

A typing in progress engine is a way that people can see that a user is typing in real time. Allowing you to see the other user is replying or sending a message.

The typing in progress engine must work in real time! This meaning that if you start typing to a certain user this user will be able to see that you are typing.

This engine must have/display:

    A websocket to establish the connection with both users
    An animation so that the user can see that you are typing, this animation should be smooth (no interruptions/janks) and just enough to draw attention for the user to see (user friendly)
    The name of the user that is typing
    Whenever the user stops typing or finishes the conversation, it should not display the animation

To help with the display of the typing in progress you can take a look on the js event list, mainly the Keyboard events and the Focus events

## How to use the Application

1. Clone this repository on your computer using the following command:
```bash
  git clone https://learn.zone01dakar.sn/git/sbadiane/real-time-forum-typing-in-progress.git
```

2. Go to the project directory:
```bash
  cd real-time-forum-typing-in-progress/
```

3. Run the application using Node.js with the following command:
```bash
  go run .
```


## Authors

* [sbadiane](https://learn.zone01dakar.sn/git/sbadiane)
* [ssock](https://learn.zone01dakar.sn/git/ssock)


---

Connect, Share, Engage - Building Community Together!✨
