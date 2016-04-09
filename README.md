# goDash
Little Go program to react to Amazon Dash Button clicks

All it does right now is to print "Dash!" whenever it sees an ARP request from a
dash button.

This is so for only tested under Linux and must be run as root or - much -
better granted the net_raw capability.  This is done using

```
% go build
% sudo setcap cap_net_raw+pe goDash
```

Keep in mind that this has to be repeated everytime you run `go build` because
the capability is removed every time the file is written or replaced.
