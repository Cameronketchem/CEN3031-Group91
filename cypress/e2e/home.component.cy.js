describe('Home to Wallet spec', () => {
  it('Goes to Connect Wallet page when clicked', () => {
    cy.visit('http://localhost:4200')
    cy.get('.welcome-msg-background').invoke('css', 'display', 'none');
    cy.get('.connect-wallet-btn').click()
    cy.url().should('include', 'http://localhost:4200/connect-wallet');
  });
})

describe('Wallet to Home spec', () => {
  it('Goes to home page when clicked', () => {
    cy.visit('http://localhost:4200/connect-wallet')
    cy.get('.welcome-msg-background').invoke('css', 'display', 'none');
    cy.get('.app-title').click()
    cy.url().should('include', 'http://localhost:4200');
  });
})