describe('HomeComponent', () => {
  
  it('Should create the app', () => {
    cy.visit('http://localhost:4200')
    cy.get('app-root').should('exist')
  });

  it(`Should have as title 'CrowdNFT'`, () => {
    cy.visit('http://localhost:4200')
    cy.get('app-root aT').should('have.text', 'CrowdNFT')
  });


  it('Should have the github message', () => {
    cy.visit('http://localhost:4200')
    cy.get('app-root div.github-link').should('have.text',' Want to see how we did it?  Check out our repo!')
  });

  it('Should contain a link to connect wallet', () => {
    cy.visit('http://localhost:4200')
    cy.get('app-root divB').should('have.attr', 'routerLink', 'connect-wallet') 
  });

  it('Goes to Connect Wallet page when clicked', () => {
    cy.visit('http://localhost:4200')
    cy.get('.connect-wallet-btn').click()
    cy.url().should('include', 'http://localhost:4200/connect-wallet');
  });

  it('Has nft option', () =>{
    let dropdown;
    cy.visit('http://localhost:4200')
    dropdown = document.querySelectorAll("select")
    expect(dropdown).to.exist
    cy.get('app-root option').should('have.attr','value', 'nft')
  });
  
  it('Allows user or NFT search', () =>{
    let dropdown;
    cy.visit('http://localhost:4200')
    dropdown = document.querySelectorAll("select")
    expect(dropdown).to.exist;
    cy.get('.search-input-container select').should('have.text', 'NFTsUsers')
    
  });

  it('Redirects to nft page when searching for a nft', () => {
    cy.visit('http://localhost:4200');
    cy.get('.search-input-container input[type="text"]').type('1');
    cy.get('.search-input-container select').select('nft');
    cy.get('.search-icon').click();
    cy.url().should('include', 'nftcard/1');
  });

  it('Redirects to profile view page when searching for a user', () => {
    cy.visit('http://localhost:4200');
    cy.get('.search-input-container input[type="text"]').type('CRAZYGUY24');
    cy.get('.search-input-container select').select('Users');
    cy.get('.search-icon').click();
    cy.url().should('include', 'user/CRAZYGUY24');
  });
})


