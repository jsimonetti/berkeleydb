// Copyright 2013 Matt Butcher
// MIT License
// This file contains a number of wrapper functions that make it
// possible for Go to call the method-style functions on BerkeleyDB
// structs.

#include <db.h> 

int go_db_open(DB *dbp, DB_TXN *txnid, char *filename, char *dbname, DBTYPE type, u_int32_t flags, int mode) {
  return dbp->open(dbp, txnid, filename, dbname, type, flags, mode);
}

int go_db_close(DB *dbp, u_int32_t flags) {
  if (dbp == NULL) return 0;

  return dbp->close(dbp, flags);
}

int go_db_get_open_flags(DB *dbp, u_int32_t *open_flags) {
  return dbp->get_open_flags(dbp, open_flags);
}

int go_db_remove(DB *dbp, char *filename) {
  // dbp->close(dbp, 0);
  return dbp->remove(dbp, filename, NULL, 0);
}

int go_db_rename(DB *dbp, char *oldname, char *newname) {
  return dbp->rename(dbp, oldname, NULL, newname, 0);
}
