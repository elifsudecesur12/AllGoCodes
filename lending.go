package main

import (
    "fmt"
    "math/big"
)

type LendingPlatform struct {
    users        map[string]*User
    loans        map[string]*Loan
    nextLoanID   int
}

type User struct {
    name        string
    wallet      *big.Int
    borrowed    *big.Int
}

type Loan struct {
    id          string
    borrower    *User
    amount      *big.Int
    interest    float64
    repaid      bool
}

func NewLendingPlatform() *LendingPlatform {
    return &LendingPlatform{
        users: make(map[string]*User),
        loans: make(map[string]*Loan),
    }
}

func (lp *LendingPlatform) CreateUser(name string, wallet *big.Int) {
    lp.users[name] = &User{
        name: name,
        wallet: wallet,
        borrowed: new(big.Int),
    }
}

func (lp *LendingPlatform) CreateLoan(borrowerName string, amount *big.Int, interest float64) string {
    user, exists := lp.users[borrowerName]
    if !exists {
        return "User not found"
    }

    loanID := fmt.Sprintf("Loan%d", lp.nextLoanID)
    loan := &Loan{
        id: loanID,
        borrower: user,
        amount: amount,
        interest: interest,
        repaid: false,
    }

    user.borrowed.Add(user.borrowed, amount)
    lp.loans[loanID] = loan
    lp.nextLoanID++

    return loanID
}

func (lp *LendingPlatform) RepayLoan(loanID string, amount *big.Int) string {
    loan, exists := lp.loans[loanID]
    if !exists {
        return "Loan not found"
    }

    if loan.repaid {
        return "Loan already repaid"
    }

    borrower := loan.borrower
    remainingBalance := new(big.Int).Sub(loan.amount, amount)

    if remainingBalance.Cmp(borrower.wallet) > 0 {
        return "Insufficient funds to repay the loan"
    }

    borrower.wallet.Sub(borrower.wallet, amount)
    borrower.borrowed.Sub(borrower.borrowed, amount)

    if remainingBalance.Cmp(big.NewInt(0)) == 0 {
        loan.repaid = true
    } else {
        loan.amount = remainingBalance
    }

    return "Loan repayment successful"
}

func main() {
    lp := NewLendingPlatform()

    lp.CreateUser("Alice", big.NewInt(100))
    lp.CreateUser("Bob", big.NewInt(50))

    loanID := lp.CreateLoan("Alice", big.NewInt(75), 0.1)
    fmt.Printf("Loan ID: %s\n", loanID)

    repaymentStatus := lp.RepayLoan(loanID, big.NewInt(85))
    fmt.Printf("Repayment Status: %s\n", repaymentStatus)
}
