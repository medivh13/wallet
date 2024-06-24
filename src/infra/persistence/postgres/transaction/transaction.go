package transaction

import (
	dto "wallet/src/app/dto/transaction"
	"wallet/src/infra/helper"

	"log"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	Transfer(data *dto.TransferReqDTO) error
	GetTopTransactionByUser(walletID int64) ([]*dto.GetTopTransRespDTO, error)
	GetOverallTopTransactions() ([]*dto.GetOverallRespDTO, error)
}

const (
	Transfer = `UPDATE public.wallets
	SET balance = balance - $1
	WHERE id = $2
	AND balance >= $3
	`

	Receive = `UPDATE public.wallets
	SET balance = balance + $1
	WHERE id = (
		SELECT w.id
		FROM public.wallets w
		INNER JOIN public.users u ON u.id = w.user_id
		WHERE u.username = $2
		LIMIT 1
	)`

	TransferTransaction = `INSERT INTO public.transactions
	(wallet_id, type, amount)
	VALUES ($1, 'transfer', -$2::NUMERIC)
	`

	ReceiveTransaction = `INSERT INTO public.transactions
	(wallet_id, type, amount)
	VALUES (
		(SELECT w.id
		FROM public.wallets w
		INNER JOIN public.users u ON u.id = w.user_id
		WHERE u.username = $1
		LIMIT 1),
		'transfer', $2
	)
	`

	GetTopTransactionByUser = `SELECT 
	u.username,
    t.amount
	FROM 
	    public.transactions t
	INNER JOIN 
	    public.wallets w ON t.wallet_id = w.id
	INNER JOIN 
	    public.users u ON w.user_id = u.id
	WHERE 
	    t.type = 'transfer' AND w.id = $1
	ORDER BY 
	    ABS(t.amount) DESC
	LIMIT 10;
	`

	GetOverallTopTransactions = `
	SELECT 
    u.username,
    SUM(ABS(t.amount)) AS transacted_value
	FROM 
	    public.transactions t
	INNER JOIN 
	    public.wallets w ON t.wallet_id = w.id
	INNER JOIN 
	    public.users u ON w.user_id = u.id
	WHERE 
	    t.type = 'transfer'
	    AND t.amount < 0
	GROUP BY 
	    u.username
	ORDER BY 
	    transacted_value DESC
	LIMIT 10;
	`
)

var statement PreparedStatement

type PreparedStatement struct {
	getTopTransaction *sqlx.Stmt
	getOverall        *sqlx.Stmt
}

type transactionRepo struct {
	Connection *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	repo := &transactionRepo{
		Connection: db,
	}
	InitPreparedStatement(repo)
	return repo
}

func (p *transactionRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *transactionRepo) {
	statement = PreparedStatement{
		getTopTransaction: m.Preparex(GetTopTransactionByUser),
		getOverall:        m.Preparex(GetOverallTopTransactions),
	}
}

func (p *transactionRepo) Transfer(data *dto.TransferReqDTO) error {

	tx, err := p.Connection.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx TopUp  : ", err.Error())
		return err
	}
	defer func(tx *sqlx.Tx) {
		if err != nil {
			tx.Rollback()
			log.Println("Rolling back transaction due to:", err)
		} else {
			err = tx.Commit()
			if err != nil {
				log.Println("Failed to commit transaction:", err.Error())
			}
		}
	}(tx)

	resultTransfer, err := tx.Exec(Transfer, data.Amount, data.WalletIDSender, data.Amount)

	if err != nil {
		log.Println("Failed Query Transfer: ", err.Error())
		return err
	}

	row, _ := resultTransfer.RowsAffected()
	if row < 1 {
		log.Println("Failed Query Transfer: ", helper.ErrInsufficientBalance)
		err = helper.ErrInsufficientBalance
		return err
	}

	_, err = tx.Exec(TransferTransaction, data.WalletIDSender, data.Amount)

	if err != nil {
		log.Println("Failed Query Create Transaction Transfer : ", err.Error())
		return err
	}

	resultReceive, err := tx.Exec(Receive, data.Amount, data.ToUserName)

	if err != nil {
		log.Println("Failed Query Create Receive : ", err.Error())
		return err
	}

	rowReceive, _ := resultReceive.RowsAffected()
	if rowReceive < 1 {
		log.Println("Failed Query Receive: ", helper.ErrUserNotFound)
		err = helper.ErrUserNotFound
		return err
	}

	_, err = tx.Exec(ReceiveTransaction, data.ToUserName, data.Amount)

	if err != nil {
		log.Println("Failed Query Create Transaction Receive : ", err.Error())
		return err
	}

	return nil
}

func (p *transactionRepo) GetTopTransactionByUser(walletID int64) ([]*dto.GetTopTransRespDTO, error) {

	var resultData []*dto.GetTopTransRespDTO

	err := statement.getTopTransaction.Select(&resultData, walletID)

	if err != nil {
		return nil, err
	}

	return resultData, nil
}

func (p *transactionRepo) GetOverallTopTransactions() ([]*dto.GetOverallRespDTO, error) {

	var resultData []*dto.GetOverallRespDTO

	err := statement.getOverall.Select(&resultData)

	if err != nil {
		return nil, err
	}

	return resultData, nil
}
