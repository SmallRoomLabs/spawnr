# spawnr
### Run apps in background and kill them automatically at exit for golang


Sometimes you need to have some helper/server applications running in the background to properly run your code.  It can be something more advenced like redis/mysql or just some API servers as your backend.

You could then either start them up manually, start developing your code and then remember to shut them all down when finsihed.  Or you could have spawnr automatically run your API servers and databases for you and then shut it all down when your code exits.

---
#### Version history
 
v0.0 - Sep 29 2015 - Started project 