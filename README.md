# PerkDB

A persistent key-value database with support for arbitrary indexes. It
provides a REST API and stores documents as JSON data.

**NOTICE: Due to the change in license for BerkeleyDB, I have discontinued development on this package.**

## Sub-Packages

* `perkdb/berkeleydb`: Go bindings for the BerkeleyDB C library.

### BerkeleyDB Bindings

This package provides BerkeleyDB wrappers for the C library using `cgo`.

To build, you will need a relatively recent version of BerkeleyDB.
