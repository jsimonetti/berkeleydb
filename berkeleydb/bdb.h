// Copyright 2013 Matt Butcher
// MIT License

// Open a database.
extern int go_db_open(DB *, DB_TXN *, char *, char *, DBTYPE, u_int32_t, int);
// Close a database.
extern int go_db_close(DB *, u_int32_t);
extern int go_db_get_open_flags(DB *, u_int32_t *);
extern int go_db_remove(DB *, char *);
extern int go_db_rename(DB *, char *, char *);
