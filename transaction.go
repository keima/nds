package nds

import (
	"appengine"
	"appengine/datastore"
)

// RunInTransaction works just like datastore.RunInTransaction however it
// interacts correcly with memory and memcache if a context generated by
// NewContext is used.
func RunInTransaction(c appengine.Context, f func(tc appengine.Context) error,
	opts *datastore.TransactionOptions) error {

	if cc, ok := c.(*context); ok {
		return runInTransaction(cc, f, opts)
	}
	return datastore.RunInTransaction(c, f, opts)
}

func runInTransaction(cc *context, f func(tc appengine.Context) error,
	opts *datastore.TransactionOptions) error {

	return datastore.RunInTransaction(cc, func(tc appengine.Context) error {
		tcc := &context{
			Context: tc,

			inTransaction: true,
		}
		return f(tcc)
	}, opts)
}
