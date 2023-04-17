describe('HomeComponent', () => {
  
  it('should create the app', () => {
    cy.visit('http://localhost:4200')
    cy.get('app-root').should('exist')
  });

  it(`should have as title 'CrowdNFT'`, () => {
    cy.visit('http://localhost:4200')
    cy.get('app-root aT').should('have.text', 'CrowdNFT')
  });

  it('should render welcome user', () => {
    cy.visit('http://localhost:4200')
    cy.get('app-root h1').should('have.text','Welcome to CrowdNFT!')
  });

  it('should have the github link', () => {
    cy.visit('http://localhost:4200')
    cy.get('app-root div.github-link').should('have.text',' Want to see how we did it?  Check out our repo!')
  });

  it('should contain a link to connect wallet', () => {
    cy.visit('http://localhost:4200')
    cy.get('app-root divB').should('have.attr', 'routerLink', 'connect-wallet') 
  });
})
