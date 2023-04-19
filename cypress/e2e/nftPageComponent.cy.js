describe('nftPageComponent', () => {
  it('Should create the app', () => {
    cy.visit('http://localhost:4200/nftcard/12')
    cy.get('app-root').should('exist')
  });

  it(`Should have as title 'CrowdNFT'`, () => {
    cy.visit('http://localhost:4200/nftcard/12')
    cy.get('app-root aT').should('have.text', 'CrowdNFT')
  });

  it('should not display NFT details when ID is invalid', () => {
    cy.window().then((win) => {
      cy.visit('http://localhost:4200/nftcard/1200000')
    });
    cy.contains('.no-nft-found').should('have.text', 'No NFT was found with the ID 1200000')
  });
  
  it('Should display data when ID is valid', () => {
    cy.window().then((win) => {
      cy.visit('http://localhost:4200/nftcard/12')
    });
    cy.get('app-root img').should('have.attr', 'src')
 });

 it('Redirects to new nft page when searching for a nft', () => {
  cy.visit('http://localhost:4200');
  cy.get('.search-input-container input[type="text"]').type('20');
  cy.get('.search-input-container select').select('nft');
  cy.get('.search-icon').click();
  cy.url().should('include', 'nftcard/20');
});

it('Redirects to profile view page when searching for a user', () => {
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

})