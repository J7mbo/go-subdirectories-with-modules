### What have you done!?

This is the simplest example of using subdirectories with go (which is not how the makers of go intended it to be used).

With go modules, we have app/directory/directory, which we didn't have before (I think!)

Nevertheless, it's a step closer to being able to write cleaner code in separate files instead of 1 million line files.

Run with `go run .`

---

### Additional things to point out

I'm not sticking with the convention of `x.com/y/z` that is demanded of me. If I'm not writing a library to be consumed
by others and is only standalone by itself, this is a much simpler way of doing things.

I've yet to stumble upon any problems with writing things this way. Let's see how it goes.
