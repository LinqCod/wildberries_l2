package main

import (
	"fmt"
	"log"
)

type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
}

func newWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("Creating new account...")
	walletFacade := &walletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
	}
	fmt.Println("Account created")
	return walletFacade
}

func (w *walletFacade) addMoneyToWallet(accountID string, securityCode, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	return nil
}

func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	return nil
}

type account struct {
	name string
}

func newAccount(accountName string) *account {
	return &account{
		name: accountName,
	}
}

func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account name is incorrect")
	}
	fmt.Println("Account verified")
	return nil
}

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkCode(code int) error {
	if s.code != code {
		return fmt.Errorf("security code is incorrect")
	}
	fmt.Println("Security code verified")
	return nil
}

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("balance is not sufficient")
	}
	fmt.Println("Wallet balance is sufficient")
	w.balance -= amount
	return nil
}

func main() {
	walletFacade := newWalletFacade("Maxim", 1234)
	err := walletFacade.addMoneyToWallet("Maxim", 1234, 150)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	err = walletFacade.deductMoneyFromWallet("Maxim", 1234, 100)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
}
