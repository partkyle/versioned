Versioned
=========

A quick example of how to compile a version *into* a go application.

How to use
==========

Run the build script.

Ex:

```
~/versioned $ bin/build
VERSION not specified, using best known tagref "v0.0.1"
Success! Execute using "./versioned -version"
```

Then, by running `./versioned -version` you should see output similar to this:

```
~/versioned $ ./versioned -version
version: v0.0.1
```

You can manually specify a version as well:

```
~/versioned $ VERSION=v11.0 bin/build
Success! Execute using "./versioned -version"
~/versioned $ ./versioned -version
version: v11.0

```
