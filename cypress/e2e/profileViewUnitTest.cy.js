describe('profileViewComponent', () => {
  it('Should create the app', () => {
    cy.visit('http://localhost:4200/user/1')
    cy.get('app-root').should('exist')
  });

  it(`Should have as title 'CrowdNFT'`, () => {
    cy.visit('http://localhost:4200/user/1')
    cy.get('app-root aT').should('have.text', 'CrowdNFT')
  });

  it('Should display error if user does not exist', () => {
    cy.visit('http://localhost:4200/user/PEAKTRADER')
    cy.get('.no-user-found').should('have.text', `No user was found with the ID 'PEAKTRADER' `)
});

it('Should display user info if user does exist', () =>{
  cy.window().then((win) => {
    cy.visit('http://localhost:4200/user/user123')
    win.validID = true
    win.invalidID = false
  });
  cy.get('app-root img').should('have.attr', 'src')
  });

  it('Redirects to nft page when searching for a nft', () => {
    cy.visit('http://localhost:4200');
    cy.get('.search-input-container input[type="text"]').type('20');
    cy.get('.search-input-container select').select('nft');
    cy.get('.search-icon').click();
    cy.url().should('include', 'nftcard/20');
  });
  
  it('Redirects to new profile view page when searching for a user', () => {
    cy.visit('http://localhost:4200');
    cy.get('.search-input-container input[type="text"]').type('CRAZYGUY24');
    cy.get('.search-input-container select').select('Users');
    cy.get('.search-icon').click();
    cy.url().should('include', 'user/CRAZYGUY24');
  });
  
   it('Goes to home page when clicked', () => {
    cy.visit('http://localhost:4200/connect-wallet')
    cy.get('.app-title').click()
    cy.url().should('include', 'http://localhost:4200');
  });
  
  it('Goes to Connect Wallet page when clicked', () => {
    cy.visit('http://localhost:4200')
    cy.get('.connect-wallet-btn').click()
    cy.url().should('include', 'http://localhost:4200/connect-wallet');
  });
});
