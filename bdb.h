// Copyright 2013 Matt Butcher
// MIT License

// Open a database.
extern int go_db_open(DB *, DB_TXN *, char *, char *, DBTYPE, u_int32_t, int);
// Close a database.
extern int go_db_close(DB *, u_int32_t);
extern int go_db_get_open_flags(DB *, u_int32_t *);
extern int go_db_remove(DB *, char *);
extern int go_db_rename(DB *, char *, char *);
extern int go_env_open(DB_ENV *, char *, u_int32_t, u_int32_t);
extern int go_env_close(DB_ENV *, u_int32_t);


// Convenience functions MAY BE REMOVED.
int go_db_put_string(DB *, char *, char *, u_int32_t);
int go_db_get_string(DB *, char *, char *);
int go_db_del_string(DB *, char *);
int go_db_cursor(DB *, DBC **);
int go_cursor_get_next(DBC *, char *, char *);
int go_cursor_get_prev(DBC *, char *, char *);
int go_cursor_get_first(DBC *, char *, char *);
int go_cursor_get_last(DBC *, char *, char *);
