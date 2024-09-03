package genent

type CommitRollbacker interface {
	Commit() error
	Rollback() error
}

type Tx[T CommitRollbacker] struct {
	commited bool
	v        T
}

func WrapTx[T CommitRollbacker](v T) *Tx[T] {
	return &Tx[T]{v: v}
}

func (t *Tx[T]) Commit() error {
	if err := t.v.Commit(); err != nil {
		return err
	}

	t.commited = true

	return nil
}

func (t *Tx[T]) Rollback() error {
	if t.commited {
		return nil
	}

	return t.v.Rollback()
}

func (t *Tx[T]) Unwrap() T {
	return t.v
}
