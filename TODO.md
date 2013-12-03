TODO
====

* Investigate switching to 7x --sfx / code.google.com/p/lzma
* Write a go binary that does this instead of a shell script for increased portability
	* This may be impossible (or difficult) if the libraries don't expose info for self-extracting archives.
* I think there's a bug if you do no web hosting, because Get() doesn't call Initialize()
* Change Initialize() to init()?
