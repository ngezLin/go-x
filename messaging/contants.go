package messaging

type PublisherTxnStatusFlag int16

const (
	// PublisherTxnFlagUninitialized when txnmgr is created
	PublisherTxnFlagUninitialized PublisherTxnStatusFlag = 1 << iota
	// PublisherTxnFlagInitializing when txnmgr is initializing
	PublisherTxnFlagInitializing
	// PublisherTxnFlagReady when is ready to receive transaction
	PublisherTxnFlagReady
	// PublisherTxnFlagInTransaction when transaction is started
	PublisherTxnFlagInTransaction
	// PublisherTxnFlagEndTransaction when transaction will be committed
	PublisherTxnFlagEndTransaction
	// PublisherTxnFlagInError when having abortable or fatal error
	PublisherTxnFlagInError
	// PublisherTxnFlagCommittingTransaction when committing txn
	PublisherTxnFlagCommittingTransaction
	// PublisherTxnFlagAbortingTransaction when committing txn
	PublisherTxnFlagAbortingTransaction
	// PublisherTxnFlagAbortableError when producer encounter an abortable error
	// Must call AbortTxn in this case.
	PublisherTxnFlagAbortableError
	// PublisherTxnFlagFatalError when producer encounter an fatal error
	// Must Close an recreate it.
	PublisherTxnFlagFatalError
)
