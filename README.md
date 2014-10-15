Fuzzy Proxy
===========

A proxy build specifically for fuzzing images. 

Wait. What?
===========

Imagine you're working for a web hosting company, or similar, and you're in charge of looking at complaints against customer content.  Much of the kinds of things you'd be forced to look at would be horrific, terrifying, and unforgettable.  The kinds of things that some people find exciting are disturbing and damaging to others...

This proxy is for the others who have the unpleasant task of reviewing this content.  Fuzzy Proxy can allow you to see the images and even make determinations without having to view all of it in tortorously high resolution.

Getting
=======

If you don't feel comfortable builing the software yourself then you can download a binary from the build directory. There is currently one for osx and one for linux (both amd64)

Running
=======

```
Usage of fuzzyproxy:
  -l="127.0.0.1:7070": Proxy listening address
  -p="": Proxy password
  -u="": Proxy username.  Empty string disables authentication.
```

Once you have the proxy up and running you can then configure your browser to use the proxy (it's an http connect proxy)

What's it look like?
====================

Here's a small image example (snapped from imgur.com)
![Small Example](https://raw.githubusercontent.com/apokalyptik/fuzzyproxy/master/Example-Thumbnails.png "Small Example")

heres a larger image example (snapped from imgur.com again)
![Large Example](https://raw.githubusercontent.com/apokalyptik/fuzzyproxy/master/Example-Large.png "Large Example")
